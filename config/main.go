package main

import (
	"fmt"
	"golang/config/conf"
)

func main() {
	conf, err := config.NewConfig("./aaa.conf")

	if err != nil {
		switch err {
		case config.SectionDelimNotMatch:
			fmt.Printf("errCode = %s, errMsg = %s", err, "section name must be surrounding by '[' and ']'.")
		case config.KeyValueFormatError:
			fmt.Printf("errCode = %s, errMsg = %s", err, "key and value must be connected by '='.")
		}
	} else {
		// fmt.Printf("config => %#v\n", conf)
		fmt.Printf("sections : => %#v\n", conf.GetSections())
		fmt.Printf("kay value info :\n %s\n", conf.ToString())
	}
}
