package parser

import (
	"errors"

	"github.com/rickmoonex/gotigen/internal/models"
)

// ParseThingsDBEnums parses an array of ThingsDBEnum into ParsedEnum
func ParseThingsDBEnums(sourceEnums []models.ThingsDBEnum) (*[]models.ParsedEnum, error) {
	var finalEnums []models.ParsedEnum

	for _, v := range sourceEnums {
		parsedEnum := models.ParsedEnum{}

		parsedEnum.Name = v.Name
		parsedEnum.Methods = parseThingsDBMethods(v.Methods)

		members, err := parseThingsDBEnumMembers(v.Members)
		if err != nil {
			return nil, err
		}

		parsedEnum.Members = members
		parsedEnum.DefaultValue = getDefaultEnumValue(members, v.Default)

		finalEnums = append(finalEnums, parsedEnum)
	}

	return &finalEnums, nil
}

// getDefaultEnumValue looks up the enums default value based on the member key.
// defaults to 0 if not found.
func getDefaultEnumValue(members []models.ParsedEnumMember, key string) int {
	value := 0

	for _, v := range members {
		if v.Key == key {
			return v.Value
		}
	}

	return value
}

// parseThingsDBEnumMembers parses an array of ThingsDBEnumMember into ParsedEnumMember
func parseThingsDBEnumMembers(sourceMembers []models.ThingsDBEnumMember) ([]models.ParsedEnumMember, error) {
	var finalMembers []models.ParsedEnumMember

	for _, v := range sourceMembers {
		parsedMember := models.ParsedEnumMember{}

		key, ok := v[0].(string)
		if !ok {
			return nil, errors.New("enum member key cannot be cast to a string")
		}

		value, ok := v[1].(int8)
		if !ok {
			return nil, errors.New("enum member value cannot be cast to an int")
		}

		parsedMember.Key = key
		parsedMember.Value = int(value)

		finalMembers = append(finalMembers, parsedMember)
	}

	return finalMembers, nil
}
