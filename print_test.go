// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gb3sum_test

import (
	"encoding/hex"
	"testing"

	"github.com/sorairolake/gb3sum"
)

func TestWriteUntaggedChecksum(t *testing.T) {
	t.Parallel()

	filename := "foo.txt"

	digest, err := hex.DecodeString("94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff")
	if err != nil {
		t.Fatal(err)
	}

	checksum := gb3sum.WriteUntaggedChecksum(filename, digest)
	expected := "94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff  foo.txt"

	if checksum != expected {
		t.Errorf("expected checksum line `%v`, got `%v`", expected, checksum)
	}
}

func TestWriteTaggedChecksum(t *testing.T) {
	t.Parallel()

	filename := "foo.txt"

	digest, err := hex.DecodeString("94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff")
	if err != nil {
		t.Fatal(err)
	}

	checksum := gb3sum.WriteTaggedChecksum(filename, digest)
	expected := "BLAKE3 (foo.txt) = 94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff"

	if checksum != expected {
		t.Errorf("expected checksum line `%v`, got `%v`", expected, checksum)
	}
}
