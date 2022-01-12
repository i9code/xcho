package xhttp

const (
	SerializerJson  Serializer = "json"
	SerializerProto Serializer = "proto"
	SerializerXml   Serializer = "xml"
	SerializerBytes Serializer = "bytes"
)

type Serializer string
