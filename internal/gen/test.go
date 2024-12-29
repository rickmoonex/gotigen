package gen

type Enum1 int

const (
	Enum1_OPTION1 Enum1 = 0
	Enum1_OPTION2 Enum1 = 1
	Enum1_OPTION3 Enum1 = 2
)

// Print it to the screen!!
func (p *Enum1) print() (interface{}, error) {
	return nil, nil
}
