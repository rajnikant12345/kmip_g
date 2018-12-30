package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type ProtocolVersion struct {
	ProtocolVersionMajor kmipbin.KmipInt	`kmip:"42006A"`
	ProtocolVersionMinor kmipbin.KmipInt	`kmip:"42006B"`
}
/*
func (r *ProtocolVersion) Unmarshal(b []byte) ([]byte, error) {
	for {
		if len(b) <= 0 {
			break
		}
		tag := kmipbin.KmipTagType{}
		tag.UnMarshalBin(b[:4])
		switch int(tag.Tag) {
		case 0x42006B:
			b = b[8:]
			var min kmipbin.KmipInt
			r.ProtocolVersionMinor = &min
			r.ProtocolVersionMinor.UnMarshalBin(b[:8])
			b = b[8:]
		case 0x42006A:
			b = b[8:]
			var maj kmipbin.KmipInt
			r.ProtocolVersionMajor = &maj
			r.ProtocolVersionMajor.UnMarshalBin(b[:8])
			b = b[8:]
		default:
			return b, nil
		}
	}
	return b, nil
}
*/