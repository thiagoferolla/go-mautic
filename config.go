package go_mautic

type ClientConfig struct {
	baseUrl  string
	user     string
	password string
}

func (config *ClientConfig) SetBaseUrl(baseUrl string) {
	config.baseUrl = baseUrl
}

func (config *ClientConfig) SetUser(user string) {
	config.user = user
}

func (config *ClientConfig) SetPassword(password string) {
	config.password = password
}
