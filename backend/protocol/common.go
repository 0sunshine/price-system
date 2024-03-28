package protocol

type CommonFilterBase struct {
	PageSize    int `json:"page-size"`
	CurrentPage int `json:"current-page"`
}

type CommonRequestBase[T interface{}] struct {
	Token  string `json:"token"`
	Filter T      `json:"filter"`
}

type CommonRespondse[T interface{}] struct {
	Code int `json:"code"`
	Msg  int `json:"msg"`
	Data T   `json:"data"`
}

type NoRespondseData struct {
}

type CommonNoDataRespondse CommonRespondse[NoRespondseData]
