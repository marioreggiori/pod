package global

var isVerbose = false

func IsVerbose() bool {
	return isVerbose
}

func SetIsVerbose(state bool) {
	isVerbose = state
}
