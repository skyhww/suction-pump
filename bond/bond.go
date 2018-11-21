package bond

import (
	"suction-pump/input"
	"suction-pump/output"
)

type Bond struct {
	in    input.Input
	out   output.Output
	param map[string]string
}

func (bond *Bond) Validate() error {
	for _,enrty:=range bond.in.GetEntry(){
		enrty.GetName()
	}
}
