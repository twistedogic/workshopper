package config

import (
	"io/ioutil"
)

type File string

func (f File) Content() ([]byte, error) {
	return ioutil.ReadFile(string(f))
}
