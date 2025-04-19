package gcloud

import (
	"encoding/json"
	"log"
	"os/exec"

	aw "github.com/deanishe/awgo"
)

const gcloudPath = "/Users/dinesh.chikkanna/google-cloud-sdk/bin/gcloud"

func ListConfigs() ([]*Config, string, error) {
	out, err := exec.Command(gcloudPath,
		"config", "configurations", "list", "--format=json").Output()
	if err != nil {
		log.Println("LOG: error in gcloud command", err)
		return nil, "", err
	}

	var rawConfigs []GCloudConfig
	if err := json.Unmarshal(out, &rawConfigs); err != nil {
		log.Println("LOG: error unmarshalling JSON", err)
		return nil, "", err
	}

	log.Println("LOG: rawConfigs", rawConfigs)

	var configs []*Config
	var active string

	for _, raw := range rawConfigs {
		conf := raw.ToConfig()
		configs = append(configs, conf)
		if active == "" {
			active = raw.Name
		}
	}

	return configs, active, nil
}

func GetActiveConfig(wf *aw.Workflow) *Config {
	out, err := exec.Command(gcloudPath,
		"config", "configurations", "list", "--filter=is_active:true", "--format=json").Output()
	if err != nil {
		log.Println("LOG: error in gcloud command", err)
		return nil
	}

	var configs []GCloudConfig
	if err := json.Unmarshal(out, &configs); err != nil || len(configs) == 0 {
		log.Println("LOG: error unmarshalling JSON or no active config found", err)
		return nil
	}

	log.Println("LOG: active config", configs[0])

	return configs[0].ToConfig()
}
