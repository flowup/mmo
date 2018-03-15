package docker

import "strings"

// Image is structure to hold docker image name info
type Image struct {
	Name     string
	Tag      string
	Registry string
}

// ImageFromString parses string and returns Docker image
func ImageFromString(name string) *Image {

	slashes := strings.Count(name, "/")
	registry := ""
	if slashes > 1 {
		names := strings.Split(name, "/")
		registry = names[0]
		name = strings.Join(names[1:], "/")
	}

	tags := strings.Split(name, ":")
	tag := "latest"
	image := ""

	if len(tags) > 1 {
		image = tags[0]
		name = strings.Join(tags[1:], ":")
		tag = name
	} else {
		image = name
	}

	return &Image{Registry: registry, Name: image, Tag: tag}
}

// GetFullname is function that returns full name of docker image consisting of the registry, name and the tag
func (i *Image) GetFullname() string {
	return i.Registry + "/" + i.Name + ":" + i.Tag
}
