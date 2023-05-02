package go_mautic

type ClientConfig struct {
	baseUrl  string
	user     string
	password string
}

func Config() *ClientConfig {
	return &ClientConfig{}
}

func (config *ClientConfig) SetBaseUrl(baseUrl string) *ClientConfig {
	config.baseUrl = baseUrl

	return config
}

func (config *ClientConfig) SetUser(user string) *ClientConfig {
	config.user = user

	return config
}

func (config *ClientConfig) SetPassword(password string) *ClientConfig {
	config.password = password

	return config
}
