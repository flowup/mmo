package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/utils"
	"github.com/flowup/mmo/utils/dockercmd"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Builder is structure to hold instance of service builder
type Builder struct {
	cli           *client.Client
	repository    string
	builtServices []Image
}

// GetBuilder is function to get instance of builder
func GetBuilder(repo string) (*Builder, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Builder{cli: cli, repository: repo, builtServices: make([]Image, 0)}, nil
}

// BuildService is function to build service's binary and alpine docker image
func (b *Builder) BuildService(service string) (Image, error) {

	err := b.buildBinary(service)
	if err != nil {
		return Image{}, errors.Wrap(err, "Failed to build binary of the service: "+service)
	}

	img, err := b.buildImage(service)
	if err != nil {
		return Image{}, errors.Wrap(err, "Failed to build docker image of the service: "+service)
	}

	b.builtServices = append(b.builtServices, img)

	return img, nil
}

// PushService is function to push image to local minikube registry
func (b *Builder) PushService(image Image) error {
	err := utils.PushImage(b.cli, image.GetFullname())
	if err != nil {
		return errors.Wrap(err, "Error pushing image "+image.GetFullname())
	}

	return nil
}

// Clean is function to remove built images - can be used to after pushing images to external registry
func (b *Builder) Clean() error {
	for _, service := range b.builtServices {
		b.cli.ImageRemove(context.Background(), service.GetFullname(), types.ImageRemoveOptions{})
	}

	return nil
}

func (b *Builder) buildBinary(service string) error {

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cmd := "go build -a -o " + service + "/bin/main ./" + service + "/cmd/" + service

	err = utils.PullImage(b.cli, dockercmd.Golang)
	if err != nil {
		return err
	}

	cont, err := b.cli.ContainerCreate(context.Background(), &container.Config{
		Image:      dockercmd.Golang,
		Cmd:        []string{"bash", "-c", cmd},
		WorkingDir: "/go/src/" + b.repository,
		Env: []string{
			"GOOS=linux",
			"CGO_ENABLED=0",
		},
	}, &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: pwd,
				Target: "/go/src/" + b.repository,
			},
		},
	}, nil, "")

	if err != nil {
		return err
	}

	// TODO: error of container not returned
	err = utils.ContainerRunStdout(b.cli, cont.ID)
	return err
}

func (b *Builder) buildImage(name string) (Image, error) {

	h := sha1.New()
	timeNow, err := time.Now().MarshalBinary()
	if err != nil {
		return Image{}, err
	}
	_, err = h.Write(timeNow)
	if err != nil {
		return Image{}, err
	}

	var img = Image{}

	img.Registry = dockercmd.MinikubeRegistry
	img.Name = b.repository + "-" + name
	img.Tag = strings.ToLower(base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(h.Sum(nil)))

	log.Debug("Building image: ", img.GetFullname())
	buildOptions := types.ImageBuildOptions{
		Tags: []string{img.GetFullname()},
	}

	ctx, err := b.createContext(name)
	if err != nil {
		return Image{}, err
	}

	res, err := b.cli.ImageBuild(context.Background(), ctx, buildOptions)
	// looks like we need to read the stream until the end because of some docker
	// daemon asynchronity. Response is returned before the command is finised.
	// ReadAll will wait until the docker stream is ultimately closed.
	ioutil.ReadAll(res.Body)

	return img, err
}

func (b *Builder) createContext(service string) (*bytes.Buffer, error) {

	buffer := new(bytes.Buffer)
	tarWriter := tar.NewWriter(buffer)
	defer tarWriter.Close()

	dockerfile := `FROM alpine
		ADD bin/main /main
		CMD ["./main"]`

	var files = []struct {
		Name string
		Body []byte
		Mode int64
	}{
		{"Dockerfile", []byte(dockerfile), 0600},
		{"bin/main", nil, 0755},
	}

	for _, file := range files {
		if file.Body == nil {
			body, err := ioutil.ReadFile(service + "/" + file.Name)
			if err != nil {
				return nil, err
			}

			file.Body = body
		}
		header := &tar.Header{
			Name: file.Name,
			Mode: file.Mode,
			Size: int64(len(file.Body)),
		}
		if err := tarWriter.WriteHeader(header); err != nil {
			return nil, err
		}
		if _, err := tarWriter.Write([]byte(file.Body)); err != nil {
			return nil, err
		}
	}

	if err := tarWriter.Close(); err != nil {
		return nil, err
	}

	return buffer, nil
}
