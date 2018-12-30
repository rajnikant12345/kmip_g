package parser

import "github.com/rajnikant12345/kmip_g/kmipbin"

type BatchItem struct {
	Operation         *kmipbin.KmipEnum       `kmip:"42005C"`
	UniqueBatchItemID *kmipbin.KmipByteString `kmip:"420093"`
	RequestPayload    *RequestPayload         `kmip:"42007C"`
	MessageExtension  *MessageExtension       `kmip:"420051"`
}

func (r *BatchItem) Unmarshal(b []byte) ([]byte, error) {

	for {
		if len(b) <= 0 {
			return b, nil
		}
		tag := kmipbin.KmipTagType{}
		tag.UnMarshalBin(b[:4])
		switch int(tag.Tag) {
		case 0x420079:

			var l kmipbin.KmipLength
			l.UnMarshalBin(b[4:8])
			b = b[8+int(l):]
			//r.RequestPayload = RequestPayload{}
			//b,_ = r.RequestPayload.Unmarshal(b)

		case 0x42005C:

			b = b[8:]
			var op kmipbin.KmipEnum
			r.Operation = &op
			r.Operation.UnMarshalBin(b[:8])
			b = b[8:]

		case 0x420093:

			var l kmipbin.KmipLength
			l.UnMarshalBin(b[4:8])
			b = b[8:]
			le := kmipbin.PadLength(int(l))
			var uvi kmipbin.KmipByteString
			r.UniqueBatchItemID = &uvi
			r.UniqueBatchItemID.UnMarshalBin(b[:le], int(l))
			b = b[le:]

		case 0x420051:

			var l kmipbin.KmipLength
			l.UnMarshalBin(b[4:8])
			b = b[8+int(l):]

		default:
			return b, nil

		}
	}
}
