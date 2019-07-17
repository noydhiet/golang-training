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

func HelloDaerah(name string, jenis_kelamin string, city string) string {
	var helloOutput string
	switch strings.ToUpper(city) {
	case "JAKARTA":
		{
			if strings.ToUpper(jenis_kelamin) == strings.ToUpper("PRIA") {
				helloOutput = fmt.Sprintf("Hi MR %s", name)
			} else {
				helloOutput = fmt.Sprintf("Hi MRS %s", name)
			}
		}
	case "BANDUNG":
		{
			if strings.ToUpper(jenis_kelamin) == strings.ToUpper("PRIA") {
				helloOutput = fmt.Sprintf("Wilujeng MR %s", name)
			} else {
				helloOutput = fmt.Sprintf("Wilujeng MRS %s", name)
			}
		}

	case "MEDAN":
		{
			if strings.ToUpper(jenis_kelamin) == strings.ToUpper("PRIA") {
				helloOutput = fmt.Sprintf("Horas MR %s", name)
			} else {
				helloOutput = fmt.Sprintf("Horas MRS %s", name)
			}
		}

	}
	return helloOutput
}
