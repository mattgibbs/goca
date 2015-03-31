package goca

/*
#cgo CFLAGS: -I . -I /Users/mgibbs/Downloads/base-3.14.12.4/include -I /Users/mgibbs/Downloads/base-3.14.12.4/include/os/Darwin
#cgo LDFLAGS: -L/Users/mgibbs/Downloads/base-3.14.12.4/lib/darwin-x86 -lca -lcom
#include "cago.h"
#include "cadef.h"
*/
import "C"
import (
	"unsafe"
)

type ConnectionCallback func(connected bool)

type PV struct {
	Pvname string
	chid C.chid
	connected bool
	connection_callback ConnectionCallback
}

//export handleConnection
func handleConnection(chid C.chid, op C.long) {
	var pv_ptr unsafe.Pointer = C.ca_puser(chid)
	pv := (*PV)(pv_ptr)
	pv.chid = chid
	var connected bool
	if op == C.CA_OP_CONN_UP {
		connected = true
	} else {
		connected = false
	}
	pv.connected = connected
	pv.connection_callback(connected)
}

func NewPV(name string, cb ConnectionCallback) (*PV, int) {
	pv := &PV{
		Pvname: name,
		connection_callback: cb,
	}
	var chid C.chid
	ret_status := int(C.ca_create_channel_cgo(C.CString(pv.Pvname), unsafe.Pointer(C.handleConnection_cgo), unsafe.Pointer(pv), C.int(50), &chid))
	return pv, ret_status
}

func ContextCreate(preemptive bool) int {
	var mode = map[bool]int{true: 1, false: 0}
	return int(C.ca_context_create_cgo(C.int(mode[preemptive])))
}

func ContextDestroy() {
	C.ca_context_destroy()
}