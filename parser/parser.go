package parser

import (
	"errors"

	"github.com/rickmoonex/gotigen/models"
)

// TypeMap represents a mapping between the ThingsDB type and the Golang type
var TypeMap = map[string]string{
	"str": "string",
	"#":   "int",
}

// ParseType takes a type defined in ThingsDB and parses it to a Golang type
func ParseType(source string) string {
	val, ok := TypeMap[source]
	if !ok {
		return ""
	}
	return val
}

// ParseTypeInfoResponse parses the results of a type info query into a ParsedType object
func ParseTypeInfoResponse(response interface{}) (*models.ParsedType, error) {
	parsedType := &models.ParsedType{}

	resMap, ok := response.(map[string]interface{})
	if !ok {
		return nil, errors.New("error parsing response to map[string]interface{}")
	}

	// Parsing of the name
	nameValue, ok := resMap["name"].(string)
	if !ok {
		return nil, errors.New("`name` property on res is not a string")
	}

	parsedType.Name = nameValue

	fieldsValue, ok := resMap["fields"].([]interface{})
	if !ok {
		return nil, errors.New("`fields` property cannot be parsed")
	}

	for _, e := range fieldsValue {
		field := models.ParsedTypeField{}

		typedE, ok := e.([]interface{})
		if !ok {
			return nil, errors.New("field cannot be parsed to string array")
		}

		field.Name = typedE[0].(string)
		field.Type = ParseType(typedE[1].(string))

		parsedType.Fields = append(parsedType.Fields, field)
	}

	return parsedType, nil
}
