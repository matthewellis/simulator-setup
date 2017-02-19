package simulatorsetup

type iOSVershion struct {
	major int
	minor int
}

func (iOSVershionA iOSVershion) Equal(iOSVershionB iOSVershion, matchMinor bool) bool {
	if iOSVershionA.major == iOSVershionB.major {

		if iOSVershionA.minor == iOSVershionB.minor || matchMinor == false {
			return true
		}
	}

	return false
}
