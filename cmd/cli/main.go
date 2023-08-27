package main

import (
	"fmt"
	"os"

	"berkeleytrue/gogal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
