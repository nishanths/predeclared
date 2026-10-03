package field

type t1 byte

func (byte *t1) m1() {} // want "^byte: shadows predeclared identifier$" "^byte: same name as predeclared identifier$"
func (t *t1) byte()  {} // want "^byte: same name as predeclared identifier$"

type _ struct {
	int8           // want "^int8: same name as predeclared identifier$"
	int16   uint16 // want "^int16: same name as predeclared identifier$"
	normal  uint16
	recover interface{ uint() uint32 } // want "^recover: same name as predeclared identifier$" "^uint: same name as predeclared identifier$"
}

type i1 interface {
	uintptr
	~float64
	normal() uint32
	uint() uint32 // want "^uint: same name as predeclared identifier$"
	new() any     // want "^new: same name as predeclared identifier$"
}
