package gen

import (
	"fmt"
	"strings"

	"github.com/rickmoonex/gotigen/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// GenerateTypeCode generates Golang code for a struct representing a ParsedType
func GenerateTypeCode(parsedType *models.ParsedType) (string, error) {
	structFields := ""

	for _, field := range parsedType.Fields {
		parsed := parseFieldsToStruct(field)

		structFields += fmt.Sprintf("%s\n", parsed)
	}

	return fmt.Sprintf("type %s struct {\n%s}\n", parsedType.Name, structFields), nil
}

// parseFieldsToStruct parses a ParsedTypeField implementation into the contents of a struct
func parseFieldsToStruct(data models.ParsedTypeField) string {
	fieldName := toCamelCase(data.Name)

	return fmt.Sprintf("%s%s %s", "\t", fieldName, data.Type)
}

// toCamelCase convets a snake_case or kebab-case key to CamelCase
func toCamelCase(input string) string {
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == '_' || r == '-'
	})

	caser := cases.Title(language.English)

	for i := range parts {
		parts[i] = caser.String(parts[i])
	}

	return strings.Join(parts, "")
}
