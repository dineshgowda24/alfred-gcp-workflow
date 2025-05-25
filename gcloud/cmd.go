package gcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/config"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/log"
)

func runGCloudCmd[T any](cfg *Config, cmd string, extraArgs ...string) (T, error) {
	var out T
	args := buildArgs(cmd, cfg, extraArgs...)

	var stderr bytes.Buffer
	cmdExec := exec.Command(config.GetGCloudPath(), args...)
	cmdExec.Stderr = &stderr

	log.Info("gcloud command:", cmdExec.String())
	raw, err := cmdExec.Output()
	if err != nil {
		return out, fmt.Errorf("gcloud command failed: %s", stderr.String())
	}

	if err := json.Unmarshal(raw, &out); err != nil {
		return out, fmt.Errorf("failed to parse gcloud JSON output: %w", err)
	}

	return out, nil
}

func buildArgs(cmd string, cfg *Config, extra ...string) []string {
	args := append([]string{cmd}, extra...)

	if cfg != nil && cfg.Name != "" {
		args = append(args, "--configuration="+cfg.Name)
	}

	if !containsFlag(args, "--format") {
		args = append(args, "--format=json")
	}

	return args
}

func containsFlag(args []string, flag string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, flag) || arg == flag {
			return true
		}
	}
	return false
}
