
type {{ $.WireGRPCTypeName }} *grpc.ClientConn
type {{ $.WireHTTPTypeName }} *http.Client

func {{ $.ProviderName }}Client(client {{ $.WireGRPCTypeName }}) {{ $.Name }}Client {
	return New{{ $.Name }}Client((*grpc.ClientConn)(client))
}

func {{ $.ProviderName }}HTTPClient(client {{ $.WireHTTPTypeName }}) {{ $.Name }}HTTPClient {
	return New{{ $.Name }}HTTPClient((*http.Client)(client))
}