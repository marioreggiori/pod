package global

// --verbose
var isVerbose = false

func IsVerbose() bool {
	return isVerbose
}

func SetIsVerbose(state bool) {
	isVerbose = state
}

// --tag
var imageTag string

func ImageTag() string {
	return imageTag
}

func SetImageTag(tag string) {
	imageTag = tag
}
