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

package upgrade

import (
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/puctl/upgrade"
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/util/cli"
	"github.com/spf13/cobra"
)

func yumUpgradeCommand() *cobra.Command {
	cmd := cli.BuildCommand("yum-upgrade",
		"Upgrades the Pure1 Unplugged yum package",
		"Upgrades the Pure1 Unplugged package using yum. Use this once you've mounted an upgrade .iso either as a disk or using 'puctl upgrade mount-iso [file]'.",
		upgrade.YumUpgrade,
	)

	return cmd
}
