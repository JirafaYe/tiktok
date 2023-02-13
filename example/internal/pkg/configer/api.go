package configer

var handler *Manager

func init() {
	handler = New()
}

func ReadConfig(c Config) error {
	return handler.ReadConfig(c)
}
