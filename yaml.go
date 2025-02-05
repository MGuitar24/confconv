package confconv

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func ConvertJSONToYAML(data []byte) ([]byte, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}
	ymlData, err := yaml.Marshal(jsonData)
	return ymlData, err
}

func RewriteJSONToYAML(path string) error {
	if fileExists(path) {
		fileData, err := readFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		jsonData, err := ConvertJSONToYAML(fileData)
		if err != nil {
			return fmt.Errorf("failed to convert file to JSON %s: %w", path, err)
		}

		fileExt := identifyFileType(path)
		switch fileExt {
		case "json":
			{
				jsonPath := strings.Replace(path, ".json", ".yaml", 1)
				err = writeFile(jsonPath, jsonData)
				if err != nil {
					return fmt.Errorf("failed to write YAML file %s: %w", path, err)
				}
				return nil
			}
		default:
			{
				jsonPath := path + ".yaml"
				err = writeFile(jsonPath, jsonData)
				if err != nil {
					return fmt.Errorf("failed to write YAML file %s: %w", path, err)
				}
				return nil
			}
		}
	} else {
		return fmt.Errorf("file %s does not exist", path)
	}
}
