// Copyright (C) 2024 T-Force I/O
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

package xlib

import "github.com/rs/zerolog"

type Bytes []byte

func (slice Bytes) MarshalZerologArray(arr *zerolog.Array) {
	arr.Bytes(slice)
}

type IntSlice []int

func (slice IntSlice) MarshalZerologArray(arr *zerolog.Array) {
	for _, i := range slice {
		arr.Int(i)
	}
}

type Int32Slice []int32

func (slice Int32Slice) MarshalZerologArray(arr *zerolog.Array) {
	for _, i := range slice {
		arr.Int32(i)
	}
}

type Int64Slice []int64

func (slice Int64Slice) MarshalZerologArray(arr *zerolog.Array) {
	for _, i := range slice {
		arr.Int64(i)
	}
}

type StringSlice []string

func (slice StringSlice) MarshalZerologArray(arr *zerolog.Array) {
	for _, i := range slice {
		arr.Str(i)
	}
}

type UintSlice []uint

func (slice UintSlice) MarshalZerologArray(arr *zerolog.Array) {
	for _, i := range slice {
		arr.Uint(i)
	}
}

type Uint32Slice []uint32

func (slice Uint32Slice) MarshalZerologArray(arr *zerolog.Array) {
	for _, i := range slice {
		arr.Uint32(i)
	}
}

type Uint64Slice []uint64

func (slice Uint64Slice) MarshalZerologArray(arr *zerolog.Array) {
	for _, i := range slice {
		arr.Uint64(i)
	}
}
