package cache

import "encoding/json"

type JsonEncoder struct{}

func (val JsonEncoder) Dump(fromObj interface{}, toObj interface{}) error {
	fromBytes, err := val.Encode(fromObj)
	if err != nil {
		return err
	}
	return val.Decode(fromBytes, toObj)
}

func (val JsonEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (val JsonEncoder) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
