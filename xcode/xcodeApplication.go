package xcapp

import (
  "fmt"
  "os"
  "github.com/StevenWatremez/xcon/filesystem"
  log "github.com/sirupsen/logrus"
  plist "howett.net/plist"
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

type versionPlist struct {
  Version string `plist:"CFBundleShortVersionString"`
}

func fetchVersion(path string) string {
  filename := path + "/version.plist"
  xmlFile, err := os.Open(filename)
  if err != nil {
    log.Warnf("plist: error opening plist: %s", err)
    return ""
  }
  defer xmlFile.Close()

  var data versionPlist
  decoder := plist.NewDecoder(xmlFile)
  errDecode := decoder.Decode(&data)
  if errDecode != nil {
    log.Warn(errDecode)
    return ""
  }
  log.Infof("xcode version : %s", data.Version)
  return data.Version
}

type infoPlist struct {
  Icon string `plist:"CFBundleIconFile"`
}

func fetchIconName(path string) string {
  filename := path + "/Info.plist"
  xmlFile, err := os.Open(filename)
  if err != nil {
    log.Warnf("plist: error opening plist: %s", err)
    return ""
  }
  defer xmlFile.Close()

  var data infoPlist
  decoder := plist.NewDecoder(xmlFile)
  errDecode := decoder.Decode(&data)
  if errDecode != nil {
    log.Warn(errDecode)
    return ""
  }
  log.Infof("xcode icon name : %s", data.Icon)
  return data.Icon
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