package web

import "fmt"

type DbMeta struct {
	Id          string
	Name        string
	Kind        string
	Service     string
	Version     string
	Tag         string
	Credentials string
	Config      string
	Schema      string
	Dictionary  string //dictionary data
	Data        string //use case static data
	Sequence    bool
	Tables      []string
}

type AppMeta struct {
	Name         string
	Description  string
	OriginURL    string
	Sdk          string
	Docker       bool
	Args         []string
	Config       string
	Build        string //build path
	DbConfigPath string
	Assets       []string
	Selenium     map[string]interface{}
	HTTP         map[string]interface{}
	REST         map[string]interface{}
}

func (m *AppMeta) GetArguments(dockerfile bool) string {
	var result = ""
	if len(m.Args) == 0 {
		return result
	}
	for _, item := range m.Args {
		if len(result) > 0 || dockerfile {
			result += ","
		}
		if dockerfile {
			result += fmt.Sprintf(`"%v"`, item)
			continue
		}
		result += fmt.Sprintf(`%v`, item)
	}
	return result
}
