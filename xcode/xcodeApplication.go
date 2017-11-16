package xcapp

import (
  "fmt"
  "os"
  "github.com/StevenWatremez/xcon/filesystem"
)

type XcodeDetail struct {
  RootPath, ContentsPath, ResourcesPath, Version, IconName string
}

func fetchPath(application string) string {
  var xcodePath string = xcfs.CheckXcodePath(application)
  if xcodePath != "" {
    return xcodePath
  } else {
    // search Xcode path by name
    return xcfs.FindXcodePath(application)
  }
}

func fetchVersion(path string) string {
  return "9.0"
}

func fetchIconName(path string) string {
  return "Xcode"
}

func ParseXcodeApplication(application, xcodeVersion string) XcodeDetail {
  // TODO : Manage path for XCode application
  path := fetchPath(application)
  
  if path == "" {
    fmt.Printf("\x1b[31;1mFailed to find path with application : %s !\x1b[0m", application)
    os.Exit(1)
  }

  contentsPath := path + "/Contents"
  resourcesPath := contentsPath + "/Resources"
  version := xcodeVersion
  if version == "" {
      version = fetchVersion(contentsPath)
  }
  iconName := fetchIconName(contentsPath)

  xcodeDetail := XcodeDetail{ path, contentsPath, resourcesPath, version, iconName }
  return xcodeDetail
}