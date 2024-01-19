package main

import (
	"fmt"
	"go_error/application_clean"
	graphclean "go_error/graph_clean"
)

func main() {
	SaveSampleA()
}

func SaveSampleA() {
	cmd := application_clean.SaveSampleACommand{
		Field1: "invalidField1",
		Field2: "field2",
	}
	o := graphclean.SaveSampleAOutPort{}
	success, err := application_clean.SaveSampleA(cmd, &o)
	if !success {
		o.Emit()
		fmt.Printf(err.Error())
		fmt.Printf("%#v\n", o.Response)
	} else {
		fmt.Println("success")
	}
}
