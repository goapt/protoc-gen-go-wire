package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"strings"
)

//go:embed template.go.tpl
var tpl string

type service struct {
	Name     string // Greeter
	FullName string // helloworld.Greeter
	FilePath string // api/helloworld/helloworld.proto
}

func (s *service) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("wire").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return buf.String()
}

// WireTypeName is wire var type name
func (s *service) WireGRPCTypeName() string {
	return "GrpcClientConn"
}

// WireTypeName is wire var type name
func (s *service) WireHTTPTypeName() string {
	return "HttpClientConn"
}

// ProviderName is wire provider func name
func (s *service) ProviderName() string {
	return "NewWire" + s.Name
}
