package main

import (
  "fmt"
  "flag"
  "os"
  log "github.com/sirupsen/logrus"

  "github.com/StevenWatremez/xcon/xcode"
  "github.com/StevenWatremez/xcon/processing"
)

func main() {
  // Subcommands
  autoCommand := flag.NewFlagSet("auto", flag.ExitOnError)
  configCommand := flag.NewFlagSet("config", flag.ExitOnError)
  targetCommand := flag.NewFlagSet("target", flag.ExitOnError)

  // Auto subcommand

  // Config subcommand

  // Target subcommand
  // --application
  targetApplicationPtr := targetCommand.String("application", "", "Xcode application that you want to change icon.")
  targetTemplatePtr := targetCommand.String("template", "banner", "Template name.")
  targetXCodeVersionPtr := targetCommand.String("xcode-version", "", "Specify an Xcode version to override.")
	
  // Verify that a subcommand has been provided
  // os.Arg[0] is the main command
  // os.Arg[1] will be the subcommand
  if len(os.Args) < 2 {
    fmt.Println("\x1b[31;1mlist or count subcommand is required !\x1b[0m")
    os.Exit(1)
  }
  
  // Switch on the subcommand
  // Parse the flags for appropriate FlagSet
  // FlagSet.Parse() requires a set of arguments to parse as input
  // os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
  switch os.Args[1] {
  case "auto":
    autoCommand.Parse(os.Args[2:])
  case "config":
    configCommand.Parse(os.Args[2:])
  case "target":
    targetCommand.Parse(os.Args[2:])
  default:
    flag.PrintDefaults()
    os.Exit(1)
  }

  fmt.Println(`
   |\  —————————————————————————————————————————————————————————————————  /|
   | |                            Welcome To xcon                        | |
   |/  —————————————————————————————————————————————————————————————————  \|
  `)

  // Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
  // FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
  if autoCommand.Parsed() {
    log.Warn("not implemented yet.")
  } else if configCommand.Parsed() {
    log.Warn("not implemented yet.")
  } else if targetCommand.Parsed() {
    if *targetApplicationPtr == "" {
      targetCommand.PrintDefaults()
      os.Exit(1)
    }

    fmt.Printf(`---------------------------------------------------------------
      application: '%s' 
      template: '%s'
      xcode-version: '%s'
---------------------------------------------------------------
`,
      *targetApplicationPtr,
      *targetTemplatePtr,
      *targetXCodeVersionPtr)

    log.Warn("under construct.")
    xcode := xcapp.ParseXcodeApplication(*targetApplicationPtr, *targetXCodeVersionPtr)

    data, err := Asset("template-" + *targetTemplatePtr + ".png")
    if err != nil {
      // Asset was not found.
      fmt.Println("Asset was not found !")
      os.Exit(1)
    }
    
    //
    templateImage := proc.ProcessTemplate(data, xcode.Version)

    iconFileName := "/" + xcode.IconName + ".icns"
    rootPathIcns := xcode.RootPath + iconFileName
    resourcesPathIcns := xcode.ResourcesPath + iconFileName
    //"/Users/swatremez/Desktop/test.icns"
    proc.CreateIcns(templateImage, rootPathIcns)
    proc.CreateIcns(templateImage, resourcesPathIcns)
    
    log.Info("Xcode detail: ", xcode)
  }
}