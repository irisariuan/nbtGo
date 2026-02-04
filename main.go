package main

import (
	"fmt"
	"goNbt/lib"
	"goNbt/lib/nbt"
	"os"
)

func main() {
	filepath := "./test.dat"
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	allBytes, err := lib.UnzipReader(file)

	fmt.Println("All bytes", len(allBytes))
	if err != nil {
		panic(err)
	}

	tag, err := nbt.ParseNBT(allBytes, false)
	nbt.PrintTag(tag)
}
