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
	"github.com/tforce-io/tf-golib/strfmt"
)

func TestSaveMappings(t *testing.T) {
	execs, imgs := testingMappingData()

	tests := []struct {
		group    string
		mappings []*testingMapping
		expected int64
	}{
		{"new_mapping", []*testingMapping{execs[5], execs[6]}, 19},
		{"duplicated_mapping", []*testingMapping{imgs[7], imgs[7]}, 17},
		{"duplicated_mapping", []*testingMapping{imgs[9], imgs[15]}, 18},
		{"same_hash_same_filename", []*testingMapping{imgs[15], imgs[15]}, 18},
		{"same_hash_different_filename", []*testingMapping{imgs[16], imgs[17]}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			ctx := getAndReseedMappingDB("SaveMappings")
			var mappings []*Mapping
			for _, m := range tt.mappings {
				mappings = append(mappings, m.Mapping())
			}
			ctx.SaveMappings(mappings)
			count, err := ctx.Count(&Mapping{}, nil, nil)
			if err != nil {
				t.Error(err)
			}
			if count != tt.expected {
				t.Errorf("wrong number of records. expected %v actual %v", tt.expected, count)
			}
		})
	}
}

type testingMapping struct {
	ID        uuid.UUID
	HashID    Bytes32
	Directory string
	FileName  string
}

func (m *testingMapping) Mapping() *Mapping {
	fileName := strfmt.NewFileNameFromStr(m.FileName)
	return &Mapping{
		ID:        m.ID,
		HashID:    m.HashID,
		Directory: m.Directory,
		Name:      fileName.Name,
		Extension: fileName.Extension,
	}
}

