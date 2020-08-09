package main


import (
	"fmt"
	"os"
)

//os.Args demo
func main() {
	// os.Args是一个[]string
	// os.Args[0] 是当前执行的程序
	fmt.Println(os.Args)
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
