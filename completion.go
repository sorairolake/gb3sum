// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func genCompletion(cmd *cobra.Command, shell string) error {
	var err error

	switch strings.ToLower(shell) {
	case "bash":
		err = cmd.Root().GenBashCompletionV2(os.Stdout, true)
	case "fish":
		err = cmd.Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		err = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
	case "zsh":
		err = cmd.Root().GenZshCompletion(os.Stdout)
	default:
		panic("gb3sum: unknown shell")
	}

	if err != nil {
		return err
	}

	return nil
}
