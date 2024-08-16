package betterj

import (
	"encoding/json"
)

func MinifyJ(s string) (string, error) {
	var jsonobj interface{}
	err := json.Unmarshal([]byte(s), &jsonobj)
	if err != nil {
		return "", err
	}
	marshalled, err := json.Marshal(jsonobj)
	if err != nil {
		return "", err
	}
	return string(marshalled), nil

}

func BeautifyJ(s string, indentWith string) (string, error) {
	var jsonobj interface{}
	err := json.Unmarshal([]byte(s), &jsonobj)
	if err != nil {
		return "", err
	}

	marshalled, err := json.MarshalIndent(jsonobj, " ", indentWith)
	if err != nil {
		return "", err
	}
	return string(marshalled), nil
}
