package util

import "encoding/json"

func PrintToString(in any) string {
	b, err := json.MarshalIndent(in, "", " ")
	if err != nil {
		return err.Error()
	}

	return string(b)
}
