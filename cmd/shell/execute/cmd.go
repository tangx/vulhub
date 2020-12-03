package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type Cmd struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func (c *Cmd) String() string {
	s := fmt.Sprintf("%s %s", c.Command, strings.Join(c.Args, " "))
	// logrus.Debugln(s)
	return s
}

func (c *Cmd) Execute() (string, error) {
	if c.Command == "" {
		return "", errors.New("Command is required")
	}

	cmd := c.String()
	out, err := exec.Command(c.Command, c.Args...).Output()
	logrus.Debugln(cmd)
	if err != nil {
		logrus.Errorf("\t exec failed\n\t%s\n", err.Error())
		return "", err
	}

	result := fmt.Sprintf("$ %s \n\n%s", cmd, out)

	return string(result), nil
}
