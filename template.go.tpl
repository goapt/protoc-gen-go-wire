
type {{ $.WireTypeName }} *grpc.ClientConn

func {{ $.ProviderName }}(client {{ $.WireTypeName }}) {{ $.Name }}Client {
	return New{{ $.Name }}Client((*grpc.ClientConn)(client))
}
