package main

import (
	"os"

	"github.com/huof6829/testproj/hello/aaa"
	"github.com/huof6829/testproj/utils"
)

func main() {
	// fmt.Println(hello.Greet())

	rootCmd := aaa.NewAccountNonce()
	if err := utils.ExecuteCmd(rootCmd, "0"); err != nil {
		os.Exit(1)
	}
}
