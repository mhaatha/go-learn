package challenge

type Config struct {
	ServerIP       string `json:"host"`
	Port           int    `json:"port"`
	Timeout        int    `json:"timeout,omitempty"`
	DatabasePass   string `json:"-"`
	EnableFeatureX *bool  `json:"enable_feature_x,omitempty"`
}
