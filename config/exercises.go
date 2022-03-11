package config

import (
	"gopkg.in/yaml.v3"
)

type Exec struct {
	Variables Variables `yaml:"variables"`
	Scripts   `yaml:",inline"`
}

func (e Exec) exec(script Script) error {
	if script == nil {
		return nil
	}
	vars := make([]string, 0)
	if e.Variables != nil {
		vars = e.Variables.Env()
	}
	return script.Exec(vars...)
}

func (e Exec) Setup() error   { return e.exec(e.BeforeScript) }
func (e Exec) Check() error   { return e.exec(e.Script) }
func (e Exec) Cleanup() error { return e.exec(e.AfterScript) }

type Exercise struct {
	Name     string
	Stage    string `yaml:"stage"`
	Exec     `yaml:",inline"`
	Markdown `yaml:",inline"`
}

type Exercises []Exercise

func (e *Exercises) UnmarshalYAML(value *yaml.Node) error {
	var name string
	execs := make(Exercises, 0)
	for _, node := range value.Content {
		switch node.Kind {
		case yaml.ScalarNode:
			name = node.Value
		case yaml.MappingNode:
			var exec Exercise
			if err := node.Decode(&exec); err != nil {
				return err
			}
			exec.Name = name
			execs = append(execs, exec)
		}
	}
	*e = execs
	return nil
}
