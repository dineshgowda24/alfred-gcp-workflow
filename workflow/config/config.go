package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	aw "github.com/deanishe/awgo"
)

const fileName = "config.json"

type ConfigFile struct {
	GCloudPath       string `json:"gcloud_path"`
	GCloudConfigPath string `json:"gcloud_config_path"`
}

var cfgFile *ConfigFile

// Init loads the singleton config instance.
// Must be called once on workflow startup.
func Init(wf *aw.Workflow) {
	cfgFile = ensure(wf)
}

// GetConfigFile returns the loaded config singleton.
func GetConfigFile() *ConfigFile {
	return cfgFile
}

// GetGCloudPath returns the gcloud binary path from the config file.
func GetGCloudPath() string {
	cf := GetConfigFile()
	return cf.GCloudPath
}

// GetGCloudConfigPath returns the gcloud config path from the config file.
func GetGCloudConfigPath() string {
	cf := GetConfigFile()
	return cf.GCloudConfigPath
}

// UpdateConfigFile overwrites config.json with the given config.
func UpdateConfigFile(wf *aw.Workflow, cfg *ConfigFile) {
	write(wf, cfg)
}

func ensure(wf *aw.Workflow) *ConfigFile {
	path := configPath(wf)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create an empty config file
		write(wf, &ConfigFile{})
		return &ConfigFile{}
	} else if err != nil {
		panic(err)
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg ConfigFile
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}

func write(wf *aw.Workflow, cfg *ConfigFile) {
	path := configPath(wf)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(err)
	}

	if _, err := f.Write(data); err != nil {
		panic(err)
	}
}

func configPath(wf *aw.Workflow) string {
	return filepath.Join(wf.DataDir(), fileName)
}
