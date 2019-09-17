package factom

import ()

const (
	Endpoint = "https://api.factom.pro"
	Version  = "v1"
)

type Factom interface {
	GetAPIInfo() interface{}
}

type Context struct {
	apiKey string
}

func NewFactom(apiKey string) Factom {

	return &Context{apiKey}

}
