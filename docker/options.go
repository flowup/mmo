package docker

type RunOptions struct {
	Image         string
	Command       string
	Volumes       []string
	PortPublishes []string
	Args          []string
	AutoRemove	  bool
	Environment   map[string]string
}

func CreateRunOptions(image string, command string) *RunOptions {
	return &RunOptions{Image: image, Command: command}
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

func (o *RunOptions) ToArgs() []string {
	args := make([]string, 0)
	args = append(args, o.Command)
	for _, volume := range o.Volumes {
		args = append(args, "-v")
		args = append(args, volume)
	}

	for _, port := range o.PortPublishes {
		args = append(args, "-p")
		args = append(args, port)
	}

	for key, value := range o.Environment {
		args = append(args, "-e")
		args = append(args, key+"="+value)
	}

	if o.AutoRemove {
		args = append(args, "--rm")
	}

	if o.Image != "" {
		args = append(args, o.Image)
	}
	args = append(args, o.Args...)
	return args
}
