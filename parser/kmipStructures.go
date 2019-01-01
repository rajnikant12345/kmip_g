package parser

type RequestMessage struct {
	RequestHeader *RequestHeader `kmip:"420077"`
	BatchItem     []*BatchItem   `kmip:"42000F"`
}

type KmipStruct struct {
	RequestMessage *RequestMessage `kmip:"420078"`
	ResponseMessage *ResponseMessage `kmip:"42007B"`
}

type KmipStructResponse struct {
	ResponseMessage *ResponseMessage `kmip:"42007B"`
}

type ResponseMessage struct {
	ResponseHeader *ResponseHeader `kmip:"42007A"`
	BatchItem     []*BatchItem   `kmip:"42000F"`
}
