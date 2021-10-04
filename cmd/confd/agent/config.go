package agent

import (
	"context"
	"grape/api/confd"
	"os"
	"os/exec"
	"strings"
	"time"
)

func WriteConfigFiles(cf *confd.Configs) error {
	for _, file := range cf.FileConfigs {
		err := os.WriteFile(file.Path, []byte(file.Content), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleUpdateConfig(cf *confd.Configs) {
	appLock.Lock()
	defer appLock.Unlock()
	log.Infof("update configs, type: %v", cf.RestartType)
	switch cf.RestartType {
	case confd.Configs_None:
		log.Info("skip update")
		return
	case confd.Configs_WriteFiles:
		err := WriteConfigFiles(cf)
		if err != nil {
			log.Errorf("write config files err: %v", err)
			return
		}
	case confd.Configs_Kill:
		killApplication(app)
		<-app.done
		app = newAppCmd(cf)
	case confd.Configs_Command:
		if cf.RestartCommand == "" {
			log.Errorf("empty RestartCommand")
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		args := strings.Split(cf.RestartCommand, " ")
		err := exec.CommandContext(ctx, "sh", args...)
		if err != nil {
			log.Errorf("exec RestartCommand %s err: %v", cf.RestartCommand, err)
		}
	}
	log.Info("configs updated")
}
