package service

import (
	fmt "fmt"
	"strings"
)

func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}

func HelloDaerah(name string, gender string, city string) string {
	var daerahOutput string

	if strings.ToUpper(city) == strings.ToUpper("JAKARTA") {
		if strings.ToUpper(gender) == strings.ToUpper("MALE") {
			daerahOutput = fmt.Sprintln("Halo Mister, %s", name)
		} else {
			daerahOutput = fmt.Sprintln("Halo Miss, %s", name)
		}
	} else if strings.ToUpper(city) == strings.ToUpper("MEDAN") {
		if strings.ToUpper(gender) == strings.ToUpper("MALE") {
			daerahOutput = fmt.Sprintln("Horas Mister, %s", name)
		} else {
			daerahOutput = fmt.Sprintln("Horas Miss, %s", name)
		}
	} else if strings.ToUpper(city) == strings.ToUpper("BANDUNG") {
		if strings.ToUpper(gender) == strings.ToUpper("MALE") {
			daerahOutput = fmt.Sprintln("Wiunjeng Mister, %s", name)
		} else {
			daerahOutput = fmt.Sprintln("Wiunjeng Miss, %s", name)
		}
	}

	return daerahOutput
}
