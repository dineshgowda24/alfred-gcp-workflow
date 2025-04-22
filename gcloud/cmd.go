package gcloud

import (
	"encoding/json"
	"os/exec"
)

const gcloudPath = "/Users/dinesh.chikkanna/google-cloud-sdk/bin/gcloud"

func runGCloudCmd[T any](config *Config, subcommand string, args ...string) (T, error) {
	var result T
	cmdArgs := buildGCloudArgs(subcommand, config, args...)

	output, err := exec.Command(gcloudPath, cmdArgs...).Output()
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(output, &result)
	return result, err
}

func buildGCloudArgs(subcommand string, config *Config, args ...string) []string {
	cmdArgs := []string{subcommand}
	cmdArgs = append(cmdArgs, args...)

	cmdArgs = addConfigFlag(cmdArgs, config)
	cmdArgs = addFormatFlag(cmdArgs)

	return cmdArgs
}

func addConfigFlag(args []string, config *Config) []string {
	if config != nil && config.Name != "" {
		args = append(args, "--configuration="+config.Name)
	}
	return args
}

func addFormatFlag(args []string) []string {
	args = append(args, "--format=json")
	return args
}
