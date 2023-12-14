// Copyright The Enterprise Contract Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package validate

import (
	"github.com/spf13/cobra"

	"github.com/enterprise-contract/ec-cli/internal/definition"
	"github.com/enterprise-contract/ec-cli/internal/image"
	"github.com/enterprise-contract/ec-cli/internal/input"
)

var ValidateCmd *cobra.Command

func init() {
	ValidateCmd = NewValidateCmd()
}

func init() {
	ValidateCmd.AddCommand(validateImageCmd(image.ValidateImage))
	ValidateCmd.AddCommand(validateDefinitionCmd(definition.ValidateDefinition))
	ValidateCmd.AddCommand(validateInputCmd(input.ValidateInput))
}

func NewValidateCmd() *cobra.Command {
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate conformance with the Enterprise Contract",
	}
	validateCmd.PersistentFlags().Bool("show-successes", false, "")
	return validateCmd
}
