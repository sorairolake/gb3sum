// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gb3sum

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "gb3sum [OPTIONS] [FILE]...",
	Short:   "Print and check BLAKE3 checksums",
	Version: version,
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		if ignoreMissing, _ := cmd.Flags().GetBool("ignore-missing"); ignoreMissing {
			if err := cmd.MarkFlagRequired("check"); err != nil {
				panic(err)
			}
		}

		if quiet, _ := cmd.Flags().GetBool("quiet"); quiet {
			if err := cmd.MarkFlagRequired("check"); err != nil {
				panic(err)
			}
		}

		if status, _ := cmd.Flags().GetBool("status"); status {
			opt.quiet = true
			if err := cmd.MarkFlagRequired("check"); err != nil {
				panic(err)
			}
		}

		if strict, _ := cmd.Flags().GetBool("strict"); strict {
			if err := cmd.MarkFlagRequired("check"); err != nil {
				panic(err)
			}
		}

		if warn, _ := cmd.Flags().GetBool("warn"); warn {
			if err := cmd.MarkFlagRequired("check"); err != nil {
				panic(err)
			}
		}

		generateCompletion, _ := cmd.Flags().GetString("generate-completion")
		switch strings.ToLower(generateCompletion) {
		case "", "bash", "fish", "powershell", "zsh":
			break
		default:
			return fmt.Errorf("gb3sum: unknown shell `%v`", generateCompletion)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		isSucceeded, err := run(cmd, args)
		if err != nil {
			return err
		}

		if !isSucceeded {
			os.Exit(1)
		}

		return nil
	},
	DisableFlagsInUseLine: true,
}

type options struct {
	check              bool
	length             uint
	tag                bool
	ignoreMissing      bool
	quiet              bool
	status             bool
	strict             bool
	warn               bool
	generateCompletion string
}

var opt options

func init() {
	RootCmd.Flags().BoolVarP(&opt.check, "check", "c", false, "read checksums from [FILE] and check them")
	RootCmd.Flags().UintVarP(&opt.length, "length", "l", 32, "the number of output bytes")
	RootCmd.Flags().BoolVar(&opt.tag, "tag", false, "print BSD-style output")
	RootCmd.Flags().BoolVar(&opt.ignoreMissing, "ignore-missing", false, "ignore missing files when checking checksums")
	RootCmd.Flags().BoolVarP(&opt.quiet, "quiet", "q", false, "skip printing OK for each successfully verified file")
	RootCmd.Flags().BoolVar(&opt.status, "status", false, "indicates the validation result with the exit status without printing anything")
	RootCmd.Flags().BoolVar(&opt.strict, "strict", false, "exit non-zero if any line in the file is invalid")
	RootCmd.Flags().BoolVarP(&opt.warn, "warn", "w", false, "warn about improperly formatted checksum lines")
	RootCmd.Flags().StringVar(&opt.generateCompletion, "generate-completion", "", "generate shell completion")

	RootCmd.MarkFlagsMutuallyExclusive("quiet", "status")
}