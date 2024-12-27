package models

// ThingsDBProcedure represents a procedure definition returned from ThingsDB
type ThingsDBProcedure struct {
	Arguments       []string `mapstructure:"arguments"`
	CreatedAt       int      `mapstructure:"created_at"` // Unix timestamp
	Definition      string   `mapstructure:"definition"`
	Doc             string   `mapstructure:"doc"`
	Name            string   `mapstructure:"name"`
	WithSideEffects bool     `mapstructure:"with_side_effects"`
}

// ParsedProcedure represents a ThingsDBProcedure that has been parsed and sanitized
type ParsedProcedure struct {
	Name      string
	Doc       string
	Arguments []string
}
