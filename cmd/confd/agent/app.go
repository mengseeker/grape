package agent

import (
	"context"
	"grape/api/confd"
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

var (
	appLock sync.Mutex
	app     *application
)

type application struct {
	lock     sync.Locker
	cmd      *exec.Cmd
	signKill bool
	done     chan struct{}
}

func handleApplication(ctx context.Context, ch <-chan *confd.Configs) {
	cf := <-ch
	if cf.Version == "" {
		log.Warn("no configuration was found")
	}
	err := WriteConfigFiles(cf)
	if err != nil {
		log.Fatalf("fail to write config file: %v", err)
	}

	app = newAppCmd(cf)
	go handleSign()

	select {
	case <-ctx.Done():
		return
	case cf := <-ch:
		handleUpdateConfig(cf)
	}
}

func GetRunCommand(cf *confd.Configs) string {
	run := config.run
	if run == "" {
		run = cf.Run
	}
	if run == "" {
		log.Fatal("run command undefined, exit")
	}
	return run
}

func waitApplication(a *application) {
	pid := a.cmd.Process.Pid
	log.Infof("application started at %d", pid)
	err := a.cmd.Wait()
	close(a.done)
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.signKill {
		log.Warnf("application %d sign killed by confd", pid)
		return
	}
	if err != nil {
		log.Fatalf("application %d exit unexpected: %v", pid, err)
	} else {
		log.Infof("application %d complete, exit", pid)
		os.Exit(0)
	}
}

func killApplication(a *application) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.signKill = true
	a.cmd.Process.Signal(os.Interrupt)
	go killApplicationForce(a)
}

func killApplicationForce(a *application) {
	select {
	case <-a.done:
		return
	case <-time.After(ApplicationKillTimeout * time.Second):
		log.Warnf("application %d unable to exit in %ds, force killed", a.cmd.Process.Pid, ApplicationKillTimeout)
		a.cmd.Process.Kill()
	}
}

func newAppCmd(cf *confd.Configs) *application {
	run := GetRunCommand(cf)
	log.Infof("start application: %s", run)
	args := strings.Split(run, " ")
	c := exec.Command(args[0], args[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	// cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	c.Env = os.Environ()
	err := c.Start()
	if err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
	a := &application{
		lock:     &sync.Mutex{},
		cmd:      c,
		signKill: false,
		done:     make(chan struct{}),
	}
	go waitApplication(a)
	return a
}

func handleSign() {
	c := make(chan os.Signal, 3)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sign := <-c
	appLock.Lock()
	defer appLock.Unlock()
	log.Infof("get sign %v", sign)
	err := app.cmd.Process.Signal(sign)
	if err != nil {
		log.Fatalf("failed to notify application: %v", err)
	}
	<-time.After(time.Second * 10)
	log.Fatal("try exit application in 10s, but not exit in time")
	// os.Exit(1)
}
