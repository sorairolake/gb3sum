// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

// Gb3sum prints and checks BLAKE3 checksums.
//
// Usage:
//
//	gb3sum [OPTIONS] [FILE]...
//
// Arguments:
//
//	[FILE]
//		Files to hash, or checksum files to check.
//
// Options:
//
//	-c, --check
//		Read checksums from _FILE_ and check them.
//	-l, --length <BYTE>
//		The number of output bytes.
//	--tag
//		Print BSD-style output.
//	--ignore-missing
//		Ignore missing files when checking checksums.
//	-q, --quiet
//		Skip printing OK for each successfully verified file.
//	--status
//		Indicates the validation result with the exit status without
//		printing anything.
//	--strict
//		Exit non-zero if any line in the file is invalid.
//	-w, --warn
//		Warn about improperly formatted checksum lines.
//	-h, --help
//		Print help message.
//	-v, --version
//		Print version number.
//	--generate-completion <SHELL>
//		Generate shell completion.
package main
