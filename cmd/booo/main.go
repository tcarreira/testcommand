package main

import (
	"fmt"

	"github.com/tcarreira/testcommand/pkg/cmd/testcommand"
)

func main() {
	fmt.Println("Hello from testcommand!")
	testcommand.Execute()
}
