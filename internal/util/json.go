package util
import "encoding/json"

func BytesToJson(bytes []byte) (any, error) {
	var err error

	// Convert JSON bytes to map
	var result any
	if err = json.Unmarshal(bytes, &result); err != nil {
		return nil, nil
	}


	return result, nil
}