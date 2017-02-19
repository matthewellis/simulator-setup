package simulatorsetup

import (
	"fmt"
	"log"
)

func (device Device) deviceDataPath() string {
	return basePath + device.UDID + "/data/"
}

func (device Device) dumpSimulatorData(dirs []string) {
	path := device.deviceDataPath()

	for _, dir := range dirs {
		err := copy_folder(path+dir, "./simulatorSettings/"+dir)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("copy finish")
		}
	}

}

func (device Device) applySimulatorData(removeFirst bool) {
	path := device.deviceDataPath()

	if removeFirst {
		removeContents(path)
	}

	err := copy_folder("./simulatorSettings/", path)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("copy finish")
	}
}
