package service

import "fmt"

func HelloDaerah(name, jenisKelamin, asalKota string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("%s %s %s", asalKota, jenisKelamin, name)

	return helloOutput
}
