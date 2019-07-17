package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}
type HelloDaerahRequest struct {
	NAME string `json:"name"`
	GENDER string `json:"jenis_kelamin"`
	KOTA string `json:"asal_kota"`
}

type HelloDaerahResponse struct {
	MESSAGE string `json:"message"`
}