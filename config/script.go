package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	tempDir     = "/tmp"
	tempPattern = "workshopper"
)

type Variables map[string]string

func (vars Variables) Set() error {
	for k, v := range vars {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (vars Variables) Env() []string {
	envs := make([]string, 0, len(vars))
	for k, v := range vars {
		s := k + "=" + v
		envs = append(envs, s)
	}
	return envs
}

type Script []string

func (s Script) file() (*os.File, error) {
	f, err := ioutil.TempFile(tempDir, tempPattern)
	if err != nil {
		return f, err
	}
	for _, line := range s {
		if _, err := f.WriteString(line); err != nil {
			return f, err
		}
	}
	return f, nil
}

func (s Script) Exec(envs ...string) error {
	f, err := s.file()
	if err != nil {
		return err
	}
	filename := f.Name()
	defer os.Remove(filename)
	cmd := exec.Command("/bin/sh", filename)
	cmd.Env = envs
	b, err := cmd.CombinedOutput()
	fmt.Println(string(b))
	return err
}

type Scripts struct {
	BeforeScript Script `yaml:"before_script"`
	Script       Script `yaml:"script"`
	AfterScript  Script `yaml:"after_script"`
}
