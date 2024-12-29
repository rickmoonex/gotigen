package gen

import (
	"fmt"

	"github.com/rickmoonex/gotigen/internal/models"
)

// GenerateParsedEnums takes a ParsedEnum and returns the Go code as []byte
func GenerateParsedEnums(enum models.ParsedEnum) []byte {
	var finalCode string

	finalCode += fmt.Sprintf("%s\n\n", genEnumType(enum))
	finalCode += genEnumConsts(enum)

	for _, v := range enum.Methods {
		finalCode += "\n\n"

		finalCode += genMethod(v, enum.Name)
	}

	return []byte(finalCode)
}

// genEnumType generates the Go type definition for the ParsedEnum
func genEnumType(enum models.ParsedEnum) string {
	return fmt.Sprintf("type %s int", enum.Name)
}

// genEnumConsts generates the Go const for the ParsedEnum's members
func genEnumConsts(enum models.ParsedEnum) string {
	finalCode := "const (\n"

	for _, v := range enum.Members {
		finalCode += fmt.Sprintf("\t%s_%s %s = %v\n", enum.Name, v.Key, enum.Name, v.Value)
	}

	finalCode += ")"

	return finalCode
}
