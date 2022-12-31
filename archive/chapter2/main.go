package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	msg := fmt.Sprintf("Hello %s\n", arg)

	f, err := os.Create("hello.txt") // 変更する
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}
	defer f.Close() // 追加する

	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Printf("failed to write message to file \n: %v", err)
		return
	}
	fmt.Println("Done!")
}
