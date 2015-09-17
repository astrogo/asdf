//+build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	fname := os.Args[1]
	fmt.Printf("file: %q\n===\n", fname)
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n===\n", string(buf))
	var data yaml.MapSlice
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", data)
}
