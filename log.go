package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	nowDateString := fmt.Sprintf("%02d%02d%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Println(nowDateString)

	filename := "./db.qa.20200426.log"
	//ReadLine(filename)

	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		str, err := readLine(r)
		if err != nil {
			break
		}
		fmt.Println(str)
	}
}

// func ReadLine(filename string) {

// }

func readLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}
