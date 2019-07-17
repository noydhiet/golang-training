package service

import (
	"strings"
	 "fmt"
)

func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}
func HelloDaerah(name string, jenis_kelamin string, asal_kota string ) string {
	var helloOutput string
	switch strings.ToUpper(asal_kota) {
	case "JAKARTA":
		{
			if strings.ToUpper(jenis_kelamin) == "MALE"{
				helloOutput = fmt.Sprintf("Hai Mr %s", name)
			}else {
				helloOutput = fmt.Sprintf("Hai Mrs %s", name)
			}
		}
	case "BANDUNG":
		{
			if strings.ToUpper(jenis_kelamin) == "MALE"{
				helloOutput = fmt.Sprintf("Wilujeng Mr %s", name)
			}else {
				helloOutput = fmt.Sprintf("Wilujeng Mrs %s", name)
			}
		}
	case "MEDAN":
		{
			if strings.ToUpper(jenis_kelamin) == "MALE"{
				helloOutput = fmt.Sprintf("Horas Mr %s", name)
			}else {
				helloOutput = fmt.Sprintf("Horas Mrs %s", name)
			}
		}
	}
	// helloOutput = fmt.Sprintf("Hello, %s %s %s", name, jenis_kelamin, asal_kota)
	return helloOutput
}
