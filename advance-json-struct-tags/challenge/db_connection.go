package challenge

type Option func(*DBConfig)

func NewDB(host, password string, opts ...Option) *DBConfig {
	config := &DBConfig{
		Host:     host,
		Password: password,
		Port:     5432,
		UseSSL:   true,
	}

	for _, opt := range opts {
		opt(config)
	}

	return config
}

func WithPort(p int) func(*DBConfig) {
	return func(d *DBConfig) {
		d.Port = p
	}
}

func WithMaxIdleTime(t int) func(*DBConfig) {
	return func(d *DBConfig) {
		d.MaxIdleTime = &t
	}
}

func WithSSL(active bool) func(*DBConfig) {
	return func(d *DBConfig) {
		d.UseSSL = active
	}
}
