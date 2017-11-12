package main

import (
  "fmt"
  "flag"
  "os"
)

func main() {
	templatePtr := flag.String("template", "banner", "Template name.")
  xcVersionPtr := flag.String("xcode-version", "", "Specify an Xcode version to override.")
  applicationPtr := flag.String("application", "", "Xcode application that you want to change icon.")
  flag.Parse()

  // test if application is empty
  if *applicationPtr == "" {
    flag.PrintDefaults()
    os.Exit(1)
  }

  fmt.Printf("Welcome to xcone !.\n")

  fmt.Printf("template: %s, xcode-version: %s, application: %s\n", *templatePtr, *xcVersionPtr, *applicationPtr)
}