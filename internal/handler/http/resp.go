package http

type JSONSerializable any

type Resp[DataT JSONSerializable] struct {
	Message string `json:"message"`
	Result  DataT  `json:"result"`
	Meta    string `json:"meta"`
}

// ErrForResp Info about error for response
type ErrForResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RespErr Http Response when error happened
type RespErr struct {
	Err  ErrForResp `json:"error"`
	Meta string     `json:"meta"`
}

func (r RespErr) WithMeta(meta string) RespErr {
	r.Meta = meta
	return r
}

func RespFromErr(err Err) RespErr {
	return RespErr{
		Err: ErrForResp{
			Code:    err.code,
			Message: err.Error(),
		},
	}
}

func RespValidationErr(err error) RespErr {
	vError := ValidationError(err)
	return RespFromErr(vError)

}
