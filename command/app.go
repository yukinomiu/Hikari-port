package main

import "runtime"

func main() {
	loadConfig()
	start()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
