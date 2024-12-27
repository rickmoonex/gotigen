package parser

import (
	"github.com/rickmoonex/gotigen/internal/models"
)

// defaultType is the type that is used when no other can be matched
const defaultType = "interface{}"

// typeMap is used to map ThingsDB types to Golang types
var typeMap = map[string]string{
	"#":   "int",
	"str": "string",
}

// parseType takes in a ThingsDB type and returns its correct Golang type
// If no type can be matched it defaults to defaultType.
func parseType(sourceType string) string {
	val, ok := typeMap[sourceType]
	if !ok {
		return defaultType
	}
	return val
}

// ParseThingsDBTypes parses an array of ThingsDBType into an array of ParsedType
func ParseThingsDBTypes(sourceTypes []models.ThingsDBType) *[]models.ParsedType {
	var finalTypes []models.ParsedType

	for _, v := range sourceTypes {
		parsedType := models.ParsedType{}

		parsedType.Name = v.Name
		parsedType.Fields = parseThingsDBFields(v.Fields)
		parsedType.Methods = parseThingsDBMethods(v.Methods)

		finalTypes = append(finalTypes, parsedType)
	}

	return &finalTypes
}

// parseThingsDBFields parses an array of ThingsDBTypeField into an array of ParsedTypeField
func parseThingsDBFields(sourceFields []models.ThingsDBTypeField) []models.ParsedTypeField {
	var finalFields []models.ParsedTypeField

	for _, v := range sourceFields {
		parsedField := models.ParsedTypeField{}

		parsedField.Name = v[0]
		parsedField.Type = parseType(v[1])

		finalFields = append(finalFields, parsedField)
	}

	return finalFields
}

// parseThingsDBMethods parses an array of ThingsDBMethod into an array of ParsedMethod
func parseThingsDBMethods(sourceMethods map[string]models.ThingsDBMethod) []models.ParsedMethod {
	var finalMethods []models.ParsedMethod

	for k, v := range sourceMethods {
		parsedMethod := models.ParsedMethod{}

		parsedMethod.Name = k
		parsedMethod.Doc = v.Doc
		parsedMethod.Arguments = v.Arguments

		finalMethods = append(finalMethods, parsedMethod)
	}

	return finalMethods
}
