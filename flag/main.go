// main
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addCmd := flag.NewFlagSet("adduser", flag.ExitOnError)
	data1 := addCmd.String("name", "", "输入姓名")
	data2 := addCmd.Int("age", 18, "输入年龄")
	data3 := addCmd.Bool("single", false, "光棍？")
	data4 := addCmd.Float64("balance", 0, "余额")
	if len(os.Args) < 2 {
		addCmd.Usage()
		return
	}
	addCmd.Parse(os.Args[2:])

	if addCmd.Parsed() {
		fmt.Printf("data1=%s\n", *data1)
		fmt.Printf("data2=%d\n", *data2)
		fmt.Printf("data3=%v\n", *data3)
		fmt.Printf("data4=%f\n", *data4)
		fmt.Println()
	}
	fmt.Println("Hello World!")
}
