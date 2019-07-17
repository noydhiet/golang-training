package service

import (
	 "fmt"
	 "strings"
)

//Hello World Function
func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}

//Hello Daerah Function
func HelloDaerah(name string, gender string, city string) string {
	var helloOutput string
	switch strings.ToUpper(city) {
	case "JAKARTA":
		{
			if strings.ToUpper(gender) == "PRIA" {
				helloOutput = fmt.Sprintf("HI MR %s",name )
			} else {
				helloOutput = fmt.Sprintf("HI MRS %s",name )
			}
		}
	case "BANDUNG": 
		{
			if strings.ToUpper(gender) == "PRIA" {
				helloOutput = fmt.Sprintf("WILUJENG MR %s",name )
			} else {
				helloOutput = fmt.Sprintf("WILUJENG MRS %s",name )
			}
		}
	case "MEDAN":
		{
			if strings.ToUpper(gender) == "PRIA" {
				helloOutput = fmt.Sprintf("HORAS MR %s",name )
			} else {
				helloOutput = fmt.Sprintf("HORAS MRS %s",name )
			}
		}
	}
	return helloOutput
	
}
