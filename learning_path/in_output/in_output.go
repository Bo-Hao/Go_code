package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("file error!")
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	fmt.Println(stat, "stat")
	if err != nil {
		fmt.Println("stat error!")
		return
	}

	bs := make([]byte, stat.Size()) //stat.Size() is equal to its length.
	fmt.Println(stat.Size())

	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("_ error!")
		return
	}
	str := string(bs)
	fmt.Println(str)
}
