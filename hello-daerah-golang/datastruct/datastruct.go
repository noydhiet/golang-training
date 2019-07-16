package datastruct

type HelloDaerahRequest struct {
	NAME   string `json:"name"`
	SEX    string `json:"sex"`
	ORIGIN string `json:"origin_town"`
}

type HelloDaerahResponse struct {
	MESSAGE string `json:"message"`
}
