package challenge

type DBConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Password    string `json:"-"`
	MaxIdleTime *int   `json:"max_idle_time,omitempty"`
	UseSSL      bool   `json:"ssl_enabled"`
}
