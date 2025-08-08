package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// init initializes Viper configuration when the package is imported.
func init() {
	if err := NewViperConfig(); err != nil {
		panic(err)
	}
}

var once sync.Once

// NewViperConfig initializes the global Viper configuration exactly once.
// It is safe to call this function multiple times from different places.
func NewViperConfig() (err error) {
	once.Do(func() {
		err = newViperConfig()
	})
	return
}

// newViperConfig sets up Viper with:
// - config name/type and search path resolved relative to the caller
// - environment variable overrides (e.g. GOOGLE_API_KEY)
// - reads the config file into memory
func newViperConfig() error {
	relPath, err := getRelativePathFromCaller()
	if err != nil {
		return err
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(relPath)

	// Replace underscores with hyphens for env-to-key mapping
	// Example: LLM_GEMINI_APIKEY -> llm-gemini-apikey
	viper.EnvKeyReplacer(strings.NewReplacer("_", "-"))
	viper.AutomaticEnv()

	// Bind GOOGLE_API_KEY to config key: llm.gemini.apikey
	_ = viper.BindEnv("llm.gemini.apikey", "GOOGLE_API_KEY")
	return viper.ReadInConfig()
}

// getRelativePathFromCaller computes the config search path
// relative to the current working directory and the file location
// of this package, to make config discovery stable in different
// execution contexts (tests, binaries, IDE run, etc.).
func getRelativePathFromCaller() (relPath string, err error) {
	callerPwd, err := os.Getwd()
	if err != nil {
		return
	}
	_, here, _, _ := runtime.Caller(0)
	relPath, err = filepath.Rel(callerPwd, filepath.Dir(here))
	fmt.Printf("caller from: %s, here: %s, relpath: %s", callerPwd, here, relPath)
	return
}
