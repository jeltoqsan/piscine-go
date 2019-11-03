package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	counter := 0
	filename := ""
	for i, v := range os.Args {
		counter++
		if i == 1 {
			filename = v
		}
	}
	if counter < 2 {
		fmt.Println("File name missing")
		return
	}
	if counter > 2 {
		fmt.Println("Too many arguments")
		return
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(data))
}
