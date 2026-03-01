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

	fmt.Printf("Snapshot detected with commit ID: %s\n", snapshot)
}
