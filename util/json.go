package util

import "encoding/json"

func UnmarshalFromInterface(value interface{}) (any, error) {

	switch v := value.(type) {
	case string:
		return map[string]interface{}{"String": v}, nil
	case int, int8, int16, int32, int64:
		return map[string]interface{}{"Int": v}, nil
	case uint, uint8, uint16, uint32, uint64:
		return map[string]interface{}{"Unsigned Int": v}, nil
	case float32, float64:
		return map[string]interface{}{"Float": v}, nil
	case bool:
		return map[string]interface{}{"Bool": v}, nil
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, nil
	}

	var parsedValue map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &parsedValue)
	if err != nil {
		return nil, err
	}
	return parsedValue, nil

}
