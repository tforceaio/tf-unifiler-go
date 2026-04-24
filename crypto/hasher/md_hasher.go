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
	"crypto/md5"

	"golang.org/x/crypto/md4"
)

// Compute the MD4 checksum of a file.
func HashMd4(fPath string) (*HashResult, error) {
	return hashFile(fPath, md4.New(), "md4")
}

// Compute the MD5 checksum of a file.
func HashMd5(fPath string) (*HashResult, error) {
	return hashFile(fPath, md5.New(), "md5")
}
