package service

import (
	"fmt"
	"strings"
)

func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}

func HelloDaerah(name, jenis_kelamin, asal_kota string) string {
	var helloOutput string
	if strings.ToTitle(asal_kota) == strings.ToTitle("Jakarta") {
		if strings.ToTitle(jenis_kelamin) == strings.ToTitle("Pria") {
			helloOutput = fmt.Sprintf("Hi Mr. %s", name)
		} else {
			helloOutput = fmt.Sprintf("Hi Mrs. %s", name)
		}
	} else {
		if strings.ToTitle(asal_kota) == strings.ToTitle("Bandung") {
			if strings.ToTitle(jenis_kelamin) == strings.ToTitle("Pria") {
				helloOutput = fmt.Sprintf("Wilujeng Mr. %s", name)
			} else {
				helloOutput = fmt.Sprintf("Wilujeng Mrs. %s", name)
			}
		} else {
			if strings.ToTitle(asal_kota) == strings.ToTitle("Medan") {
				if strings.ToTitle(jenis_kelamin) == strings.ToTitle("Pria") {
					helloOutput = fmt.Sprintf("Horas Mr. %s", name)
				} else {
					helloOutput = fmt.Sprintf("Horas Mrs. %s", name)
				}
			}
		}
	}
	return helloOutput
}
