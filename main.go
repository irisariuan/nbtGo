package main

import (
	"encoding/json"
	"fmt"
	"goNbt/lib"
	"goNbt/lib/nbt"
	"os"
)

func main() {
	allBytes, err := lib.UnzipReader(os.Stdin)

	if err != nil {
		panic(err)
	}

	tag, err := nbt.ParseNBT(allBytes, false)
	if err != nil {
		panic(err)
	}
	jsonTag, err := json.MarshalIndent(tag, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonTag))
}
