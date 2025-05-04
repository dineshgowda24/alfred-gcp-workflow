package gcloud

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"

	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/config"
)

var (
	configsLoaded     bool
	activeConfigCache *Config
	allConfigsCache   []*Config
)

type Config struct {
	Name    string
	Project string
	Region  *Region
}

func (c *Config) GetRegionName() string {
	if c.Region != nil {
		return c.Region.Name
	}
	return ""
}

// CacheKey returns a unique cache key for the config.
// We use the prefix combined with the project name to avoid collisions.
// Note: We intentionally avoid using the config name, because multiple configs
// can point to the same project or config names can change over time.
// This ensures the cache remains stable and doesn't invalidate unnecessarily.
func (c *Config) CacheKey(prefix string) string {
	base := prefix + "_" + c.Project
	if c.Region != nil && c.Region.Name != "" {
		base += "_" + c.Region.Name
	}

	return base
}

// GetActiveConfig lazy loads the active gcloud config
// subsequent calls will return the cached config
func GetActiveConfig() *Config {
	if activeConfigCache != nil {
		return activeConfigCache
	}

	path, err := activeConfigPath()
	if err != nil {
		log.Println("failed to get active config path:", err)
		return nil
	}

	activeConfigCache = loadConfig(path)
	return activeConfigCache
}

func activeConfigPath() (string, error) {
	path := getGCloudConfigPath()
	data, err := os.ReadFile(filepath.Join(path, "active_config"))
	if err != nil {
		return "", err
	}

	name := strings.TrimSpace(string(data))
	return filepath.Join(path, "configurations", "config_"+name), nil
}

func getGCloudConfigPath() string {
	path := config.GetGCloudConfigPath()
	return filepath.Clean(path)
}

func loadConfig(path string) *Config {
	iniFile, err := ini.Load(path)
	if err != nil {
		log.Println("error loading config file:", path, err)
		return nil
	}

	name := strings.TrimPrefix(filepath.Base(path), "config_")
	project := iniFile.Section("core").Key("project").String()

	return &Config{
		Name:    name,
		Project: project,
	}
}

// GetAllConfigs lazy loads all gcloud configs
// subsequent calls will return the cached configs
func GetAllConfigs() ([]*Config, error) {
	if configsLoaded {
		return allConfigsCache, nil
	}

	dir, err := configsDir()
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Println("error reading configurations directory:", err)
		return nil, err
	}

	var configs []*Config
	for _, f := range files {
		if !isValidConfigFile(f) {
			continue
		}
		cfg := loadConfig(filepath.Join(dir, f.Name()))
		if cfg != nil {
			configs = append(configs, cfg)
		}
	}

	allConfigsCache = configs
	configsLoaded = true
	return allConfigsCache, nil
}

func configsDir() (string, error) {
	return filepath.Join(getGCloudConfigPath(), "configurations"), nil
}

func isValidConfigFile(f os.DirEntry) bool {
	return !f.IsDir() && f.Type().IsRegular() && strings.HasPrefix(f.Name(), "config_")
}
