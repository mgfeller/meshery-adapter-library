// Copyright 2020 Layer5, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package adapter_test contains tests that cannot be contained in package adapter because of circular dependencies.
package adapter_test

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/config"
	"github.com/layer5io/meshery-adapter-library/config/provider"
	"github.com/layer5io/meshery-adapter-library/status"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/utils"
	"path"
	"testing"
)

var (
	configRootPath = path.Join(utils.GetHome(), ".meshery")
	serviceName    = "test-adapter"

	serverDefaults = map[string]string{
		"name":     serviceName,
		"port":     "11111",
		"traceurl": "none",
		"version":  "v0.1.0",
	}

	meshSpecDefaults = map[string]string{
		"name":   "Adapter",
		"status": status.NotInstalled,
	}

	viperDefaults = map[string]string{
		"filepath": configRootPath,
		"filename": "adapter",
		"filetype": "yaml",
	}

	kubeConfigDefaults = map[string]string{
		provider.FilePath: configRootPath,
		provider.FileType: "yaml",
		provider.FileName: "kubeconfig",
	}

	operations = adapter.Operations{}
)

func createLogger() (logger.Handler, error) {
	logger, err := logger.New(serviceName, logger.Options{Format: logger.JsonLogFormat, DebugLevel: false})
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func createConfig() (config.Handler, error) {
	return provider.NewViper(provider.Options{
		ServerConfig:   serverDefaults,
		MeshSpec:       meshSpecDefaults,
		ProviderConfig: viperDefaults,
		Operations:     operations,
	})
}

func createKubeCfg() (config.Handler, error) {
	return provider.NewViper(provider.Options{
		ProviderConfig: kubeConfigDefaults,
	})
}

func TestCreateInstance(t *testing.T) {
	log, err := createLogger()
	if err != nil {
		t.Errorf("err = %v; want 'nil'", err)
	}

	cfg, err := createConfig()
	if err != nil {
		t.Errorf("err = %v; want 'nil'", err)
	}

	kubeCfg, err := createKubeCfg()
	if err != nil {
		t.Errorf("err = %v; want 'nil'", err)
	}

	handler := &adapter.Adapter{Config: cfg, Log: log, KubeconfigHandler: kubeCfg}
	handler.CreateInstance([]byte(""), "", nil)
}
