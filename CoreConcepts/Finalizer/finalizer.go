/*
	Finalizer is a function which only gets called for an object when it that object has lost it's
	reference from everywhere. In such a case the GC(garbage collector will clean the resource). So,
	before the GC cleans this object the finalizer function is called.
	Finalizers can be added to an object using runtime package.
	DOUBT:
	Need to figure out why the finalizer is called only when I call the GC the 2nd time.
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

type timeEater struct {
	t int
}

func eatTime() {
	//time.Sleep(1 * time.Second)
	te := &timeEater{1}
	runtime.SetFinalizer(te, func(te *timeEater) { fmt.Println("This is my will pleae read it as I'm going to die.") })
	fmt.Println("Functions ends, now the local objects created won't be having referrence anywhere else.")
}

func main() {
	fmt.Println("##################FINALIZER################")
	eatTime()
	runtime.GC()
	time.Sleep(1 * time.Second)
	runtime.GC()
	time.Sleep(5 * time.Second)
}
