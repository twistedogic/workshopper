package config

import (
	"testing"

	"gopkg.in/yaml.v3"
)

type Executable interface {
	Setup() error
	Check() error
	Cleanup() error
}

func checkExec(t *testing.T, e Executable) {
	if err := e.Setup(); err != nil {
		t.Fatal(err)
	}
	if err := e.Check(); err != nil {
		t.Fatal(err)
	}
	if err := e.Cleanup(); err != nil {
		t.Fatal(err)
	}
}

func Test_Config(t *testing.T) {
	cases := map[string]struct {
		input string
	}{
		"base": {
			input: "testdata/base.yaml",
		},
	}
	for name := range cases {
		tc := cases[name]
		t.Run(name, func(t *testing.T) {
			c, err := FromYAMLFile(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			b, err := yaml.Marshal(&c)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(b))
			for _, e := range c.Exercises {
				checkExec(t, e)
			}
		})
	}
}
