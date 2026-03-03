package cmd

import (
	"fmt"
	"litevc/utils"
)

func Detect() {
	commitId, err := utils.CheckSnapshot()

	if err != nil {
		fmt.Println("Something wrong happen: ", err)
		return
	}

	snapshot, err := utils.LoadSnapshot(commitId)
	if err != nil {
		fmt.Println("Error loading snapshot:", err)
		return
	}

	currentSnapshot, err := utils.BuildCurrentSnapshot(".")
	if err != nil {
		fmt.Println("Error building current snapshot:", err)
		return
	}

	added, modified, deleted := utils.CompareSnapshot(snapshot, currentSnapshot)

	fmt.Printf("Added files: %v\n", added)
	fmt.Printf("Modified files: %v\n", modified)
	fmt.Printf("Deleted files: %v\n", deleted)

	//build index entries
	indexEntries := utils.BuildIndexEntries(added, modified, deleted, currentSnapshot)

	//write index entries to index file
	err = utils.WriteIndex(indexEntries)
	if err != nil {
		fmt.Println("Error writing index:", err)
		return
	}

	fmt.Println("Changes has added to stage area successfully and index file has updated")
}
