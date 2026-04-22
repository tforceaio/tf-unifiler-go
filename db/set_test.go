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

package db

import (
	"testing"

	"github.com/google/uuid"
)

func TestSaveSets(t *testing.T) {
	sets := testingSetData()
	tests := []struct {
		group    string
		mappings []*testingSet
		expected int64
	}{
		{"new_set", []*testingSet{sets[2]}, 3},
		{"duplicated_set_name", []*testingSet{sets[0], sets[1]}, 2},
		{"duplicated_set_name", []*testingSet{sets[1], sets[2], sets[3]}, 4},
		{"same_name", []*testingSet{sets[3], sets[3], sets[3]}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			ctx := getAndReseedSetDB("SaveSets")
			var sets []*Set
			for _, s := range tt.mappings {
				sets = append(sets, s.Set())
			}
			ctx.SaveSets(sets)
			count, err := ctx.Count(&Set{}, nil, nil)
			if err != nil {
				t.Error(err)
			}
			if count != tt.expected {
				t.Errorf("wrong number of records. expected %v actual %v", tt.expected, count)
			}
		})
	}
}

type testingSet struct {
	ID   uuid.UUID
	Name string
}

func (s *testingSet) Set() *Set {
	return &Set{
		ID:   s.ID,
		Name: s.Name,
	}
}

func testingSetData() []*testingSet {
	return []*testingSet{
		{uuid.MustParse("0194d4fd-2ca9-7356-9c26-f289b32cdd9a"), "exes"},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), "imgs"},
		{uuid.MustParse("0194d7e9-14f4-76b9-ab96-bb335dd97619"), "isos"},
		{uuid.MustParse("0194d8d7-5fb9-7821-8e40-cf5cde6f7d6c"), "photos"},
	}
}

func getAndReseedSetDB(function string) *DbContext {
	ctx := getTestDB("Set", function)
	sets := testingSetData()
	tx := ctx.db.Begin()
	tx.Where("1 = 1").Delete(&Set{})
	for _, s := range sets[0:2] {
		tx.Create(s.Set())
	}
	tx.Commit()
	return ctx
}
