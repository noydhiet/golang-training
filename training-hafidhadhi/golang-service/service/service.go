package service

import (
	"fmt"
	"strings"
)

func HelloWorld(name string) string {

	var helloOutput string

	helloOutput = fmt.Sprintf("Hello, %s", name)

	return helloOutput

}

func HelloDaerah(name string, gender string, city string) string {

	var helloOutput string

	if strings.ToTitle(city) == strings.ToTitle("Jakarta") {
		if strings.ToTitle(gender) == strings.ToTitle("Pria") {
			helloOutput = fmt.Sprintf("Hi Mr. %s", name)
		} else {
			helloOutput = fmt.Sprintf("Hi Mrs. %s", name)
		}
	} else {
		if strings.ToTitle(city) == strings.ToTitle("Bandung") {
			if strings.ToTitle(gender) == strings.ToTitle("Pria") {
				helloOutput = fmt.Sprintf("Wilujeng Mr. %s", name)
			} else {
				helloOutput = fmt.Sprintf("Wilujeng Mrs. %s", name)
			}
		} else {
			if strings.ToTitle(city) == strings.ToTitle("Medan") {
				if strings.ToTitle(gender) == strings.ToTitle("Pria") {
					helloOutput = fmt.Sprintf("Horas Mr. %s", name)
				} else {
					helloOutput = fmt.Sprintf("Horas Mrs. %s", name)
				}
			}
		}
	}
	return helloOutput

}
