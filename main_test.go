// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmdtest"
)

var testFiles = []string{
	"b3sums_512.txt",
	"b3sums_failed.txt",
	"b3sums_improperly_formatted.txt",
	"b3sums_missing_files.txt",
	"b3sums_missing_multiple_files.txt",
	"b3sums_multiple_failed.txt",
	"b3sums_multiple_improperly_formatted.txt",
	"b3sums_tagged.txt",
	"b3sums_untagged.txt",
	"foo.txt",
	"fox.txt",
}

func copyFile(srcFile, dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func TestCLI(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "windows" {
		t.SkipNow()
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if err := exec.Command("go", "build").Run(); err != nil {
		t.Fatal(err)
	}

	defer os.Remove("gb3sum")

	ts, err := cmdtest.Read("testdata")
	if err != nil {
		t.Fatal(err)
	}

	ts.Setup = func(rootDir string) error {
		for _, testFile := range testFiles {
			err := copyFile(filepath.Join(wd, "testdata", testFile), filepath.Join(rootDir, testFile))
			if err != nil {
				return err
			}
		}

		return nil
	}
	ts.Commands["gb3sum"] = cmdtest.Program("gb3sum")
	ts.Run(t, false)
}
