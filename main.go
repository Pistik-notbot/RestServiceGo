package main

import (
	"os"
)

func main() {
	_, err := os.Create("out/newfile.txt")
	if err != nil {
		panic(err)
	}

}
