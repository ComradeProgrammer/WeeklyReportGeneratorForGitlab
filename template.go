package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed template.md
var Template string

func ReplaceInTemplate(m map[string]string) string {
	res := Template
	for k, v := range m {
		res = strings.ReplaceAll(res, fmt.Sprintf("{{%s}}", k), v)
	}
	return res
}
