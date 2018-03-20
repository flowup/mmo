package utils

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// RecursiveCopy downloads every folder and file from bucket with provided path to out
// example RecurisveCopy(client, "flowtest", "test1/", "/tmp/test")
func RecursiveCopy(client *storage.Client, bucket, path, out string) error {
	ctx := context.Background()
	it := client.Bucket(bucket).Objects(ctx, &storage.Query{
		Prefix: path,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		// Directory
		if strings.HasPrefix(attrs.ContentType, "application/x-www-form-urlencoded") {
			err = os.MkdirAll(out+strings.TrimPrefix(attrs.Name, path), 0777)
			if err != nil {
				return err
			}
			continue
		}
		newPath := strings.TrimPrefix(attrs.Name, path)
		fileName := newPath
		newPath = filepath.Dir(newPath)
		if newPath == "." {
			err = os.MkdirAll(out, 0777)
		} else {
			err = os.MkdirAll(out+newPath, 0777)
		}
		if err != nil {
			return err
		}
		data, err := readObject(client, bucket, attrs.Name)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(out+fileName, data, 0777)
		if err != nil {
			return err
		}

	}
	return nil
}

func readObject(client *storage.Client, bucket, object string) ([]byte, error) {
	ctx := context.Background()
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
