package version

type Version struct {}

func NewVersion() *Version {
	return &Version{}
}

func (v *Version) Get() string {
	return "v1.0.2"
}

func Get() string {
	return "v1.0.2"
}
