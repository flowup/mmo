package docker

import "testing"

func TestImageFromString(t *testing.T) {
	tests := []struct {
		Image    string
		Expected Image
	}{
		{
			Image: "flowup/mmo-gen-grpc",
			Expected: Image{
				Registry: "",
				Name:     "flowup/mmo-gen-grpc",
				Tag:      "latest",
			},
		}, {
			Image: "flowup/mmo-gen-grpc:0.1",
			Expected: Image{
				Registry: "",
				Name:     "flowup/mmo-gen-grpc",
				Tag:      "0.1",
			},
		}, {
			Image: "gcr.io/flowup/mmo-gen-grpc",
			Expected: Image{
				Registry: "gcr.io",
				Name:     "flowup/mmo-gen-grpc",
				Tag:      "latest",
			},
		}, {
			Image: "gcr.io/flowup/mmo-gen-grpc:0.5",
			Expected: Image{
				Registry: "gcr.io",
				Name:     "flowup/mmo-gen-grpc",
				Tag:      "0.5",
			},
		}, {
			Image: "gcr.io/flowup/mmo-gen-grpc:0.5:0.5",
			Expected: Image{
				Registry: "gcr.io",
				Name:     "flowup/mmo-gen-grpc",
				Tag:      "0.5:0.5",
			},
		},
	}

	for _, test := range tests {
		output := ImageFromString(test.Image)
		if output.Name != test.Expected.Name {
			t.Errorf("Image name incorrect, got: %s, want: %s.", output.Name, test.Expected.Name)
		}
		if output.Registry != test.Expected.Registry {
			t.Errorf("Image registry incorrect, got: %s, want: %s.", output.Registry, test.Expected.Registry)
		}
		if output.Tag != test.Expected.Tag {
			t.Errorf("Image tag incorrect, got: %s, want: %s.", output.Tag, test.Expected.Tag)
		}
	}

}
