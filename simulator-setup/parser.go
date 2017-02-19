package simulatorsetup

import (
	"os"
	"strconv"
	"strings"

	"github.com/matthewellis/simulator-setup/fileIO"

	"github.com/mkrautz/plist/xmlplist"
)

type DevicePlist struct {
	UDID       string `plist:"UDID"`
	DeviceType string `plist:"deviceType"`
	Name       string `plist:"name"`
	Runtime    string `plist:"runtime"`
}

func parseDevicePlist(file *os.File) Device {
	deviceSet := new(DevicePlist)
	fileContent, _ := fileIO.Load(file)
	xmlplist.Unmarshal(fileContent, deviceSet)
	return deviceSet.toDevice()
}

func (devicePlist DevicePlist) toDevice() Device {
	device := Device{}
	device.UDID = devicePlist.UDID
	device.Name = devicePlist.Name
	osType, deviceiOSVershion := convertRuntime(devicePlist.Runtime)
	device.osType = osType
	device.iOSVershion = deviceiOSVershion

	return device
}

func convertRuntime(runtime string) (string, iOSVershion) {
	runtimeArray := strings.Split(runtime, ".")
	runtimeString := runtimeArray[len(runtimeArray)-1]
	runtimeArray = strings.Split(runtimeString, "-")
	major, _ := strconv.ParseInt(runtimeArray[1], 10, 0)
	minor, _ := strconv.ParseInt(runtimeArray[2], 10, 0)
	return runtimeArray[0], iOSVershion{int(major), int(minor)}
}
