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

func HelloDaerah(name string, jenis_kelamin string, asal_kota string) string {
	var helloDaerahOutput string
	// sapaan := ""
	// panggilan := ""

	switch strings.ToUpper(asal_kota) {
	case "JAKARTA":
		{
			if strings.ToUpper(jenis_kelamin) == "PRIA" {
				helloDaerahOutput = fmt.Sprintf("Hi, Mr. %s ", name)
			} else {
				helloDaerahOutput = fmt.Sprintf("Hi, Mrs. %s ", name)
			}
		}
	case "BANDUNG":
		{
			if strings.ToUpper(jenis_kelamin) == "PRIA" {
				helloDaerahOutput = fmt.Sprintf("Wilujeng, Mr. %s ", name)
			} else {
				helloDaerahOutput = fmt.Sprintf("Wilujeng, Mrs. %s ", name)
			}
		}
	case "MEDAN":
		{
			if strings.ToUpper(jenis_kelamin) == "PRIA" {
				helloDaerahOutput = fmt.Sprintf("Horas, Mr. %s ", name)
			} else {
				helloDaerahOutput = fmt.Sprintf("Horas, Mrs. %s ", name)
			}
		}

	}

	// if strings.ToLower(asal_kota) == "Jakarta" {
	// 	sapaan = "Hi"
	// } else if strings.ToLower(asal_kota) == "Bandung" {
	// 	sapaan = "Wilujeng"
	// } else {
	// 	sapaan = "Horas"
	// }

	// if strings.ToLower(jenis_kelamin) == "PRIA" {
	// 	panggilan = "MR."
	// } else {
	// 	panggilan = "MRS."
	// }

	// helloDaerahOutput = fmt.Sprintf("%s ", sapaan, name, panggilan)

	return helloDaerahOutput
}
