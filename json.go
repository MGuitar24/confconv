package confconv

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

func ConvertYAMLToJSON(data []byte) ([]byte, error) {
	var ymlData map[string]interface{}
	err := yaml.Unmarshal(data, &ymlData)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	return jsonData, err
}

func RewriteYAMLToJSON(path string) error {
	if fileExists(path) {
		fileData, err := readFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		jsonData, err := ConvertYAMLToJSON(fileData)
		if err != nil {
			return fmt.Errorf("failed to convert file to JSON %s: %w", path, err)
		}

		fileExt := identifyFileType(path)
		switch fileExt {
		case "yml":
			{
				jsonPath := strings.Replace(path, ".yml", ".json", 1)
				err = writeFile(jsonPath, jsonData)
				if err != nil {
					return fmt.Errorf("failed to write JSON file %s: %w", path, err)
				}
				return nil
			}
		case "yaml":
			{
				jsonPath := strings.Replace(path, ".yaml", ".json", 1)
				err = writeFile(jsonPath, jsonData)
				if err != nil {
					return fmt.Errorf("failed to write JSON file %s: %w", path, err)
				}
				return nil
			}
		default:
			{
				jsonPath := path + ".json"
				err = writeFile(jsonPath, jsonData)
				if err != nil {
					return fmt.Errorf("failed to write JSON file %s: %w", path, err)
				}
				return nil
			}
		}
	} else {
		return fmt.Errorf("file %s does not exist", path)
	}
}
