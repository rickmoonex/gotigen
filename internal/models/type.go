package models

// ThingsDBType represents a type definition returned from ThingsDB
type ThingsDBType struct {
	Name       string                    `mapstructure:"name"`
	Fields     []ThingsDBTypeField       `mapstructure:"fields"`
	HideId     bool                      `mapstructure:"hide_id"`
	Methods    map[string]ThingsDBMethod `mapstructure:"methods"`
	ModifiedAt int                       `mapstructure:"modified_at"` // Unix timestamp
	CreatedAt  int                       `mapstructure:"created_at"`  // Unix timestamp
	Relations  interface{}               `mapstructure:"relations"`
	Id         int                       `mapstructure:"type_id"`
	WarpOnly   bool                      `mapstructure:"wrap_only"`
}

// ThingsDBTypeField represents the field of a type returned by ThingsDB
// It's an array with a fixed length of 2. Where the first item represents
// the name of the field, and the second item the type. Both values are strings.
type ThingsDBTypeField []string

// ThingsDBMethod represents a method in ThingsDB
type ThingsDBMethod struct {
	Arguments       []string `mapstructure:"arguments"`
	Definition      string   `mapstructure:"definition"`
	Doc             string   `mapstructure:"doc"`
	WithSideEffects bool     `mapstructure:"with_side_effects"`
}

// ParsedType represents a ThingsDBType that has been parsed and sanitized
type ParsedType struct {
	Name    string
	Fields  []ParsedTypeField
	Methods []ParsedMethod
}

// ParsedTypeField represents a ThingsDBTypeField that has been parsed and sanitized
type ParsedTypeField struct {
	Name string
	Type string
}

// ParsedMethod represents a ThingsDBMethod that has been parsed and sanitized
type ParsedMethod struct {
	Name      string
	Doc       string
	Arguments []string
}
