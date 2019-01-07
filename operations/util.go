package operations

import "github.com/rajnikant12345/kmip_g/parser"

func PrepareResponseMessage() *parser.KmipStructResponse {
	p := parser.KmipStructResponse{}
	p.ResponseMessage = &parser.ResponseMessage{}
	p.ResponseMessage.ResponseHeader = &parser.ResponseHeader{}
	return nil
}