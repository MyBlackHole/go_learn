package main

import (
	"log"
	"os"
	"time"

	"github.com/hibiken/asynq"
	"gopkg.in/yaml.v3"
)

func main() {
    provider := &FileBasedConfigProvider{filename: "./periodic_task_config.yml"}

    mgr, err := asynq.NewPeriodicTaskManager(
        asynq.PeriodicTaskManagerOpts{
            RedisConnOpt:               asynq.RedisClientOpt{Addr: "localhost:6379"},
            PeriodicTaskConfigProvider: provider,         // this provider object is the interface to your config source
            SyncInterval:               10 * time.Second, // this field specifies how often sync should happen
    })
    if err != nil {
        log.Fatal(err)
    }

    if err := mgr.Run(); err != nil {
         log.Fatal(err)
    }
}

// FileBasedConfigProvider implements asynq.PeriodicTaskConfigProvider interface.
type FileBasedConfigProvider struct {
     filename string
}

// Parses the yaml file and return a list of PeriodicTaskConfigs.
func (p *FileBasedConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
    data, err := os.ReadFile(p.filename)
    if err != nil {
        return nil, err
    }
    var c PeriodicTaskConfigContainer
    if err := yaml.Unmarshal(data, &c); err != nil {
        return nil, err
    }
    var configs []*asynq.PeriodicTaskConfig
    for _, cfg := range c.Configs {
         configs = append(configs, &asynq.PeriodicTaskConfig{Cronspec: cfg.Cronspec, Task: asynq.NewTask(cfg.TaskType, nil)})
    }
    return configs, nil
}

type PeriodicTaskConfigContainer struct {
    Configs []*Config `yaml:"configs"`
}

type Config struct {
    Cronspec string `yaml:"cronspec"`
    TaskType string `yaml:"task_type"`
}
