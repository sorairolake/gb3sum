// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import "os"

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
