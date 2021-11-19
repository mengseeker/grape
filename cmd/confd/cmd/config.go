package cmd

import (
	"context"
	"grape/api/v1/confd"
	"os"
	"strings"
)

var (
	lastConfigVersion = ""
)

func WriteConfigFiles(cf *confd.Configs) error {
	if cf.Version == "" {
		return nil
	}
	if lastConfigVersion >= cf.Version {
		log.Infof("version %s loaded, ignore update", cf.Version)
		return nil
	}
	for _, file := range cf.FileConfigs {
		dir := file.Path[0:strings.LastIndex(file.Path, "/")]
		if dir != "" {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		if err := os.WriteFile(file.Path, []byte(file.Content), 0644); err != nil {
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
	log.Infof("update configs, type: %v", cf.RestartType)
	if cf.RestartType == confd.RestartType_None {
		log.Info("skip update")
		return nil
	}

	// udpate config files
	if err := WriteConfigFiles(cf); err != nil {
		return err
	}

	// do restart application
	switch cf.RestartType {
	case confd.RestartType_Kill:
		return app.RestartByKill(cf.RunCmd)
	case confd.RestartType_Command:
		return app.RestartByCommand(cf.RunCmd, cf.RestartCommand)
	}
	return nil
}
