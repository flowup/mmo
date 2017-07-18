package docker

type Image struct {
	Name     string
	Tag      string
	Registry string
}

func (i *Image) GetFullname() string {
	return i.Registry + "/" + i.Name + ":" + i.Tag
}
