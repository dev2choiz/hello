package version

const version = "v1.0.3"

type Version struct {}

// NewVersion return a *Version instance
func NewVersion() *Version {
	return &Version{}
}

// Get return the app version
func (v *Version) Get() string {
	return version
}

// Get return the app version
// Deprecated
func Get() string {
	return version
}
