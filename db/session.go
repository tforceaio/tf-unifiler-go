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
	"time"

	"github.com/google/uuid"
)

// Session groups changes in other tables for easier tracking and monitoring.
type Session struct {
	ID   uuid.UUID `gorm:"column:id;primaryKey"`
	Time time.Time `gorm:"time"`
}

// Return new Session.
func NewSession(id uuid.UUID, time time.Time) *Session {
	return &Session{
		ID:   id,
		Time: time,
	}
}

// SessionChangesCount counts number of records in tracking tables.
type SessionChangesCount struct {
	Hash    int64
	Mapping int64
	Set     int64
	SetHash int64
}

// Count number of records for tables that support sessions.
func (c *DbContext) CountSessionChanges(id uuid.UUID) (*SessionChangesCount, error) {
	hashCount, err := c.Count(&Hash{}, "session_id = ?", id)
	if err != nil {
		return nil, err
	}
	mappingCount, err := c.Count(&Mapping{}, "session_id = ?", id)
	if err != nil {
		return nil, err
	}
	setCount, err := c.Count(&Set{}, "session_id = ?", id)
	if err != nil {
		return nil, err
	}
	setHashCount, err := c.Count(&SetHash{}, "session_id = ?", id)
	if err != nil {
		return nil, err
	}
	return &SessionChangesCount{
		Hash:    hashCount,
		Mapping: mappingCount,
		Set:     setCount,
		SetHash: setHashCount,
	}, nil
}

// Get Session by ID.
func (c *DbContext) GetSession(id uuid.UUID) (*Session, error) {
	return c.findSession(id)
}

// Get all Sessions ordered by time.
func (c *DbContext) GetSessions() ([]*Session, error) {
	return c.findSessions()
}

// Get latest Session ordered by time.
func (c *DbContext) GetLatestSession() (*Session, error) {
	sessions, err := c.findSessions()
	if err != nil {
		return nil, err
	}
	if len(sessions) == 0 {
		return nil, nil
	}
	return sessions[0], nil
}

// Save Sessions to database.
func (c *DbContext) SaveSessions(sessions []*Session) error {
	ids := make([]uuid.UUID, len(sessions))
	for i, session := range sessions {
		ids[i] = session.ID
	}
	changedSessions, err := c.findSessionsByIDs(ids)
	if err != nil {
		return err
	}
	changedSessionsMap := map[uuid.UUID]uuid.UUID{}
	for _, session := range changedSessions {
		changedSessionsMap[session.ID] = session.ID
	}
	newSessions := []*Session{}
	for _, session := range sessions {
		if _, ok := changedSessionsMap[session.ID]; ok {
			continue
		}
		newSessions = append(newSessions, session)
		changedSessionsMap[session.ID] = session.ID
	}
	return c.writeSessions(newSessions, []*Session{})
}

// Return Session that has specified ID.
func (c *DbContext) findSession(id uuid.UUID) (*Session, error) {
	var doc *Session
	result := c.db.Model(&Session{}).
		Where("id = ?", id).
		First(&doc)
	if c.isEmptyResultError(result.Error) {
		return nil, nil
	}
	return doc, result.Error
}

// Return all Sessions ordered by time descending.
func (c *DbContext) findSessions() ([]*Session, error) {
	var docs []*Session
	result := c.db.Model(&Session{}).
		Order("time DESC").
		Find(&docs)
	return docs, result.Error
}

// Return all Sessions that have specified SessionIDs.
func (c *DbContext) findSessionsByIDs(ids uuid.UUIDs) ([]*Session, error) {
	var docs []*Session
	result := c.db.Model(&Session{}).
		Where("id IN ?", ids).
		Order("time DESC").
		Find(&docs)
	return docs, result.Error
}

// Insert new Sessions and update old Sessions in one transaction.
func (c *DbContext) writeSessions(newSessions []*Session, _ []*Session) error {
	tx := c.db.Begin()
	for _, session := range newSessions {
		if session.ID == uuid.Nil {
			var err error
			session.ID, err = uuid.NewV7()
			if err != nil {
				return err
			}
		}
		result := tx.Create(session)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}
