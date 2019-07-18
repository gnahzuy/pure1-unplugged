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
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/puctl/infra"
	"github.com/PureStorage-OpenConnect/pure1-unplugged/pkg/util/cli"
	"github.com/spf13/cobra"
)

func kubeCommand() *cobra.Command {
	cmd := cli.BuildCommand("kube",
		"Shim to kubectl using local credentials",
		"CLI tools to manage the application platform (kubernetes) to run Pure1 Unplugged.\n\n**HINT: use '--' before passing args for kubectl, ex: 'puctl infra kube -- get pods --all-namespaces'",
		infra.KubeCTL,
	)

	return cmd
}
