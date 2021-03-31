package gobber

import "unsafe"

/*eface is the internal struct of a empty interface
ref: https://github.com/golang/go/blob/095ba225973152734b0722e7c5758accb2639c15/src/runtime/runtime2.go#L207
*/
type eface struct {
	typ *struct{}
	ptr unsafe.Pointer
}

func ifaceToPtr(inter interface{}) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&inter)).ptr
}
