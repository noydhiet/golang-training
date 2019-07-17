package datastruct

type HelloDaerahRequest struct {
	NAME          string `json:"name`
	JENIS_KELAMIN string `json:"jenis_kelamin"`
	ASAL_KOTA     string `json:"asal_kota"`
}

type HelloWorldResponse struct {
	STATUS  int    `json:"code"`
	MESSAGE string `json:"message"`
}

type HelloDaerahResponse struct {
	STATUS  int    `json:"code`
	MESSAGE string `json:"message"`
}
