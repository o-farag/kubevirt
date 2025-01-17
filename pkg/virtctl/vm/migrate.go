/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2023 Red Hat, Inc.
 *
 */

package vm

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	v1 "kubevirt.io/api/core/v1"

	"kubevirt.io/kubevirt/pkg/virtctl/clientconfig"
	"kubevirt.io/kubevirt/pkg/virtctl/templates"
)

const COMMAND_MIGRATE = "migrate"

func NewMigrateCommand() *cobra.Command {
	c := Command{command: COMMAND_MIGRATE}
	cmd := &cobra.Command{
		Use:     "migrate (VM)",
		Short:   "Migrate a virtual machine.",
		Example: usage(COMMAND_MIGRATE),
		Args:    cobra.ExactArgs(1),
		RunE:    c.migrateRun,
	}
	cmd.Flags().BoolVar(&dryRun, dryRunArg, false, dryRunCommandUsage)
	cmd.SetUsageTemplate(templates.UsageTemplate())
	return cmd
}

func (o *Command) migrateRun(cmd *cobra.Command, args []string) error {
	vmiName := args[0]

	virtClient, namespace, _, err := clientconfig.ClientAndNamespaceFromContext(cmd.Context())
	if err != nil {
		return err
	}

	dryRunOption := setDryRunOption(dryRun)

	err = virtClient.VirtualMachine(namespace).Migrate(context.Background(), vmiName, &v1.MigrateOptions{DryRun: dryRunOption})
	if err != nil {
		return fmt.Errorf("Error migrating VirtualMachine %v", err)
	}

	fmt.Printf("VM %s was scheduled to %s\n", vmiName, o.command)

	return nil
}
