package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloDaerahRequest struct {
	NAME   string `json:"name"`
	GENDER string `json:"gender"`
	CITY   string `json:"city"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

type HelloDaerahResponse struct {
	STATUS  int    `json:"code`
	MESSAGE string `json:"message"`
}
