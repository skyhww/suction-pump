package bond

import (
	"suction-pump/input"
	"suction-pump/output"
)

type Bond struct {
	input.Input
	output.Output
	param map[string]string
	Validator
}

func (bond *Bond) Validate() {

}
