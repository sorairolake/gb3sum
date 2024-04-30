// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
)

func unescapeFilename(filename string) (string, error) {
	return strconv.Unquote(`"` + filename + `"`)
}

type checksum struct {
	escapedFilename   string
	unescapedFilename string
	digest            []byte
}

func parseUntaggedChecksum(line string) (*checksum, error) {
	cksum := new(checksum)
	isEscaped := false

	if strings.HasPrefix(line, "\\") {
		isEscaped = true
		line = line[1:]
	}

	digestEnd := strings.Index(line, " ")
	if digestEnd == -1 {
		return nil, errors.New("gb3sum: invalid line")
	}

	digest, err := hex.DecodeString(line[:digestEnd])
	if err != nil {
		return nil, err
	}

	cksum.digest = digest

	if line[digestEnd+1] != '*' && line[digestEnd+1] != ' ' {
		return nil, errors.New("gb3sum: unexpected character")
	}

	escapedFilename := line[digestEnd+2:]
	cksum.escapedFilename = escapedFilename
	unescapedFilename := escapedFilename

	if isEscaped {
		f, err := unescapeFilename(escapedFilename)
		if err != nil {
			return nil, err
		}

		unescapedFilename = f
	}

	cksum.unescapedFilename = unescapedFilename

	return cksum, nil
}

func parseTaggedChecksum(line string) (*checksum, error) {
	cksum := new(checksum)
	isEscaped := false

	if strings.HasPrefix(line, "\\") {
		isEscaped = true
		line = line[1:]
	}

	if !strings.HasPrefix(strings.ToLower(line), "blake3") {
		return nil, errors.New("gb3sum: unknown algorithm")
	}

	if line[6] != ' ' {
		return nil, errors.New("gb3sum: unexpected character")
	}

	filenameStart := strings.Index(line, "(") + 1
	filenameEnd := strings.LastIndex(line, ")")
	escapedFilename := line[filenameStart:filenameEnd]
	cksum.escapedFilename = escapedFilename
	unescapedFilename := escapedFilename

	if isEscaped {
		f, err := unescapeFilename(escapedFilename)
		if err != nil {
			return nil, err
		}

		unescapedFilename = f
	}

	cksum.unescapedFilename = unescapedFilename

	if line[filenameEnd+1] != ' ' {
		return nil, errors.New("gb3sum: unexpected character")
	}

	if line[filenameEnd+2] != '=' {
		return nil, errors.New("gb3sum: unexpected character")
	}

	if line[filenameEnd+3] != ' ' {
		return nil, errors.New("gb3sum: unexpected character")
	}

	digest, err := hex.DecodeString(line[filenameEnd+4:])
	if err != nil {
		return nil, err
	}

	cksum.digest = digest

	return cksum, nil
}

func parseChecksum(line string) (*checksum, error) {
	if strings.HasPrefix(strings.ToLower(line), "blake3") {
		return parseTaggedChecksum(line)
	}

	return parseUntaggedChecksum(line)
}
