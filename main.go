package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"goNbt/lib"
	"goNbt/lib/nbt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "serialize" {
		reader := bufio.NewReader(os.Stdin)
		allBytes, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		var tag nbt.TagCompound
		err = json.Unmarshal(allBytes, &tag)
		if err != nil {
			panic(err)
		}
		serializedBytes, err := nbt.SerializeTag(&tag, false)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(serializedBytes)
		return
	}
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
