package docker

import (
	"github.com/docker/docker/client"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/flowup/mmo/utils"
	"github.com/flowup/mmo/utils/dockercmd"
	"os"
	"crypto/sha1"
	"encoding/base64"
	"time"
	"github.com/docker/docker/api/types"
	"archive/tar"
	"bytes"
	"io/ioutil"
	"fmt"
	"strings"
)

type Builder struct {
	cli       *client.Client
	goPackage string
}

func GetBuilder(goPackage string) (*Builder, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Builder{cli: cli, goPackage: goPackage}, nil
}

func (b *Builder) BuildService(service string) (string, error) {

	err := b.buildBinary(service)
	if err != nil {
		return "", err
	}

	image, err := b.buildImage(service)

	fmt.Println(image)

	return image, err
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
		WorkingDir: "/go/src/" + b.goPackage,
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
				Target: "/go/src/" + b.goPackage,
			},
		},
	}, nil, "")

	if err != nil {
		return err
	}

	err = utils.ContainerRunStdout(b.cli, cont.ID)
	return err
}

func (b *Builder) buildImage(service string) (string, error) {

	h := sha1.New()
	timeNow, err := time.Now().MarshalBinary()
	if err != nil {
		return "", err
	}
	h.Write(timeNow)

	imageTag := service + "-" + strings.ToLower(base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(h.Sum(nil)))

	buildTag := dockercmd.MinikubeRegistry + "/" + b.goPackage + ":" + imageTag
	buildOptions := types.ImageBuildOptions{
		Tags: []string{buildTag },
	}

	ctx, err := b.createContext(service)

	_, err = b.cli.ImageBuild(context.Background(), ctx, buildOptions)

	return buildTag, err
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
			fmt.Println(len(body))
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
