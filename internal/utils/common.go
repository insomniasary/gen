package utils

import (
	"strings"
	"unicode"
)

func ListContain(target string, list []string) bool {
	for _, item := range list {
		if target == item {
			return true
		}
	}
	return false
}

func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return Lcfirst(strings.Replace(name, " ", "", -1))
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}