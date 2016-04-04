package main

import "mime"

func init() {
	mime.AddExtensionType(".txt", "text/plain;charset=utf-8")
}
