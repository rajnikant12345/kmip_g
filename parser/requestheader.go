package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type RequestHeader struct {
	ProtocolVersion             *ProtocolVersion
	MaximumResponseSize         *kmipbin.KmipInt
	AsynchronousIndicator       *kmipbin.KmipBoolean
	AttestationCapableIndicator *kmipbin.KmipBoolean
	AttestationType             *kmipbin.KmipEnum
	Authentication              *Authentication
	BatchCount                  *kmipbin.KmipInt
}

func (r *RequestHeader) Unmarshal(b []byte) ([]byte, error) {

	for {
		if len(b) <= 0 {
			return b, nil
		}
		tag := kmipbin.KmipTagType{}
		tag.UnMarshalBin(b[:4])
		switch int(tag.Tag) {
		case 0x420069:
			b = b[8:]
			r.ProtocolVersion = &ProtocolVersion{}
			b, _ = r.ProtocolVersion.Unmarshal(b)
		case 0x42000D:
			b = b[8:]
			var bc kmipbin.KmipInt
			r.BatchCount = &bc
			r.BatchCount.UnMarshalBin(b[:8])
			b = b[8:]
		default:
			return b, nil

		}
	}
}
