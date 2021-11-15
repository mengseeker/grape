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

func handleUpdateConfigs(ctx context.Context, ch <-chan *confd.Configs, runCmd string) {
	for {
		select {
		case <-ctx.Done():
			return
		case cf := <-ch:
			if cf.Version == "" {
				log.Warn("no configuration was found")
			}
			if err := WriteConfigFiles(cf); err != nil {
				log.Errorf("fail to write configs file: %v", err)
			}
		}
	}
}
