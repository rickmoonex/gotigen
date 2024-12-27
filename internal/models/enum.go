package models

// ThingsDBEnum represent an enum returned by ThingsDB
type ThingsDBEnum struct {
	CreatedAt  int                       `mapstructure:"created_at"` // Unix timestamp
	Default    string                    `mapstructure:"default"`
	Id         int                       `mapstructure:"enum_id"`
	Members    []ThingsDBEnumMember      `mapstructure:"members"`
	Methods    map[string]ThingsDBMethod `mapstructure:"methods"`
	ModifiedAt int                       `mapstructure:"modified_at"` // Unix timestamp
	Name       string                    `mapstructure:"name"`
}

// ThingsDBEnumMember represents a member of an enum returned by ThingsDB
// It's an array with a fixed length of 2 where the first item is the
// name of the member as a string. And the second item is the integer
// associated with that member.
type ThingsDBEnumMember []interface{}

// ParsedEnum represent a ThingsDBEnum that has been parsed and sanitized
type ParsedEnum struct {
	Name         string
	DefaultValue int
	Methods      []ParsedMethod
}

// ParsedEnumMember represents a ThingsDBEnumMember that has been parsed and sanitized
type ParsedEnumMember struct {
	Key   string
	Value int
}
