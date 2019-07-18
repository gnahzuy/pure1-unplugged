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

package infra

import (
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/util/cli"
	"github.com/spf13/cobra"
)

// Command is the main command for the "infra" group of commands
func Command() *cobra.Command {
	subcmds := []*cobra.Command{
		initCommand(),
		kubeCommand(),
		resetCommand(),
		statusCommand(),
		upgradeCommand(),
		versionCommand(),
		repoCommand(),
		deletePodCommand(),
		packLogsCommand(),
	}

	cmd := cli.BuildCommand(
		"infra",
		"Pure1 Unplugged App Infrastructure Tools",
		"Pure1 Unplugged App Infrastructure Tools for configuration, deployment, troubleshooting, etc",
		nil,
	)

	cmd.AddCommand(subcmds...)

	return cmd
}
