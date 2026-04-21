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
	"github.com/google/uuid"
	"github.com/tforce-io/tf-golib/opx/slicext"
)

type Mapping struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	HashID    Bytes32   `gorm:"column:hash_id"`
	Directory string    `gorm:"column:directory"`
	Name      string    `gorm:"column:name"`
	Extension string    `gorm:"column:extension"`

	SessionID uuid.UUID `gorm:"session_id"`
}

func (e *Mapping) FullName() string {
	if e.Extension == "" {
		return e.Name
	}
	return e.Name + e.Extension
}

func NewMapping(hashID Bytes32, directory, name, extension string) *Mapping {
	return &Mapping{
		HashID:    hashID,
		Directory: directory,
		Name:      name,
		Extension: extension,
	}
}

func (ctx *DbContext) GetMappingsByHashIDs(hashes []Bytes32) ([]*Mapping, error) {
	return ctx.findMappingsByHashIDs(hashes)
}

func (ctx *DbContext) GetMappingsBySha256s(hashes []string) ([]*Mapping, error) {
	return ctx.findMappingsBySha256s(hashes)
}

func (ctx *DbContext) SaveMappings(mappings []*Mapping) error {
	hashes := make([]Bytes32, len(mappings))
	for i, mapping := range mappings {
		hashes[i] = mapping.HashID
	}
	changedMappings, err := ctx.findMappingsByHashIDs(hashes)
	if err != nil {
		return err
	}
	newMappings := []*Mapping{}
	for _, mapping := range mappings {
		existed := slicext.ContainsFunc(changedMappings, mapping, areEqualMappings)
		if existed {
			continue
		}
		newMappings = append(newMappings, mapping)
		changedMappings = append(changedMappings, mapping)
	}
	return ctx.writeMappings(newMappings, []*Mapping{})
}

func (ctx *DbContext) findMappingsByHashIDs(hashes []Bytes32) ([]*Mapping, error) {
	var docs []*Mapping
	result := ctx.db.Model(&Mapping{}).
		Where("hash_id IN ?", hashes).
		Find(&docs)
	return docs, result.Error
}

func (ctx *DbContext) findMappingsBySha256s(hashes []string) ([]*Mapping, error) {
	var docs []*Mapping
	result := ctx.db.Model(&Mapping{}).
		InnerJoins("hashes ON hashes.id = mappings.hash_id AND hashes.sha256 IN ?", hashes).
		Find(&docs)
	return docs, result.Error
}

func (ctx *DbContext) writeMappings(newMappings []*Mapping, changedMappings []*Mapping) error {
	tx := ctx.db.Begin()
	for _, mapping := range newMappings {
		if mapping.ID == uuid.Nil {
			var err error
			mapping.ID, err = uuid.NewRandom()
			if err != nil {
				return err
			}
		}
		result := tx.Create(mapping)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	for _, mapping := range changedMappings {
		result := tx.Model(&Mapping{}).
			Where("id = ?", mapping.ID).
			Updates(map[string]interface{}{
				"hash_id":    mapping.HashID,
				"directory":  mapping.Directory,
				"name":       mapping.Name,
				"extension":  mapping.Extension,
				"session_id": mapping.SessionID,
			})
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}

func areEqualMappings(x, y *Mapping) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return x.HashID == y.HashID && x.Directory == y.Directory && x.Name == y.Name && x.Extension == y.Extension
}
