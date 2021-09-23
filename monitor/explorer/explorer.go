package explorer

type Explorer struct {
	Name              string            `json:"name"`
	Coin              string            `json:"coin"`
	Url               string            `json:"url"`
	HeightJsonPattern string            `json:"heightJsonPattern"`
	HashJsonPattern   string            `json:"hashJsonPattern"`
	CustomHeaders     map[string]string `json:"customHeaders"`
	Enabled           bool              `json:"enabled"`
}
