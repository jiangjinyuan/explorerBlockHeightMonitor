package explorer

type Explorer struct {
	Name          string            `json:"name"`
	Coin          string            `json:"coin"`
	Url           string            `json:"url"`
	JsonPattern   string            `json:"jsonPattern"`
	CustomHeaders map[string]string `json:"customHeaders"`
	Enabled       bool              `json:"enabled"`
}
