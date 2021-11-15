package cmd

import (
	"context"
	"grape/api/v1/confd"
	"os"
)

var (
	lastConfigVersion = ""
)

func WriteConfigFiles(cf *confd.Configs) error {
	if lastConfigVersion == cf.Version {
		log.Infof("version %s loaded, skip update", cf.Version)
		return nil
	}
	for _, file := range cf.FileConfigs {
		err := os.WriteFile(file.Path, []byte(file.Content), 0644)
		if err != nil {
			return err
		} else {
			log.Infof("write fileConfig: %s", file.Path)
		}
	}
	lastConfigVersion = cf.Version
	return nil
}

func handleUpdateConfigs(ctx context.Context, ch <-chan *confd.Configs, app *Application) {
	for {
		select {
		case <-ctx.Done():
			return
		case cf := <-ch:
			if cf.Version == "" {
				log.Warnf("get empty configs, ignore")
			} else {
				if err := updateConfig(cf, app); err != nil {
					log.Errorf("update config err: %v", err)
				}
			}
		}
	}
}

func updateConfig(cf *confd.Configs, app *Application) error {
	if !app.Started() {
		// first start application
		app.UpdateEnv(cf.EnvConfigs)
		if err := WriteConfigFiles(cf); err != nil {
			return err
		}
		return app.Start(cf.RunCmd)
	}

	log.Infof("update configs, type: %v", cf.RestartType)
	if cf.RestartType == confd.Configs_None {
		log.Info("skip update")
		return nil
	}

	// udpate config files
	if err := WriteConfigFiles(cf); err != nil {
		return err
	}

	// update application envs, not apply to runtime
	app.UpdateEnv(cf.EnvConfigs)

	// do restart application
	switch cf.RestartType {
	case confd.Configs_Kill:
		return app.RestartByKill(cf.RunCmd)
	case confd.Configs_Command:
		return app.RestartByCommand(cf.RunCmd, cf.RestartCommand)
	}
	return nil
}
