package main

import (
	"fmt"
	"testing"
)

func Test_Execute(t *testing.T) {
	c := Cmd{
		Command: "ls",
		Args:    []string{"-a", "-l", "-h"},
	}

	out, _ := c.Execute()

	fmt.Println(out)
}
