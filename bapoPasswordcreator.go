package main

import (
	"bapoPass/bapoPass"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(bapoPass.GeneratePassword(os.Args[1], os.Args[2]))
}
