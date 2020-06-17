package model

type ProjectConf struct {
	DataBaseConfig DataBaseConfig `json:"data_base_config"`
	ServerConfig   ServerConfig   `json:"server_config"`
}

type DataBaseConfig struct {
	Type  string `json:"type"`
	Local string `json:"local"`
}

type ServerConfig struct {
	Port string `json:"port"`
	Flag int    `json:"flag"`
	Tag  int    `json:"tag"`
}
