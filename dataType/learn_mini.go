package main

import (
	"fmt"

	"github.com/sasbury/mini"
)

func LearnMini() {

	var filepath = "aaaa.conf"
	conf, err := mini.LoadConfiguration(filepath)
	if err == nil {
		fmt.Printf("%v\n", conf)
	}

	// sections := conf.SectionNames()
	dbPort := conf.IntegerFromSection("db", "port", 23306)
	// fmt.Printf("%v\n", sections)
	fmt.Printf("dbPort => %v\n", dbPort)
}
