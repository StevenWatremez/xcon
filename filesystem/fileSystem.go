package xcfs

import (
  "os"
  s "strings"
  log "github.com/sirupsen/logrus"
)

// check here if application has version.plist and info.plist
func CheckXcodePath(path string) string {
  isPath := s.Contains(path, "/")
  if isPath {
    usablePath := path
    if !s.Contains(usablePath, ".app") {
      usablePath += ".app"
    }
    log.Info("usable path =  '", usablePath, "'")
    contentsPath := usablePath + "/Contents"
    versionFileName := "version.plist"
    infoFileName := "info.plist"
    // hasVersionFile := true//os.path.exists(contentsPath + "/" + versionFileName)
    // hasInfoFile := true//os.path.exists(contentsPath + "/" + infoFileName)
    versionFilePath := contentsPath + "/" + versionFileName
    _, errVersionFile := os.Stat(versionFilePath)
    if errVersionFile != nil {
      log.Warn(errVersionFile)
      return ""
    }

    infoFilePath := contentsPath + "/" + infoFileName
    _, errInfoFile := os.Stat(infoFilePath)
    if errInfoFile != nil {
      log.Warn(errInfoFile)
      return ""
    }
    
    return usablePath
  }
  return ""
}

func FindXcodePath(name string) string {
	usablePath := name
  if !s.Contains(usablePath, ".app") {
    usablePath += ".app"
  }

  applicationsPath := "/Applications"
  xcodePath := applicationsPath + "/" + usablePath
  _, errXcodePath := os.Stat(xcodePath)
  if errXcodePath != nil {
    log.Warn(errXcodePath)
    return ""
  }
	return xcodePath
}