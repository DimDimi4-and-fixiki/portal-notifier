package entity

type JSONSerializable any

type HttpResp[DataT JSONSerializable] struct {
	Message string `json:"message"`
	Result  DataT  `json:"result"`
	Error   string `json:"error"`
	Meta    string `json:"meta"`
}
