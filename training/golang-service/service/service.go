package service

import (
	"fmt"
)

func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}

func HelloDaerah(name, jenis_kelamin, asal_kota string) string {
	var helloOutput string
	if (jenis_kelamin == "wanita") || (jenis_kelamin == "Wanita") || (jenis_kelamin == "WANITA") || (jenis_kelamin == "w") || (jenis_kelamin == "W") {
		jenis_kelamin = "MRS."
	} else if (jenis_kelamin == "pria") || (jenis_kelamin == "Pria") || (jenis_kelamin == "PRIA") || (jenis_kelamin == "p") || (jenis_kelamin == "P") {
		jenis_kelamin = "MR."
	} else {
		jenis_kelamin = "MR/MRS."
	}

	if (asal_kota == "jakarta") || (asal_kota == "Jakarta") || (asal_kota == "JAKARTA") {
		asal_kota = "Hi"
	} else if (asal_kota == "bandung") || (asal_kota == "Bandung") || (asal_kota == "BANDUNG") {
		asal_kota = "Wilujeung"
	} else if (asal_kota == "medan") || (asal_kota == "Medan") || (asal_kota == "MEDAN") {
		asal_kota = "Horas"
	} else {
		asal_kota = "Gagal"
	}

	helloOutput = fmt.Sprintf("%s %s %s", asal_kota, jenis_kelamin, name)
	return helloOutput
}
