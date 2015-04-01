cago: A Go interface to EPICS Channel Access
============================================

cago is a Go binding for the Channel Access library.  It uses cgo to call the C-based channel access library, and wraps it into an easy to use, object-oriented interface.

How to use
----------
Make a new PV struct with cago.NewPV.  Define connection and event callbacks.
	
	
	package main
	import "github.com/mattgibbs/goca"
	import (
			"fmt"
			"time"
	)

	func connectionHandler(connected bool) {
		if connected {
			fmt.Println("Connected to the PV!")
		} else {
			fmt.Println("Disconnected from the PV.")
		}
	}

	func eventHandler(val float32) {
		fmt.Printf("Got new value: %v\n", val)
	}

	func main() {
		goca.ContextCreate(true)
		mypv := goca.NewPV("mgibbsHost:ai1", connectionHandler, eventHandler, true)
		fmt.Printf("goca.NewPV() returned for pv %v\n", mypv.Pvname)
		go func() {
			fmt.Println("Waiting for callbacks...")
			time.Sleep(5 * time.Second)
			fmt.Println("OK done.")
		}()
		time.Sleep(5 * time.Second)
		mypv.Disconnect()
		time.Sleep(1 * time.Second)
		fmt.Println("Destroying Context.")
		goca.ContextDestroy()
	}