package datafile

import (
	"fmt"
	"io/ioutil"
)

func WriteF(filename2 string, data string) {
	databyte := []byte(data)
	err := ioutil.WriteFile(filename2, databyte, 0o644)
	if err != nil {
		fmt.Println(err)
	}
}
