package main

import (
	"C"
)

func main() { starter() }

func run() {
	starter()
}

//export Run
func Run() {
	run()
}

//export VoidFunc
func VoidFunc() { run() }

//export DllInstall
func DllInstall() { run() }

//export DllRegisterServer
func DllRegisterServer() { run() }

//export DllUnregisterServer
func DllUnregisterServer() { run() }

//export Exporter
func Exporter() {
	run()
}
