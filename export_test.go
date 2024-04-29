// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gb3sum

type Checksum = checksum

var (
	EscapeFilename = escapeFilename

	UnescapeFilename      = unescapeFilename
	ParseUntaggedChecksum = parseUntaggedChecksum
	ParseTaggedChecksum   = parseTaggedChecksum
	ParseChecksum         = parseChecksum

	WriteUntaggedChecksum = writeUntaggedChecksum
	WriteTaggedChecksum   = writeTaggedChecksum
)

func (c *Checksum) EscapedFilename() string {
	return c.escapedFilename
}

func (c *Checksum) UnescapedFilename() string {
	return c.unescapedFilename
}

func (c *Checksum) Digest() []byte {
	return c.digest
}
