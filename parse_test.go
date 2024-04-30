// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"encoding/hex"
	"slices"
	"testing"
)

func TestUnescapeFilename(t *testing.T) {
	t.Parallel()

	filename, err := unescapeFilename("README.md")
	if err != nil {
		t.Fatal(err)
	}

	if filename != "README.md" {
		t.Errorf("expected filename `%v`, got `%v`", "README.md", filename)
	}
}

func TestUnescapeFilenameEscaped(t *testing.T) {
	t.Parallel()

	filename, err := unescapeFilename("CODE\\\\_OF\\n_CONDUCT\\r.md")
	if err != nil {
		t.Fatal(err)
	}

	if filename != "CODE\\_OF\n_CONDUCT\r.md" {
		t.Errorf("expected filename `%v`, got `%v`", "CODE\\_OF\n_CONDUCT\r.md", filename)
	}
}

func TestParseUntaggedChecksum(t *testing.T) {
	t.Parallel()

	line := "94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff  foo.txt"

	checksum, err := parseUntaggedChecksum(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedDigest, err := hex.DecodeString("94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff")
	if err != nil {
		t.Fatal(err)
	}

	if checksum.escapedFilename != "foo.txt" {
		t.Errorf("expected escaped filename `%v`, got `%v`", "foo.txt", checksum.escapedFilename)
	}

	if checksum.unescapedFilename != "foo.txt" {
		t.Errorf("expected unescaped filename `%v`, got `%v`", "foo.txt", checksum.unescapedFilename)
	}

	if !slices.Equal(checksum.digest, expectedDigest) {
		t.Errorf("expected digest `%v`, got `%v`", expectedDigest, checksum.digest)
	}
}

func TestParseUntaggedChecksumEscaped(t *testing.T) {
	t.Parallel()

	line := "\\bee561c8ab07fec8ec114df9ffff1fce5e1e31846483e110fe6f05582bd816be  CODE\\\\_OF\\n_CONDUCT\\r.md"

	checksum, err := parseUntaggedChecksum(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedDigest, err := hex.DecodeString("bee561c8ab07fec8ec114df9ffff1fce5e1e31846483e110fe6f05582bd816be")
	if err != nil {
		t.Fatal(err)
	}

	if checksum.escapedFilename != "CODE\\\\_OF\\n_CONDUCT\\r.md" {
		t.Errorf("expected escaped filename `%v`, got `%v`", "CODE\\\\_OF\\n_CONDUCT\\r.md", checksum.escapedFilename)
	}

	if checksum.unescapedFilename != "CODE\\_OF\n_CONDUCT\r.md" {
		t.Errorf("expected unescaped filename `%v`, got `%v`", "CODE\\_OF\n_CONDUCT\r.md", checksum.unescapedFilename)
	}

	if !slices.Equal(checksum.digest, expectedDigest) {
		t.Errorf("expected digest `%v`, got `%v`", expectedDigest, checksum.digest)
	}
}

func TestParseTaggedChecksum(t *testing.T) {
	t.Parallel()

	line := "BLAKE3 (foo.txt) = 94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff"

	checksum, err := parseTaggedChecksum(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedDigest, err := hex.DecodeString("94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff")
	if err != nil {
		t.Fatal(err)
	}

	if checksum.escapedFilename != "foo.txt" {
		t.Errorf("expected escaped filename `%v`, got `%v`", "foo.txt", checksum.escapedFilename)
	}

	if checksum.unescapedFilename != "foo.txt" {
		t.Errorf("expected unescaped filename `%v`, got `%v`", "foo.txt", checksum.unescapedFilename)
	}

	if !slices.Equal(checksum.digest, expectedDigest) {
		t.Errorf("expected digest `%v`, got `%v`", expectedDigest, checksum.digest)
	}
}

func TestParseTaggedChecksumEscaped(t *testing.T) {
	t.Parallel()

	line := "\\BLAKE3 (CODE\\\\_OF\\n_CONDUCT\\r.md) = bee561c8ab07fec8ec114df9ffff1fce5e1e31846483e110fe6f05582bd816be"

	checksum, err := parseTaggedChecksum(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedDigest, err := hex.DecodeString("bee561c8ab07fec8ec114df9ffff1fce5e1e31846483e110fe6f05582bd816be")
	if err != nil {
		t.Fatal(err)
	}

	if checksum.escapedFilename != "CODE\\\\_OF\\n_CONDUCT\\r.md" {
		t.Errorf("expected escaped filename `%v`, got `%v`", "CODE\\\\_OF\\n_CONDUCT\\r.md", checksum.escapedFilename)
	}

	if checksum.unescapedFilename != "CODE\\_OF\n_CONDUCT\r.md" {
		t.Errorf("expected unescaped filename `%v`, got `%v`", "CODE\\_OF\n_CONDUCT\r.md", checksum.unescapedFilename)
	}

	if !slices.Equal(checksum.digest, expectedDigest) {
		t.Errorf("expected digest `%v`, got `%v`", expectedDigest, checksum.digest)
	}
}
