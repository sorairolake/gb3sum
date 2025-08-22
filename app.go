// SPDX-FileCopyrightText: 2024 Shun Sakai
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeebo/blake3"
)

func readFile(filename string) (data []byte, err error) {
	var reader io.Reader

	switch filename {
	case "-":
		reader = os.Stdin
	default:
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}

		defer func() {
			if e := file.Close(); e != nil {
				err = e
			}
		}()

		reader = file
	}

	return io.ReadAll(reader)
}

func escapeFilename(filename string) (string, bool) {
	escaped := false

	escapedFilename := strconv.Quote(filename)
	escapedFilename = strings.Trim(escapedFilename, "\"")

	if escapedFilename != filename {
		filename = escapedFilename
		escaped = true
	}

	return filename, escaped
}

func run(cmd *cobra.Command, args []string) (bool, error) {
	isSucceeded := true

	if opt.generateCompletion != "" {
		if err := genCompletion(cmd, opt.generateCompletion); err != nil {
			return isSucceeded, err
		}

		return isSucceeded, nil
	}

	if len(args) == 0 {
		args = append(args, "-")
	}

	if opt.check {
		unformattedLines := 0
		numReadFailures := 0
		matchedChecksums := 0
		mismatchedChecksums := 0

		for _, checksumFile := range args {
			input, err := readFile(checksumFile)
			if err != nil {
				return isSucceeded, err
			}

			reader := bytes.NewReader(input)
			scanner := bufio.NewScanner(reader)
			lineNumber := 0

			for scanner.Scan() {
				line := scanner.Text()
				lineNumber++

				checksum, err := parseChecksum(line)
				if err != nil {
					unformattedLines++

					if opt.strict {
						isSucceeded = false
					}

					if opt.warn {
						fmt.Fprintf(os.Stderr, "gb3sum: %v: %v: improperly formatted BLAKE3 checksum line\n", checksumFile, lineNumber)
					}

					continue
				}

				verifyInput, err := readFile(checksum.unescapedFilename)
				if err != nil {
					numReadFailures++

					if !opt.ignoreMissing {
						isSucceeded = false

						fmt.Fprintf(os.Stderr, "gb3sum: %v\n", err)
					}

					continue
				}

				hasher := blake3.New()
				if _, err := hasher.Write(verifyInput); err != nil {
					return isSucceeded, err
				}

				digest := hasher.Digest()

				out := make([]byte, len(checksum.digest))
				if _, err := digest.Read(out); err != nil {
					panic(err)
				}

				if slices.Equal(checksum.digest, out) {
					matchedChecksums++

					if !opt.quiet {
						fmt.Printf("%v: OK\n", checksum.escapedFilename)
					}
				} else {
					mismatchedChecksums++
					isSucceeded = false

					if !opt.status {
						fmt.Printf("%v: FAILED\n", checksum.escapedFilename)
					}
				}
			}

			if err := scanner.Err(); err != nil {
				return isSucceeded, err
			}

			if !opt.status {
				switch unformattedLines {
				case 0:
					break
				case 1:
					fmt.Fprintf(os.Stderr, "gb3sum: WARNING: %v line is improperly formatted\n", unformattedLines)
				default:
					fmt.Fprintf(os.Stderr, "gb3sum: WARNING: %v lines are improperly formatted\n", unformattedLines)
				}

				if !opt.ignoreMissing {
					switch numReadFailures {
					case 0:
						break
					case 1:
						fmt.Fprintf(os.Stderr, "gb3sum: WARNING: %v listed file could not be read\n", numReadFailures)
					default:
						fmt.Fprintf(os.Stderr, "gb3sum: WARNING: %v listed files could not be read\n", numReadFailures)
					}
				}

				switch mismatchedChecksums {
				case 0:
					break
				case 1:
					fmt.Fprintf(os.Stderr, "gb3sum: WARNING: %v computed checksum did NOT match\n", mismatchedChecksums)
				default:
					fmt.Fprintf(os.Stderr, "gb3sum: WARNING: %v computed checksums did NOT match\n", mismatchedChecksums)
				}

				if opt.ignoreMissing && (matchedChecksums == 0) {
					isSucceeded = false

					fmt.Fprintf(os.Stderr, "gb3sum: %v: no file was verified\n", checksumFile)
				}
			}
		}
	} else {
		hasher := blake3.New()
		out := make([]byte, opt.length)

		for _, filename := range args {
			input, err := readFile(filename)
			if err != nil {
				return isSucceeded, err
			}

			if _, err := hasher.Write(input); err != nil {
				return isSucceeded, err
			}

			digest := hasher.Digest()

			clear(out)

			if _, err := digest.Read(out); err != nil {
				panic(err)
			}

			var output string

			filename, escaped := escapeFilename(filename)
			if escaped {
				output = "\\"
			}

			if opt.tag {
				output += writeTaggedChecksum(filename, out)
			} else {
				output += writeUntaggedChecksum(filename, out)
			}

			fmt.Println(output)
			hasher.Reset()
		}
	}

	return isSucceeded, nil
}
