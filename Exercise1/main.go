package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	Success = iota
	Failed
)

func main() {
	var ifile *os.File
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error in reading file")
		os.Exit(Failed)

	}
	if fi.Size() > 0 {
		ifile = os.Stdin

	} else {
		if fi.Mode()&os.ModeCharDevice == os.ModeCharDevice {
			if len(os.Args) >= 2 {
				ifile, err = os.Open(os.Args[1])
			} else {
				ifile, err = os.Open(os.Args[0])
			}
			if err != nil {
				fmt.Println("Error in reading file")
				os.Exit(Failed)
			}
			defer ifile.Close()
		} else {
			ifile = os.Stdin
		}
	}
	bytes, err := ioutil.ReadAll(ifile)
	fmt.Printf("%x\n", md5.Sum(bytes))
}
