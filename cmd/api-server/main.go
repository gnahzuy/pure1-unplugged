// Copyright 2019, Pure Storage Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/apis/server"
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/clients/elastic"
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/common/hooks"
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/version"
	log "github.com/sirupsen/logrus"
)

const (
	port       = 8080
	sourceName = "api-server"
)

func main() {
	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)

	err := server.ParseAPIServerEnvironmentVariables()
	if err != nil {
		log.WithError(err).Fatal("Error loading server environment variables, exiting...")
		os.Exit(1)
		return
	}

	databaseService, err := elastic.InitializeClient(server.APIServerEnv.ElasticHost, 0, time.Second*5)
	if err != nil {
		log.WithError(err).Fatal("Error initializing elastic client, exiting...")
		os.Exit(1)
		return
	}

	err = databaseService.CreateArrayTemplate(context.Background())
	if err != nil {
		log.WithError(err).Fatal("Error initializing array template")
		os.Exit(1)
		return
	}

	errorHook, err := hooks.NewErrorLogHook(sourceName, []log.Level{log.WarnLevel, log.ErrorLevel, log.FatalLevel}, databaseService)
	if err != nil {
		log.WithError(err).Fatal("Error creating ErrorLogHook, exiting...")
		os.Exit(1)
		return
	}

	log.AddHook(errorHook)

	router := server.NewRouter()
	log.WithFields(log.Fields{
		"version": version.Get(),
		"port":    port,
	}).Debug("Starting API server")

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"port":  port,
		}).Errorf("Failed to listen and serve")
	}
}
