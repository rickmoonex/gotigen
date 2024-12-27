package models

type ParsedFile struct {
	Types []ParsedType
}

type ParsedType struct {
	Name   string
	Fields []ParsedTypeField
}

type ParsedTypeField struct {
	Name string
	Type string
}
