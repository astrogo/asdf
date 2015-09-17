//+build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/serenize/snaker"
	yaml "gopkg.in/yaml.v2"
)

type Schema struct {
	ID         string `yaml:"id"`
	Title      string
	Tag        string
	Type       string
	Properties Properties
	Required   []string
	ExtraProps bool `yaml:"additionalProperties"`
}

type Properties map[string]Property

func (props *Properties) UnmashalYAML(unmarshal func(interface{}) error) error {
	return unmarshal(props)
}

type Property struct {
	Description string
	Type        string
	Format      string
}

func main() {
	name := os.Args[1]
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	var schema Schema
	err = yaml.Unmarshal(buf, &schema)
	if err != nil {
		panic(err)
	}

	fmt.Printf("properties=%v\n", schema.Properties)

	fmt.Printf("\n\n/////\n")
	tag := strings.Split(schema.Tag, "/")
	tname := strings.Split(tag[len(tag)-1], "-")[0]
	fmt.Printf("type %s struct {\n", snaker.SnakeToCamel(tname))
	for k, prop := range schema.Properties {
		fmt.Printf("\t%s %v `yaml:\"%s\"` // %q\n",
			snaker.SnakeToCamel(k),
			prop.Type, k,
			strings.TrimSpace(prop.Description),
		)
		//fmt.Printf("// %#v\n", prop)
	}
	fmt.Printf("}\n")
}
