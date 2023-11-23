package main

import (
	"log"

	"github.com/gogo/protobuf/plugin/unmarshal"
	"github.com/gogo/protobuf/version"
)

func main() {
	marshaler := unmarshal.NewUnmarshal()
	print(version.AtLeast("v1.2.3"))
	log.Printf("%v", marshaler)
}
