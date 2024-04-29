// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gb3sum_test

import (
	"testing"

	"github.com/sorairolake/gb3sum"
)

func TestEscapeFilename(t *testing.T) {
	t.Parallel()

	filename, escaped := gb3sum.EscapeFilename("README.md")
	if filename != "README.md" {
		t.Errorf("expected filename `%v`, got `%v`", "README.md", filename)
	}

	if escaped {
		t.Error("unexpected escaped filename")
	}
}

func TestEscapeFilenameEscaped(t *testing.T) {
	t.Parallel()

	filename, escaped := gb3sum.EscapeFilename("CODE\\_OF\n_CONDUCT\r.md")
	if filename != "CODE\\\\_OF\\n_CONDUCT\\r.md" {
		t.Errorf("expected filename `%v`, got `%v`", "CODE\\\\_OF\\n_CONDUCT\\r.md", filename)
	}

	if !escaped {
		t.Error("unexpected unescaped filename")
	}
}
