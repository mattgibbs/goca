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
type SubscriptionCallback func(val float32)


type PV struct {
	Pvname string
	chid C.chid
	Connected bool
	ConnCallback ConnectionCallback
	EventCallback SubscriptionCallback
}

func (pv *PV) Connect() int {
	var chid C.chid
	ret_status := int(C.ca_create_channel_cgo(C.CString(pv.Pvname), unsafe.Pointer(C.handleConnection_cgo), unsafe.Pointer(pv), C.int(50), &chid))
	pv.chid = chid
	return ret_status
}

func (pv *PV) Disconnect() int {
	return int(C.ca_clear_channel(pv.chid))
}

func (pv *PV) Monitor() int {
	count := 0
	mask := uint(C.DBE_VALUE) | uint(C.DBE_ALARM)
	val_type := (C.chtype)(C.DBR_DOUBLE)
	var evid C.evid
	ret_status := int(C.ca_create_subscription_cgo(val_type, C.ulong(count), pv.chid, C.long(mask), unsafe.Pointer(C.handleEvent_cgo), unsafe.Pointer(pv), &evid))
	return ret_status
}

//export handleEvent
func handleEvent(pv_ptr unsafe.Pointer, chid C.chid, val_type C.long, count C.long, val_ptr unsafe.Pointer, status int) {
	pv := (*PV)(pv_ptr)
	value := float32(*((*C.double)(val_ptr)))
	pv.EventCallback(value)
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
	pv.Connected = connected
	pv.Monitor()
	pv.ConnCallback(connected)
}

func NewPV(name string, ccb ConnectionCallback, scb SubscriptionCallback, autoconnect bool) (*PV) {
	pv := &PV{
		Pvname: name,
		ConnCallback: ccb,
		EventCallback: scb,
	}
	if autoconnect {
		pv.Connect()
	}
	return pv
}

func ContextCreate(preemptive bool) int {
	var mode = map[bool]int{true: 1, false: 0}
	return int(C.ca_context_create_cgo(C.int(mode[preemptive])))
}

func ContextDestroy() {
	C.ca_context_destroy()
}