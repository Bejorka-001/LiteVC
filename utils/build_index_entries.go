package utils

import datastructures "litevc/data_structures"

func BuildIndexEntries(
	newFiles,
	modifiedFiles,
	deletedFiles []string,
	currentSnapshot map[string]string) []datastructures.IndexEntry {

	var indexEntries []datastructures.IndexEntry

	for _, file := range newFiles {
		indexEntries = append(indexEntries, datastructures.IndexEntry{
			FilePath:   file,
			Action:     "ADD",
			HashedFile: currentSnapshot[file],
		})
	}

	for _, file := range modifiedFiles {
		indexEntries = append(indexEntries, datastructures.IndexEntry{
			FilePath:   file,
			Action:     "MODIFIED",
			HashedFile: currentSnapshot[file],
		})
	}

	for _, file := range deletedFiles {
		indexEntries = append(indexEntries, datastructures.IndexEntry{
			FilePath: file,
			Action:   "DELETED",
		})
	}

	return indexEntries
}
