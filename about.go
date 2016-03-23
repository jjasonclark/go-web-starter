package main

import "fmt"

// AppName of application
var AppName = "" // Updated from build

// BuildSHA is the Git SHA of last commit used in build
var BuildSHA = "" // Updated from build

// Version of application
var Version = "" // Updated from build

func versionDisplay() string {
	return fmt.Sprintf("%s\nVersion: %s SHA %s\n", AppName, Version, BuildSHA)
}
