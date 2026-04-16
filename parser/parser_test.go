// Copyright (C) 2024 T-Force I/O
// This file is part of TF Unifiler
//
// TF Unifiler is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// TF Unifiler is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with TF Unifiler. If not, see <https://www.gnu.org/licenses/>.

package parser

import (
	"strings"
	"testing"

	"github.com/tforceaio/tf-unifiler-go/xlib"
)

func TestParseSha256Error(t *testing.T) {
	var tests = []struct {
		name    string
		content string
		err     string
	}{
		{"wrong hash size", "ca47868bca0d531a275f20e99eb04ba1 go.mod", "invalid SHA-256 hash 'ca47868bca0d531a275f20e99eb04ba1'"},
		{"wrong hash size", "eb3ff11bb7ec6a6cfc5f0ffb391af473628a443b95f5188d9b147ed9bd9862b36d8c22d796b185d16422b7be9663b4a9b9f68f6dfda1f189c9b8343492acb2a3 go.mod", "invalid SHA-256 hash 'eb3ff11bb7ec6a6cfc5f0ffb391af473628a443b95f5188d9b147ed9bd9862b36d8c22d796b185d16422b7be9663b4a9b9f68f6dfda1f189c9b8343492acb2a3'"},
		{"space in hash", "b91806c5f58c62d69dda437d20ee 86eed0044f50b780a7795d5014ed04ad1bca go.mod", "invalid SHA-256 hash 'b91806c5f58c62d69dda437d20ee'"},
		{"invalid character", "b91806c5f58c62d69dda437d20ee86eed0044f50b780a7795d5014ed04ad1bXY go.mod", "invalid SHA-256 hash 'b91806c5f58c62d69dda437d20ee86eed0044f50b780a7795d5014ed04ad1bXY'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseSha256(strings.NewReader(tt.content))
			errs := xlib.ErrString(err)
			if errs != tt.err {
				t.Errorf("wrong error. Expected %q. Actual %q.", tt.err, errs)
			}
		})
	}
}
