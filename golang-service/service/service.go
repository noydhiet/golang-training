package service

import "fmt"

func HelloWorld(name string) string {

	var helloOutput string

	helloOutput = fmt.Sprintf("Hello, %s", name)

	return helloOutput

}

func HelloDaerah(name, jenisKelamin, asalKota string) string {

	var daerahOutput string

	if jenisKelamin == "wanita" || jenisKelamin == "Wanita" || jenisKelamin == "WANITA" || jenisKelamin == "w" || jenisKelamin == "W" {
		jenisKelamin = "MRS."
	} else if jenisKelamin == "pria" || jenisKelamin == "Pria" || jenisKelamin == "PRIA" || jenisKelamin == "p" || jenisKelamin == "P" {
		jenisKelamin = "MR."
	} else {
		jenisKelamin = "MR/MRS."
	}

	if asalKota == "jakarta" || asalKota == "Jakarta" || asalKota == "JAKARTA" {
		asalKota = "Hi"
	} else if asalKota == "bandung" || asalKota == "Bandung" || asalKota == "BANDUNG" {
		asalKota = "Wilujeung"
	} else if asalKota == "medan" || asalKota == "Medan" || asalKota == "MEDAN" {
		asalKota = "Horas"
	} else {
		asalKota = "Nowhere"
	}

	daerahOutput = fmt.Sprintf("%s %s %s", asalKota, jenisKelamin, name)

	return daerahOutput
}
