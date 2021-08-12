
type {{ $.WireGRPCTypeName }} *grpc.ClientConn
{{ if $.HasHTTPRule }}
type {{ $.WireHTTPTypeName }} *http.Client
{{ end }}

func {{ $.ProviderName }}Client(client {{ $.WireGRPCTypeName }}) {{ $.Name }}Client {
	return New{{ $.Name }}Client((*grpc.ClientConn)(client))
}

{{ if $.HasHTTPRule }}
func {{ $.ProviderName }}HTTPClient(client {{ $.WireHTTPTypeName }}) {{ $.Name }}HTTPClient {
	return New{{ $.Name }}HTTPClient((*http.Client)(client))
}
{{ end }}