package config

type EnvError struct {
	message string
	data    map[string]interface{}
}

func (err *EnvError) Error() string {
	return err.message
}
