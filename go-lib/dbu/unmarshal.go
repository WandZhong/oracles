package dbu

import (
	"encoding/json"
	"fmt"

	bat "github.com/robert-zaremba/go-bat"
)

// UnmarshalJSON reads json from target (which should be either []byte or string) and writes it to target
func UnmarshalJSON(source interface{}, target interface{}) error {
	switch t := source.(type) {
	default:
		return fmt.Errorf("Unable to parse %T as json", source)
	case []byte:
		return json.Unmarshal(t, target)
	case string:
		return json.Unmarshal(bat.UnsafeStrToByteArray(t), target)
	}
}

// UnmarshalString reads string from target (which should be either []byte or string) and writes it to target
func UnmarshalString(source interface{}, target *string) (err error) {
	if source == nil {
		*target = ""
		return nil
	}
	*target, err = ConvertToString(source)
	return
}

// ConvertToString converts interface{} to string
func ConvertToString(source interface{}) (string, error) {
	switch t := source.(type) {
	default:
		return "", fmt.Errorf("Unable to parse %T as string. ", source)
	case []byte:
		return bat.UnsafeByteArrayToStr(t), nil
	case string:
		return t, nil
	}
}

// ConvertToBytes converts interface{} to []byte
func ConvertToBytes(source interface{}) ([]byte, error) {
	switch t := source.(type) {
	default:
		return nil, fmt.Errorf("Unable to parse %T as string. ", source)
	case []byte:
		return t, nil
	case string:
		return bat.UnsafeStrToByteArray(t), nil
	}
}
