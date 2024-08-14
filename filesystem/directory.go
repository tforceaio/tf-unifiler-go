package filesystem

import (
	"os"
	"path"
	"path/filepath"
)

func IsDirectory(dPath string) (bool, error) {
	fileInfo, err := os.Lstat(dPath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func IsDirectoryUnsafe(dPath string) bool {
	isDir, err := IsDirectory(dPath)
	if err != nil {
		panic(err)
	}
	return isDir
}

func listDirectory(dPath string) ([]*FsEntry, error) {
	entries, err := os.ReadDir(dPath)
	if err != nil {
		return []*FsEntry{}, err
	}
	contents := make([]*FsEntry, len(entries))
	for i, e := range entries {
		relativePath := path.Join(dPath, e.Name())
		absolutePath, err := filepath.Abs(relativePath)
		if err != nil {
			return []*FsEntry{}, err
		}
		content := &FsEntry{
			AbsolutePath: absolutePath,
			RelativePath: relativePath,
			Name:         e.Name(),
			IsDir:        e.IsDir(),
		}
		contents[i] = content
	}
	return contents, nil
}

func listEntries(entires []*FsEntry, maxDepth int, depth int) ([]*FsEntry, error) {
	contents := []*FsEntry{}
	for _, e := range entires {
		contents = append(contents, e)
		if (depth >= maxDepth && maxDepth >= 0) || !e.IsDir {
			continue
		}
		subEntries, err := listDirectory(e.RelativePath)
		if err != nil {
			return []*FsEntry{}, err
		}
		subContents, err := listEntries(subEntries, maxDepth, depth+1)
		if err != nil {
			return []*FsEntry{}, err
		}
		contents = append(contents, subContents...)
	}
	return contents, nil
}