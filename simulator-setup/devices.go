package simulatorsetup

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/matthewellis/simulator-setup/fileIO"
)

type Device struct {
	UDID        string
	Name        string
	osType      string
	iOSVershion iOSVershion
}

type DeviceList struct {
	Devices []Device
}

var basePath = os.Getenv("HOME") + "/Library/Developer/CoreSimulator/Devices/"

func getDevices() (DeviceList, error) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	devices := *new([]Device)
	for _, dirFile := range files {
		filename := dirFile.Name()
		if dirFile.IsDir() && len(filename) == 36 { // replace with regex

			file, _ := fileIO.Open(basePath + filename + "/device.plist")
			device := parseDevicePlist(file)
			devices = append(devices, device)
		}
	}

	return DeviceList{devices}, nil
}

func (deviceList DeviceList) getLatestIOS() iOSVershion {
	latestiOSVersion := iOSVershion{0, 0}

	for _, device := range deviceList.Devices {
		deviceiOSVershion := device.iOSVershion
		if deviceiOSVershion.major > latestiOSVersion.major ||
			deviceiOSVershion.minor > latestiOSVersion.minor {
			latestiOSVersion = deviceiOSVershion
		}
	}

	return latestiOSVersion
}

func (deviceList DeviceList) filterByName(name string) DeviceList {
	newDeviceList := DeviceList{}

	for _, device := range deviceList.Devices {
		if strings.HasSuffix(name, "*") {
			if strings.Contains(device.Name, strings.Replace(name, "*", "", 1)) {
				newDeviceList.Devices = append(newDeviceList.Devices, device)
			}
		} else if device.Name == name {
			newDeviceList.Devices = append(newDeviceList.Devices, device)
		}
	}

	return newDeviceList
}

func (deviceList DeviceList) filterByIOS(iOSVershion iOSVershion, matchMinor bool) DeviceList {
	newDeviceList := DeviceList{}

	for _, device := range deviceList.Devices {
		if device.iOSVershion.Equal(iOSVershion, matchMinor) {
			newDeviceList.Devices = append(newDeviceList.Devices, device)
		}
	}

	return newDeviceList
}
