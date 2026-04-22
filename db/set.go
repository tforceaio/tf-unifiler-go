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
	"github.com/google/uuid"
)

type Set struct {
	ID   uuid.UUID `gorm:"column:id;primaryKey"`
	Name string    `gorm:"column:name;uniqueIndex"`

	SessionID uuid.UUID `gorm:"session_id"`
}

func NewSet(name string) *Set {
	return &Set{
		Name: name,
	}
}

func (ctx *DbContext) GetSetByName(name string) (*Set, error) {
	return ctx.findSetByName(name)
}

func (ctx *DbContext) GetSetsByNames(names []string) ([]*Set, error) {
	return ctx.findSetsByNames(names)
}

func (ctx *DbContext) SaveSets(sets []*Set) error {
	names := make([]string, len(sets))
	for i, set := range sets {
		names[i] = set.Name
	}
	changedSets, err := ctx.findSetsByNames(names)
	if err != nil {
		return err
	}
	changedSetsMap := map[string]uuid.UUID{}
	for _, set := range changedSets {
		changedSetsMap[set.Name] = set.ID
	}
	newSets := []*Set{}
	for _, set := range sets {
		if _, ok := changedSetsMap[set.Name]; ok {
			continue
		}
		newSets = append(newSets, set)
		changedSetsMap[set.Name] = set.ID
	}
	return ctx.writeSets(newSets, []*Set{})
}

func (ctx *DbContext) findSetByName(name string) (*Set, error) {
	var doc *Set
	result := ctx.db.Model(&SetHash{}).
		Where("name = ?", name).
		First(&doc)
	if ctx.isEmptyResultError(result.Error) {
		return nil, nil
	}
	return doc, result.Error
}

func (ctx *DbContext) findSetsByNames(names []string) ([]*Set, error) {
	var docs []*Set
	result := ctx.db.Model(&Set{}).
		Where("0 = ? OR name IN ?", len(names), names).
		Order("name ASC").
		Find(&docs)
	return docs, result.Error
}

func (ctx *DbContext) writeSets(newSets []*Set, _ []*Set) error {
	tx := ctx.db.Begin()
	for _, set := range newSets {
		if set.ID == uuid.Nil {
			var err error
			set.ID, err = uuid.NewRandom()
			if err != nil {
				return err
			}
		}
		result := tx.Create(set)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}
