package proc

import (
	"os"
	"bytes"
	"image"
	"image/png"
	"github.com/tonyhb/goicns"
	log "github.com/sirupsen/logrus"
)

func ProcessTemplate(data []byte, xcodeVersion string) image.Image {

	 //img, _, _ := image.Decode(bytes.NewReader(data))

  // convert []byte to image for saving to file
  img, errPngDecode := png.Decode(bytes.NewReader(data))
  if errPngDecode != nil {
    //fmt.Println("Error decoding image %s: %s", path, err.Error())
    //fmt.Println("Error decoding image | error : %s", err.Error())
    log.Fatal(errPngDecode)
    os.Exit(1)
  }
  return img
}

func CreateIcns(img image.Image, path string) {
  icns := goicns.NewICNS(img)
  if errIcnsConstruct := icns.Construct(); errIcnsConstruct != nil {
    //fmt.Println("Error encofing ICNS %s: %s", path, err.Error())
    log.Fatalf("Error encoding ICNS | %s", errIcnsConstruct.Error())
    return
  }

  if errIcnsCreation := icns.WriteToFile(path, 0666); errIcnsCreation != nil {
    log.Fatalf("Error writing icns | %s", errIcnsCreation.Error())
    return  
  }
}