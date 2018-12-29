package parser

type RequestPayload struct {
}

func (r *RequestPayload) Unmarshal(b []byte) ([]byte, error) {

	return b, nil
}
