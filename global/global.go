package global

var flags *Flags

type Flags struct {
	Verbose       bool
	ImageTag      string
	EnvVariables  []string
	MappedPorts   []string
	MappedVolumes []string
}

func (f *Flags) Set() {
	flags = f
}

func IsVerbose() bool {
	return flags.Verbose
}

func ImageTag() string {
	return flags.ImageTag
}

func Mounts() []string {
	return flags.MappedVolumes
}

func Ports() []string {
	return flags.MappedPorts
}

func EnvVariables() []string {
	return flags.EnvVariables
}
