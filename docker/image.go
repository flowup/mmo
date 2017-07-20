package docker

// Image is structure to hold docker image name info
type Image struct {
	Name     string
	Tag      string
	Registry string
}

// GetFullname is function that returns full name of docker image consisting of the registry, name and the tag
func (i *Image) GetFullname() string {
	return i.Registry + "/" + i.Name + ":" + i.Tag
}
