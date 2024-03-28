package common

type Reponse struct {
	Data     interface{} `json:"data"`
	Paggings interface{} `json:"paggings"`
}

func ReponseData(data interface{}) *Reponse {
	return &Reponse{Data: data, Paggings: nil}
}
func MutiResponse(data, paggings interface{}) *Reponse {
	return &Reponse{Data: data, Paggings: paggings}
}
