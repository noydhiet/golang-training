package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

//HelloDaerah
type HelloDaerahRequest struct {
	NAME   string `json:"name"`
	GENDER string `json:"gender"`
	CITY   string `json:"city"`
}

type HelloDaerahResponse struct {
	MESSAGE string `json:"message"`
}
