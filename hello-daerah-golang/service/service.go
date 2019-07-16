package service

import (
	fmt "fmt"
)

func HelloDaerah(name, sex, origin string) string {
	var helloOutput string

	if origin == "Jakarta" || origin == "jakarta" || origin == "JAKARTA" || origin == "JKT" || origin == "jkt" {
		if sex == "Pria" || sex == "pria" || sex == "PRIA" {
		} else if sex == "Wanita" || sex == "wanita" || sex == "WANITA" {
			helloOutput = fmt.Sprintf("Hi Ms. %s", name)
		}
	} else if origin == "Bandung" || origin == "bandung" || origin == "BANDUNG" || origin == "BDG" || origin == "bdg" {
		if sex == "Pria" || sex == "pria" || sex == "PRIA" {
			helloOutput = fmt.Sprintf("Wilujeung Mr. %s", name)
		} else if sex == "Wanita" || sex == "wanita" || sex == "WANITA" {
			helloOutput = fmt.Sprintf("Wilujeung Ms. %s", name)
		}
	} else if origin == "Medan" || origin == "medan" || origin == "MEDAN" || origin == "MDN" || origin == "mdn" {
		if sex == "Pria" || sex == "pria" || sex == "PRIA" {
			helloOutput = fmt.Sprintf("Horas Mr. %s", name)
		} else if sex == "Wanita" || sex == "wanita" || sex == "WANITA" {
			helloOutput = fmt.Sprintf("Horas Ms. %s", name)
		}
	} else {
		helloOutput = fmt.Sprintf("Data yang di masukkan salah")
	}
	return helloOutput
}
