package main

import (
	"autofunc/autofunc"
	"fmt"
	"os"
)

func main() {

	var runing bool = true
	for runing {
		fmt.Println("\n=================== CPU INFORMATION ===================================")
		autofunc.ProgMain(os.Getpid())
		fmt.Println("\n==================== NETWORK INFORMATION ===============================")
		autofunc.InfoNetwork()
		fmt.Println("\n========================================================================")
		autofunc.FaisDodo()
	}

}
