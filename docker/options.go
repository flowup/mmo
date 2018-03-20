package docker

// RunOptions is structure that represents Docker run options
type RunOptions struct {
	Image         string
	Command       string
	Volumes       []string
	PortPublishes []string
	Args          []string
	AutoRemove    bool
	Environment   map[string]string
	WorkingDir    string
}

// CreateRunOptions creates minimal options that runs image with command and can be extended later
func CreateRunOptions(image string, command string) *RunOptions {
	return &RunOptions{
		Image:         image,
		Command:       command,
		Volumes:       make([]string, 0),
		PortPublishes: make([]string, 0),
		Args:          make([]string, 0),
		Environment:   make(map[string]string),
	}
}

// AddDockerVolume adds option to mount docker volume, "-v [volume]" in cli
func (o *RunOptions) AddDockerVolume(name string) {
	o.Volumes = append(o.Volumes, name)
}

// MountHostVolume adds option to mount host volume, "-v [host:container]" in cli
func (o *RunOptions) MountHostVolume(host string, container string) {
	o.Volumes = append(o.Volumes, host+":"+container)
}

// AddArguments adds arguments to arguments passed to running program in container
func (o *RunOptions) AddArguments(args ...string) {
	o.Args = append(o.Args, args...)
}

// AddEnvVariable adds environment variables to container
func (o *RunOptions) AddEnvVariable(name string, value string) {
	o.Environment[name] = value
}

// ToArgs creates string used in shell for running container based on RunOptions
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

	if o.WorkingDir != "" {
		args = append(args, "-w")
		args = append(args, o.WorkingDir)
	}

	if o.Image != "" {
		args = append(args, o.Image)
	}

	args = append(args, o.Args...)
	return args
}
