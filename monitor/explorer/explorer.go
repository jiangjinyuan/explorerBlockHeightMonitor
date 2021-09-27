package explorer

type Explorer struct {
	Name              string            `json:"name"`
	Coin              string            `json:"coin"`
	URL               string            `json:"url"`
	HeightJSONPattern string            `json:"heightJsonPattern"`
	HashJSONPattern   string            `json:"hashJsonPattern"`
	CustomHeaders     map[string]string `json:"customHeaders"`
	Enabled           bool              `json:"enabled"`
}
