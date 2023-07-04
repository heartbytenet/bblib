package reflection

import (
	"encoding/json"
)

func Convert(src any, dst any) (err error) {
	var (
		data []byte
	)

	data, err = json.Marshal(src)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, dst)
	if err != nil {
		return
	}

	return
}
