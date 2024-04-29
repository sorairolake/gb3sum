// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"os"

	"github.com/sorairolake/gb3sum"
)

func main() {
	if err := gb3sum.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
