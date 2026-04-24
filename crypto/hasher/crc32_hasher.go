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

import "hash/crc32"

// Compute the CRC-32 checksum of a file using the IEEE polynomial.
func HashCrc32(fPath string) (*HashResult, error) {
	return hashFile(fPath, crc32.NewIEEE(), "crc32")
}

// Compute the CRC-32C checksum of a file using the Castagnoli polynomial.
func HashCrc32c(fPath string) (*HashResult, error) {
	return hashFile(fPath, crc32.New(crc32.MakeTable(crc32.Castagnoli)), "crc32c")
}
