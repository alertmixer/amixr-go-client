package amixr

type Config struct {
	Token string
}

func (c *Config) Client() (interface{}, error) {
	client, err := NewClient(c.Token)
	return client, err
}
