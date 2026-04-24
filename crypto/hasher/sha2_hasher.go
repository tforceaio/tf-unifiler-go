// Copyright (C) 2025 T-Force I/O
// This file is part of TFunifiler
//
// TFunifiler is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// TFunifiler is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with TFunifiler. If not, see <https://www.gnu.org/licenses/>.

package hasher

import (
	"crypto/sha256"
	"crypto/sha512"
)

// Compute SHA-224 checksum of a file.
func HashSha224(fPath string) (*HashResult, error) {
	return hashFile(fPath, sha256.New224(), "sha224")
}

// Compute SHA-256 checksum of a file.
func HashSha256(fPath string) (*HashResult, error) {
	return hashFile(fPath, sha256.New(), "sha256")
}

// Compute SHA-384 checksum of a file.
func HashSha384(fPath string) (*HashResult, error) {
	return hashFile(fPath, sha512.New384(), "sha384")
}

// Compute SHA-512 checksum of a file.
func HashSha512(fPath string) (*HashResult, error) {
	return hashFile(fPath, sha512.New(), "sha512")
}