func testingMappingData() (exes, imgs []*testingMapping) {
	exes = []*testingMapping{
		{uuid.MustParse("0194d500-52f3-70af-96b8-f9945e61926e"), uuidHashID(uuid.MustParse("0194d1c3-d795-746b-a504-83573e4137b2")), "/usr/bin", "curl"},
		{uuid.MustParse("0194d500-52f3-7cf8-8f16-5672d65c7125"), uuidHashID(uuid.MustParse("0194d1c4-08d1-79e6-94a1-a9891b835cfc")), "/usr/bin", "gcc"},
		{uuid.MustParse("0194d500-52f3-701c-957d-0b4079dc2df3"), uuidHashID(uuid.MustParse("0194d1c3-9aeb-76c9-9651-0d2d18b0591a")), "/usr/bin", "git"},
		{uuid.MustParse("0194d500-52f3-74d9-baab-484bb48dde7f"), uuidHashID(uuid.MustParse("0194d1c3-420e-711e-9e5b-4dded2d9ce7d")), "/usr/bin", "go"},
		{uuid.MustParse("0194d500-52f3-7821-9715-1dd3cc078cfb"), uuidHashID(uuid.MustParse("0194d1c3-e1d1-7546-a032-72cde06b61ee")), "/usr/bin", "wget"},
		{uuid.MustParse("0194d500-52f3-74d0-8ebb-624ee24e0a73"), uuidHashID(uuid.MustParse("0194d1c4-95b9-7c32-bc65-c193239c7d40")), "/usr/bin", "nano"},
		{uuid.MustParse("0194d500-52f3-721c-b654-6eda8e5edc90"), uuidHashID(uuid.MustParse("0194d1c5-3cfd-78c4-b210-7378a714dbd1")), "/usr/bin", "zypper"},
	}

	imgs = []*testingMapping{
		{uuid.MustParse("0194d507-231a-700d-96d8-415d47a4f3d5"), uuidHashID(uuid.MustParse("0194d1d2-2508-797d-a012-3e8c02381f95")), "D:\\Images", "104481404_2269289369892972_6614553252192578924.jpg"},
		{uuid.MustParse("0194d507-231a-7e08-a9ee-60e2046c3aa3"), uuidHashID(uuid.MustParse("0194d1d2-2509-7978-96ec-f7c649b3b17d")), "D:\\Images", "120865501_4306510232713090_6524281912074268580.jpg"},
		{uuid.MustParse("0194d507-231a-76cc-b8df-8e68e9ff9d00"), uuidHashID(uuid.MustParse("0194d1d2-2509-7979-999a-8533a5c500d7")), "D:\\Images", "122679148_2779317922314746_8369294723265122474.jpg"},
		{uuid.MustParse("0194d507-231a-7230-a191-bd34c2038551"), uuidHashID(uuid.MustParse("0194d1d2-2509-797a-b36f-1892faafd2ee")), "D:\\Images", "281827975_3848323758726757_4936926896804375945.jpg"},
		{uuid.MustParse("0194d507-231a-7403-a760-146762b85d25"), uuidHashID(uuid.MustParse("0194d1d2-2509-797b-8184-e56824828fca")), "D:\\Images", "284589156_10228821387030040_5461559314640999219.jpg"},
		{uuid.MustParse("0194d507-231a-7aad-98e7-ce1ab0f1dda1"), uuidHashID(uuid.MustParse("0194d1d2-2509-797c-b4e1-958b469a6a7f")), "D:\\Images", "347395508_813195557036681_967528494166602120.jpg"},
		{uuid.MustParse("0194d507-231a-7933-97c6-4150b58deba7"), uuidHashID(uuid.MustParse("0194d1d2-2509-797d-8071-ff23cf9f18e7")), "D:\\Images", "347431900_816433553377073_6154863431794469435.jpg"},
		{uuid.MustParse("0194d507-231a-740b-b89a-44eac538f676"), uuidHashID(uuid.MustParse("0194d1d2-2509-797e-9e8c-65491842335b")), "D:\\Images", "349104259_958682395315943_6289452280342549092.jpg"},
		{uuid.MustParse("0194d507-231a-7cba-b8cd-3e24d4545238"), uuidHashID(uuid.MustParse("0194d1d2-2509-797f-bef5-37bfca0aa255")), "D:\\Images", "383446133_871775087647419_7685832678178040742.jpg"},
		{uuid.MustParse("0194d507-231a-7b06-9d98-6aa1f6b198e7"), uuidHashID(uuid.MustParse("0194d1d2-2509-7980-8e4d-b3a8f9db938f")), "D:\\Images", "417404879_943973973956363_9043642269447999040.jpg"},
		{uuid.MustParse("0194d507-231a-7acc-ae06-8418e2b013ec"), uuidHashID(uuid.MustParse("0194d1d2-2509-7981-b4ff-57f9a5732ea3")), "D:\\Images", "431000064_815846923916849_548192370849311731.jpg"},
		{uuid.MustParse("0194d507-231a-7365-923b-17fe70b28d15"), uuidHashID(uuid.MustParse("0194d1d2-2509-7982-96f4-691a1425cdda")), "D:\\Images", "457251862_902488768586469_4412682217370698292.jpg"},
		{uuid.MustParse("0194d507-231a-7697-9356-809cf336b943"), uuidHashID(uuid.MustParse("0194d1d2-2509-7983-8f8d-f30575da37e5")), "D:\\Images", "461989168_946678077487846_1616035097810312250.jpg"},
		{uuid.MustParse("0194d507-231a-7ab7-9198-ebdaad573760"), uuidHashID(uuid.MustParse("0194d1d2-2509-7984-92c6-0eaec3ba18fa")), "D:\\Images", "470667312_918872370344734_2712499210299194559.jpg"},
		{uuid.MustParse("0194d507-231a-7f23-a42f-2636faa7db35"), uuidHashID(uuid.MustParse("0194d1d2-2509-7985-aa0a-0620fa8fece8")), "D:\\Images", "471139620_1137339247755450_8509052910057230594.jpg"},
		{uuid.MustParse("0194d507-231a-7f01-9ce8-11a23fafa794"), uuidHashID(uuid.MustParse("0194d1d2-2509-7986-b269-c5ac2883b479")), "D:\\Images", "471558929_10162384100542454_2069148967528258482.jpg"},
		{uuid.MustParse("0194d507-231a-7be9-80ea-367d9d9eca2d"), uuidHashID(uuid.MustParse("0194d1d2-2509-7987-8ab6-a27b9fb06dad")), "D:\\Images", "473314735_1116165696638735_6786896280360675715.jpg"},
		{uuid.MustParse("0194d507-231a-7c4b-bf5f-44983b5b1721"), uuidHashID(uuid.MustParse("0194d1d2-2509-7987-8ab6-a27b9fb06dad")), "D:\\Images", "473314735_1116165696638735_6786896280360675715_2.jpg"},
	}
	return
}

func getAndReseedMappingDB(function string) *DbContext {
	ctx := getTestDB("Mapping", function)
	execs, imgs := testingMappingData()
	var maps []*testingMapping
	maps = append(maps, execs[0:5]...)
	maps = append(maps, imgs[0:12]...)
	tx := ctx.db.Begin()
	tx.Where("1 = 1").Delete(&Mapping{})
	for _, m := range maps {
		tx.Create(m.Mapping())
	}
	tx.Commit()
	return ctx
}
