package goft

import (
	"encoding/json"
	"fmt"
)

type Model interface {
	String() string
}

type Models string

func Makemodels(v interface{}) Models {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
	}
	return Models(b)
}