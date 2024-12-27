package models

// TypeInfo represents all information about a specific type
type TypeInfo struct {
	Name   string     `json:"name"`
	Fields [][]string `json:"fields"`
}
