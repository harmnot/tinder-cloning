package driver

type DBConf struct {
	Driver         string `json:"driver"`
	Host           string `json:"host"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Port           int    `json:"port"`
	SSL            string `json:"ssl"`
	User           string `json:"user"`
	ConMaxLifetime int    `json:"con_max_lifetime" split_words:"true"`
	ConMaxIdle     int    `json:"con_max_idle" split_words:"true"`
	ConMaxOpen     int    `json:"con_max_open" split_words:"true"`
}

type GlobalConf struct {
	DB DBConf
}
