// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import "encoding/hex"

func writeUntaggedChecksum(filename string, digest []byte) string {
	return hex.EncodeToString(digest) + "  " + filename
}

func writeTaggedChecksum(filename string, digest []byte) string {
	return "BLAKE3 (" + filename + ") = " + hex.EncodeToString(digest)
}
