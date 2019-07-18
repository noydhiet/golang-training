package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

type HelloDaerahRequest struct {
	NAME          string `json:"name"`
	JENIS_KELAMIN string `json:"jenis_kelamin"`
	ASAL_KOTA     string `json:"asal_kota"`
}

type HelloDaerahResponse struct {
	
	MESSAGE string `json:"message"`
}
