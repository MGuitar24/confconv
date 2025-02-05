package confconv

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

func ConvertJSONToYML(data []byte) ([]byte, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}
	ymlData, err := yaml.Marshal(jsonData)
	return ymlData, err
}
