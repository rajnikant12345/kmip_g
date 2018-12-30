package parser

type RequestMessage struct {
	RequestHeader RequestHeader	`kmip:"420077"`
	//BatchItem     []BatchItem
}

type KmipStruct struct {
	RequestMessage RequestMessage `kmip:"420078"`
}

func (r *RequestMessage) Unmarshal(b []byte) ([]byte, error) {
	/*r.BatchItem = make([]BatchItem, 0)

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
			b, _ = r.RequestHeader.Unmarshal(b)
		case 0x42000F:
			b = b[8:]
			bat := BatchItem{}
			b, _ = bat.Unmarshal(b)
			r.BatchItem = append(r.BatchItem, bat)
		default:
			return b, nil

		}
	}*/
	return nil,nil

}
