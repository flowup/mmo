package docker

type RunOptions struct {
	Image         string
	Volumes       []string
	PortPublishes []string
	Args          []string
	Environment   map[string]string
}

func CreateRunOptions(image string) *RunOptions {
	return &RunOptions{Image: image}
}

func (o *RunOptions) AddDockerVolume(name string) {
	o.Volumes = append(o.Volumes, name)
}

func (o *RunOptions) MountHostVolume(host string, container string) {
	o.Volumes = append(o.Volumes, host+":"+container)
}

func (o *RunOptions) AddArguments(args ...string) {
	o.Args = append(o.Args, args...)
}

func (o *RunOptions) AddEnvVariable(name string, value string) {
	o.Environment[name] = value
}
