// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gb3sum_test

import (
	"encoding/hex"
	"slices"
	"testing"

	"github.com/sorairolake/gb3sum"
)

func TestUnescapeFilename(t *testing.T) {
	t.Parallel()

	filename, err := gb3sum.UnescapeFilename("README.md")
	if err != nil {
		t.Fatal(err)
	}

	if filename != "README.md" {
		t.Errorf("expected filename `%v`, got `%v`", "README.md", filename)
	}
}

func TestUnescapeFilenameEscaped(t *testing.T) {
	t.Parallel()

	filename, err := gb3sum.UnescapeFilename("CODE\\\\_OF\\n_CONDUCT\\r.md")
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

	checksum, err := gb3sum.ParseUntaggedChecksum(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedDigest, err := hex.DecodeString("94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff")
	if err != nil {
		t.Fatal(err)
	}

	if checksum.EscapedFilename() != "foo.txt" {
		t.Errorf("expected escaped filename `%v`, got `%v`", "foo.txt", checksum.EscapedFilename())
	}

	if checksum.UnescapedFilename() != "foo.txt" {
		t.Errorf("expected unescaped filename `%v`, got `%v`", "foo.txt", checksum.UnescapedFilename())
	}

	if !slices.Equal(checksum.Digest(), expectedDigest) {
		t.Errorf("expected digest `%v`, got `%v`", expectedDigest, checksum.Digest())
	}
}

func TestParseUntaggedChecksumEscaped(t *testing.T) {
	t.Parallel()

	line := "\\bee561c8ab07fec8ec114df9ffff1fce5e1e31846483e110fe6f05582bd816be  CODE\\\\_OF\\n_CONDUCT\\r.md"

	checksum, err := gb3sum.ParseUntaggedChecksum(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedDigest, err := hex.DecodeString("bee561c8ab07fec8ec114df9ffff1fce5e1e31846483e110fe6f05582bd816be")
	if err != nil {
		t.Fatal(err)
	}

	if checksum.EscapedFilename() != "CODE\\\\_OF\\n_CONDUCT\\r.md" {
		t.Errorf("expected escaped filename `%v`, got `%v`", "CODE\\\\_OF\\n_CONDUCT\\r.md", checksum.EscapedFilename())
	}

	if checksum.UnescapedFilename() != "CODE\\_OF\n_CONDUCT\r.md" {
		t.Errorf("expected unescaped filename `%v`, got `%v`", "CODE\\_OF\n_CONDUCT\r.md", checksum.UnescapedFilename())
	}

	if !slices.Equal(checksum.Digest(), expectedDigest) {
		t.Errorf("expected digest `%v`, got `%v`", expectedDigest, checksum.Digest())
	}
}

func TestParseTaggedChecksum(t *testing.T) {
	t.Parallel()

	line := "BLAKE3 (foo.txt) = 94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff"

	checksum, err := gb3sum.ParseTaggedChecksum(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedDigest, err := hex.DecodeString("94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff")
	if err != nil {
		t.Fatal(err)
	}

	if checksum.EscapedFilename() != "foo.txt" {
		t.Errorf("expected escaped filename `%v`, got `%v`", "foo.txt", checksum.EscapedFilename())
	}

	if checksum.UnescapedFilename() != "foo.txt" {
		t.Errorf("expected unescaped filename `%v`, got `%v`", "foo.txt", checksum.UnescapedFilename())
	}

	if !slices.Equal(checksum.Digest(), expectedDigest) {
		t.Errorf("expected digest `%v`, got `%v`", expectedDigest, checksum.Digest())
	}
}

func TestParseTaggedChecksumEscaped(t *testing.T) {
	t.Parallel()

	line := "\\BLAKE3 (CODE\\\\_OF\\n_CONDUCT\\r.md) = bee561c8ab07fec8ec114df9ffff1fce5e1e31846483e110fe6f05582bd816be"

	checksum, err := gb3sum.ParseTaggedChecksum(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedDigest, err := hex.DecodeString("bee561c8ab07fec8ec114df9ffff1fce5e1e31846483e110fe6f05582bd816be")
	if err != nil {
		t.Fatal(err)
	}

	if checksum.EscapedFilename() != "CODE\\\\_OF\\n_CONDUCT\\r.md" {
		t.Errorf("expected escaped filename `%v`, got `%v`", "CODE\\\\_OF\\n_CONDUCT\\r.md", checksum.EscapedFilename())
	}

	if checksum.UnescapedFilename() != "CODE\\_OF\n_CONDUCT\r.md" {
		t.Errorf("expected unescaped filename `%v`, got `%v`", "CODE\\_OF\n_CONDUCT\r.md", checksum.UnescapedFilename())
	}

	if !slices.Equal(checksum.Digest(), expectedDigest) {
		t.Errorf("expected digest `%v`, got `%v`", expectedDigest, checksum.Digest())
	}
}
