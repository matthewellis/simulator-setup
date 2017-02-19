package simulatorsetup

import (
	"fmt"
	"os"
)

func ApplyDataToSim(deviceName string, deviceVersion string) {
	deviceList, _ := getDevices()

	fmt.Println("Apply Device: " + deviceName)

	if deviceVersion == "latest" {
		latestIosVersion := deviceList.getLatestIOS()
		deviceList = deviceList.filterByIOS(latestIosVersion, true)
	}

	deviceList = deviceList.filterByName(deviceName)
	for _, device := range deviceList.Devices {
		fmt.Println(device)
		device.applySimulatorData(true)
	}
}

func DumpSimData(deviceName string) {
	deviceList, _ := getDevices()
	latestIosVersion := deviceList.getLatestIOS()

	deviceList = deviceList.filterByIOS(latestIosVersion, true)
	deviceList = deviceList.filterByName(deviceName)

	deviceCount := len(deviceList.Devices)
	if deviceCount == 0 {
		fmt.Println("Device ", deviceName, " not found")
		fmt.Println("Available devices:")
		deviceList, _ = getDevices()
		deviceList.filterByIOS(latestIosVersion, true).printDeviceList()
		os.Exit(1)
	} else if deviceCount != 1 {
		fmt.Println("Can only dump one device, matched ", deviceCount, " devices")
		deviceList.printDeviceList()
		os.Exit(1)
	} else {
		deviceList.Devices[0].dumpSimulatorData([]string{"Containers", "Documents", "Downloads", "Library", "Media"})
	}
}

func (deviceList DeviceList) printDeviceList() {
	for _, device := range deviceList.Devices {
		fmt.Println(device.Name)
	}
}
