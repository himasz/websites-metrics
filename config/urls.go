package config

type URLConfig struct {
	URL   string `json:"url"`
	Regex string `json:"regex"`
}

type URLsConfig struct {
	URLs []URLConfig `json:"urls"`
}
