package controllers

type Context interface {
	Bind(interface{}) error
	Get(string) (interface{}, bool)
	JSON(int, interface{})
	Param(string) string
	Status(int)
}
