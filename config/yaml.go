package config

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func FormatConfig(config interface{}) {
	v := reflect.ValueOf(config)
	for i := 0; i < v.Elem().NumField(); i++ {
		filed := v.Elem().Field(i)
		if filed.Kind().String() == "struct" {
			formatStruct(filed)
		}
	}
}

func formatStruct(fileds reflect.Value) {
	for i := 0; i < fileds.NumField(); i++ {
		child := fileds.Field(i)
		if child.Kind().String() == "struct" {
			formatStruct(child)
		} else if child.Kind().String() == "string" {
			if child.CanSet() {
				child.SetString(InitConfigEnv(child.String()))
			}
		}
	}
}

func InitConfigEnv(val string) string {
	var valueReg = regexp.MustCompile("env\\(.*?\\)")
	var matched = valueReg.FindAllString(val, -1)
	for i := 0; i < len(matched); i++ {
		envArg := getEnvArgs(matched[i])
		envVal := os.Getenv(envArg)
		if envVal == "" {
			fmt.Println(envArg, " env args not be null.")
		}
		val = strings.Replace(val, matched[i], envVal, -1)
	}
	return val
}

func getEnvArgs(match string) string {
	return match[4 : len(match)-1]
}
