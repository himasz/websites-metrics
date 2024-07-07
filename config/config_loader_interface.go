package config

type ILoader interface {
	Load(file string, v interface{}) error
}
