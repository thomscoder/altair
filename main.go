package main

import "altair/src/build"

//go:generate go run script/getFiles.go

func main() {
	build.BuildAltair()
}
