package gen

import (
	"fmt"
	"strings"

	"github.com/rickmoonex/gotigen/internal/models"
)

// genMethod generates the Go func for a ParsedMethod
func genMethod(method models.ParsedMethod, parent string) string {
	finalCode := fmt.Sprintf("// %s\n", method.Doc)

	finalCode += fmt.Sprintf("func (p *%s) %s(", parent, method.Name)

	var typedArgs []string
	for _, v := range method.Arguments {
		if v == "this" {
			continue
		}
		typedArgs = append(typedArgs, fmt.Sprintf("%s interface{}", v))
	}
	finalCode += strings.Join(typedArgs, ", ")

	finalCode += ") (interface {}, error) {\n"

	// TODO: replace this with actual implementation
	finalCode += "\t return nil, nil\n"

	finalCode += "}"

	return finalCode
}
