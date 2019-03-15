package mapper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/structs"
)

// ShallowMapStructByTag takes a struct and target tag name present on fields in that struct,
// then converts it into a map[string]interface{}. The target tag should be of the format `myTag:"destinationFieldName"`,
// where `destinationFieldName` is a valid map[string] key.
func ShallowMapStructByTag(src interface{}, tagName string) (map[string]interface{}, error) {
	return mapStructByTag(src, tagName, "root", false)
}

func DeepMapStructByTag(src interface{}, tagName string) (map[string]interface{}, error) {
	return mapStructByTag(src, tagName, "root", true)
}

func mapStructByTag(src interface{}, tagName string, parentFieldName string, deepMapping bool) (map[string]interface{}, error) {
	if src == nil || !structs.IsStruct(src) {
		return nil, errors.New("Source cannot be nil and must be a struct")
	}

	result := make(map[string]interface{})

	for _, sourceField := range structs.Fields(src) {
		fieldPath := fmt.Sprintf("%s.%s", parentFieldName, sourceField.Name())
		tag := sourceField.Tag(tagName)
		if tag == "" {
			continue
		}

		options := strings.Split(tag, ",")
		if len(options) < 1 {
			continue
		}

		destinationFieldName := options[0]
		sourceValue := sourceField.Value()

		if deepMapping {
			if structs.IsStruct(sourceValue) {
				var err error
				sourceValue, err = mapStructByTag(sourceValue, tagName, fieldPath, deepMapping)

				if err != nil {
					return result, fmt.Errorf("Failed to marshal %s: %s", sourceField.Name(), err)
				}

			}
		}

		result[destinationFieldName] = sourceValue
	}

	return result, nil
}
