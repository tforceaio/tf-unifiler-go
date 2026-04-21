// Copyright (C) 2025 T-Force I/O
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

package db

import (
	"testing"

	"github.com/google/uuid"
)

func TestSaveSetHashes(t *testing.T) {
	exes, imgs := testingSetHashData()
	tests := []struct {
		group    string
		mappings []*testingSetHash
		expected int64
	}{
		{"new_set_hash", []*testingSetHash{exes[5], exes[6]}, 19},
		{"duplicated_set_hash", []*testingSetHash{exes[0], exes[1]}, 17},
		{"duplicated_set_hash", []*testingSetHash{imgs[9], imgs[10], imgs[15]}, 18},
		{"same_set_hash", []*testingSetHash{imgs[9], imgs[15], imgs[15]}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			ctx := getAndReseedSetHashDB("SaveSetHashes")
			var setHashes []*SetHash
			for _, s := range tt.mappings {
				setHashes = append(setHashes, s.SetHash())
			}
			ctx.SaveSetHashes(setHashes)
			count, err := ctx.Count(&SetHash{}, nil, nil)
			if err != nil {
				t.Error(err)
			}
			if count != tt.expected {
				t.Errorf("wrong number of records. expected %v actual %v", tt.expected, count)
			}
		})
	}
}

type testingSetHash struct {
	SetID  uuid.UUID
	HashID Bytes32
}

func (s *testingSetHash) SetHash() *SetHash {
	return &SetHash{
		SetID:  s.SetID,
		HashID: s.HashID,
	}
}

func testingSetHashData() (exes, imgs []*testingSetHash) {
	exes = []*testingSetHash{
		{uuid.MustParse("0194d4fd-2ca9-7356-9c26-f289b32cdd9a"), uuidHashID(uuid.MustParse("0194d1c3-d795-746b-a504-83573e4137b2"))},
		{uuid.MustParse("0194d4fd-2ca9-7356-9c26-f289b32cdd9a"), uuidHashID(uuid.MustParse("0194d1c4-08d1-79e6-94a1-a9891b835cfc"))},
		{uuid.MustParse("0194d4fd-2ca9-7356-9c26-f289b32cdd9a"), uuidHashID(uuid.MustParse("0194d1c3-9aeb-76c9-9651-0d2d18b0591a"))},
		{uuid.MustParse("0194d4fd-2ca9-7356-9c26-f289b32cdd9a"), uuidHashID(uuid.MustParse("0194d1c3-420e-711e-9e5b-4dded2d9ce7d"))},
		{uuid.MustParse("0194d4fd-2ca9-7356-9c26-f289b32cdd9a"), uuidHashID(uuid.MustParse("0194d1c3-e1d1-7546-a032-72cde06b61ee"))},
		{uuid.MustParse("0194d4fd-2ca9-7356-9c26-f289b32cdd9a"), uuidHashID(uuid.MustParse("0194d1c4-95b9-7c32-bc65-c193239c7d40"))},
		{uuid.MustParse("0194d4fd-2ca9-7356-9c26-f289b32cdd9a"), uuidHashID(uuid.MustParse("0194d1c5-3cfd-78c4-b210-7378a714dbd1"))},
	}

	imgs = []*testingSetHash{
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2508-797d-a012-3e8c02381f95"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7978-96ec-f7c649b3b17d"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7979-999a-8533a5c500d7"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-797a-b36f-1892faafd2ee"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-797b-8184-e56824828fca"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-797c-b4e1-958b469a6a7f"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-797d-8071-ff23cf9f18e7"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-797e-9e8c-65491842335b"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-797f-bef5-37bfca0aa255"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7980-8e4d-b3a8f9db938f"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7981-b4ff-57f9a5732ea3"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7982-96f4-691a1425cdda"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7983-8f8d-f30575da37e5"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7984-92c6-0eaec3ba18fa"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7985-aa0a-0620fa8fece8"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7986-b269-c5ac2883b479"))},
		{uuid.MustParse("0194d651-0882-771e-ae5e-b179750127d0"), uuidHashID(uuid.MustParse("0194d1d2-2509-7987-8ab6-a27b9fb06dad"))},
	}
	return
}

func getAndReseedSetHashDB(function string) *DbContext {
	ctx := getTestDB("SetHash", function)
	execs, imgs := testingSetHashData()
	var maps []*testingSetHash
	maps = append(maps, execs[0:5]...)
	maps = append(maps, imgs[0:12]...)
	tx := ctx.db.Begin()
	tx.Where("1 = 1").Delete(&SetHash{})
	for _, m := range maps {
		tx.Create(m.SetHash())
	}
	tx.Commit()
	return ctx
}
