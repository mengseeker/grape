package cmd

import (
	confdv1 "grape/api/v1/confd"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	ApplicationKillTimeout = 30
)

type Application struct {
	runCmd string
	env    map[string]string

	cmd      *exec.Cmd
	signKill bool
	done     chan struct{}

	lock sync.Mutex
}

func NewApplication(runCmd string) *Application {
	return &Application{
		runCmd: runCmd,
	}
}

func (app *Application) UpdateEnv(env []*confdv1.EnvConfig) {
	app.lock.Lock()
	defer app.lock.Unlock()
	appEnv := map[string]string{}
	for _, e := range env {
		appEnv[e.Key] = e.Value
	}
	app.env = appEnv
}

func (app *Application) Started() bool {
	return app.cmd != nil
}

func (app *Application) CreateCmd(runCmd string) {
	log.Infof("exec %q", runCmd)
	args := strings.Split(runCmd, " ")
	app.cmd = exec.Command(args[0], args[1:]...)
	app.cmd.Stdout = os.Stdout
	app.cmd.Stderr = os.Stderr
	app.cmd.Env = os.Environ()
	for k, v := range app.env {
		app.cmd.Env = append(app.cmd.Env, k+"="+v)
	}
	app.done = make(chan struct{})
}

func (app *Application) TryStart(runCmd string) error {
	if app.runCmd != "" {
		runCmd = app.runCmd
	}

	if runCmd == "" {
		// if have no command to run, just return
		return nil
	}
	app.lock.Lock()
	defer app.lock.Unlock()

	app.CreateCmd(runCmd)

	err := app.cmd.Start()
	if err != nil {
		log.Fatalf("failed to start application %q: %v", runCmd, err)
	}
	go app.handleSign()
	go app.waitApplication()
	return nil
}

func (app *Application) RestartByKill(runCmd string) error {
	if app.runCmd != "" {
		runCmd = app.runCmd
	}

	// if not started, just start
	if !app.Started() {
		app.TryStart(runCmd)
		return nil
	}

	app.killApplication()

	// wait application process exited and function waitApplication returned
	<-app.done

	app.lock.Lock()
	defer app.lock.Unlock()
	app.cmd = nil

	if runCmd == "" && app.runCmd == "" {
		// if have no command to run, just return
		return nil
	}

	app.CreateCmd(runCmd)

	err := app.cmd.Start()
	if err != nil {
		log.Fatalf("failed to start application: %v", err)
	}

	go app.waitApplication()
	return nil
}

func (app *Application) RestartByCommand(runCmd string, restartCommand string) error {
	args := strings.Split(restartCommand, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	return cmd.Run()
}

func (app *Application) RunExecApplication(runCmd string) error {
	return app.TryStart(runCmd)
}

func (app *Application) waitApplication() {
	defer close(app.done)
	pid := app.cmd.Process.Pid
	log.Infof("application started at pid: %d", pid)
	err := app.cmd.Wait()
	app.lock.Lock()
	defer app.lock.Unlock()
	if app.signKill {
		log.Warnf("application(%d) sign killed by confd", pid)
		return
	}
	if err != nil {
		log.Fatalf("application(%d) exit unexpected: %v", pid, err)
		os.Exit(1)
	} else {
		log.Infof("application(%d) complete, exit", pid)
		os.Exit(0)
	}
}

func (app *Application) killApplication() {
	app.lock.Lock()
	defer app.lock.Unlock()
	app.signKill = true
	app.cmd.Process.Signal(os.Interrupt)
	go app.killApplicationForce()
}

func (app *Application) killApplicationForce() {
	select {
	case <-app.done:
		return
	case <-time.After(ApplicationKillTimeout * time.Second):
		log.Warnf("application %d unable to exit in %ds, force killed", app.cmd.Process.Pid, ApplicationKillTimeout)
		app.cmd.Process.Kill()
	}
}

func (app *Application) handleSign() {
	c := make(chan os.Signal, 3)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sign := <-c
	app.lock.Lock()
	defer app.lock.Unlock()
	log.Infof("sign %v, exiting", sign)
	err := app.cmd.Process.Signal(sign)
	if err != nil {
		log.Fatalf("failed to notify application: %v", err)
	}
	<-time.After(time.Second * 10)
	log.Fatal("try exit application in 10s, but not exit in time")
}
