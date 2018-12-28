package parser

import (
	"github.com/rajnikant12345/kmip_g/kmipbin"
)

type ProtocolVersion struct {
	ProtocolVersionMajor kmipbin.KmipInt
	ProtocolVersionMinor kmipbin.KmipInt
}

func (r *ProtocolVersion) Unmarshal(b []byte) ([]byte,error) {
	for {
		if len(b) <= 0 {
			break
		}
		tag := kmipbin.KmipTagType{}
		tag.UnMarshalBin(b[:4])
		switch int(tag.Tag) {
		case 0x42006B:
			b = b[8:]
			r.ProtocolVersionMinor.UnMarshalBin(b[:8])
			b = b[8:]
		case 0x42006A:
			b = b[8:]
			r.ProtocolVersionMajor.UnMarshalBin(b[:8])
			b = b[8:]
		default:
			return b, nil
		}
	}
	return b,nil
}

type RequestHeader struct {
	ProtocolVersion *ProtocolVersion
	BatchCount kmipbin.KmipInt
}

func (r *RequestHeader) Unmarshal(b []byte) ([]byte,error) {

	for {
		if len(b) <= 0 {
			return b,nil
		}
		tag := kmipbin.KmipTagType{}
		tag.UnMarshalBin(b[:4])
		switch int(tag.Tag) {
		case 0x420069:
			b = b[8:]
			r.ProtocolVersion = &ProtocolVersion{}
			b,_ = r.ProtocolVersion.Unmarshal(b)
		case 0x42000D:
			b = b[8:]
			r.BatchCount.UnMarshalBin(b[:8])
			b = b[8:]
		default:
			return b, nil

		}
	}

	return b,nil
}

type RequestPayload struct {

}

func (r *RequestPayload) Unmarshal(b []byte) ([]byte,error) {

	return b,nil
}


type BatchItem struct {
	Operation kmipbin.KmipEnum
	RequestPayload RequestPayload
}

func (r *BatchItem) Unmarshal(b []byte) ([]byte,error) {

	return b,nil
}

type RequestMessage struct {
	RequestHeader *RequestHeader
	BatchItem    []BatchItem
}



func (r *RequestMessage) Unmarshal(b []byte) ([]byte,error) {
	r.BatchItem = make([]BatchItem, 0)
	for {
		if len(b) <= 0 {
			return b, nil
		}
		tag := kmipbin.KmipTagType{}
		tag.UnMarshalBin(b[:4])
		switch int(tag.Tag) {
		case 0x420077:
			b = b[8:]
			r.RequestHeader = &RequestHeader{}
			b,_ = r.RequestHeader.Unmarshal(b)
		case 0x42000F:
			var l kmipbin.KmipLength
			l.UnMarshalBin(b[4:8])
			b = b[8:]
			bat := BatchItem{}
			bat.Unmarshal(b)
			r.BatchItem = append(r.BatchItem,bat)
		default:
			return b, nil

		}
	}

	return b, nil
}
