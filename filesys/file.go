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

package filesys

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/tforce-io/tf-golib/opx"
)

type FsEntry struct {
	AbsolutePath string
	RelativePath string
	Name         string
	IsDir        bool
}

type FsEntries []*FsEntry

func (entries FsEntries) GetPaths() []string {
	fPaths := make([]string, len(entries))
	for i, e := range entries {
		fPaths[i] = e.RelativePath
	}
	return fPaths
}

func (entries FsEntries) GetAbsPaths() []string {
	fPaths := make([]string, len(entries))
	for i, e := range entries {
		fPaths[i] = e.AbsolutePath
	}
	return fPaths
}

func CreateEntry(fPath string) (*FsEntry, error) {
	absolutePath, err := GetAbsPath(fPath)
	if err != nil {
		return nil, err
	}
	fileInfo, err := os.Lstat(fPath)
	if err != nil {
		return nil, err
	}
	entry := &FsEntry{
		AbsolutePath: absolutePath,
		RelativePath: NormalizePath(fPath),
		Name:         fileInfo.Name(),
		IsDir:        fileInfo.IsDir(),
	}
	return entry, nil
}

func CreateHardlink(sPath, tPath string) error {
	ntPath := NormalizePath(tPath)
	parent, _ := filepath.Split(ntPath)
	if !IsExist(parent) {
		err := os.MkdirAll(parent, 0775)
		logger.Debug().Str("dir", parent).Str("target", tPath).Msgf("Created parent directory '%s'", parent)
		if err != nil {
			return err
		}
	}
	err := os.Link(sPath, tPath)
	if err == nil {
		logger.Debug().Str("src", sPath).Str("target", tPath).Msgf("Created link for '%s'", sPath)
	}
	return err
}

func GetAbsPath(fPath string) (string, error) {
	return filepath.Abs(fPath)
}

func IsAbsPath(fPath string) bool {
	return filepath.IsAbs(fPath)
}

func IsExist(fPath string) bool {
	_, err := os.Stat(fPath)
	return !os.IsNotExist(err)
}

func IsFile(fPath string) (bool, error) {
	fileInfo, err := os.Lstat(fPath)
	if err != nil {
		return false, err
	}
	return !fileInfo.IsDir(), nil
}

func IsFileUnsafe(fPath string) bool {
	isFile, err := IsFile(fPath)
	if err != nil {
		panic(err)
	}
	return isFile
}

func IsFileExist(fPath string) bool {
	fileInfo, err := os.Stat(fPath)
	if os.IsNotExist(err) {
		return false
	}
	return !fileInfo.IsDir()
}

func Join(elem ...string) string {
	return filepath.Join(elem...)
}

func List(fPaths []string, recursive bool) (FsEntries, error) {
	contents := make([]*FsEntry, len(fPaths))
	for i, p := range fPaths {
		entry, err := CreateEntry(p)
		if err != nil {
			return FsEntries{}, err
		}
		contents[i] = entry
	}
	maxDepth := opx.Ternary(recursive, -1, 1)
	if recursive {
		var err error
		contents, err = listEntries(contents, maxDepth, 0)
		if err != nil {
			return FsEntries{}, err
		}
	}
	return contents, nil
}

func NormalizePath(fPath string) string {
	return filepath.FromSlash(fPath) // use OS-native path separator
}

func WriteLines(fPath string, lines []string) error {
	f, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)
	for _, line := range lines {
		writer.WriteString(line)
		writer.WriteString("\n")
	}
	writer.Flush()

	return nil
}
