package utils

import (
	"cloud.google.com/go/storage"
	"context"
	"io/ioutil"
	"strings"
)

func GetFileData(uri string) ([]byte, error) {
	if strings.HasPrefix(uri, "gs://") {
		splitted := strings.SplitN(strings.TrimLeft(uri,"gs://"), "/", 2)
		if len(splitted) != 2 {
			return nil, ErrInvalidGSUri
		}
		return downloadFromGS(splitted[0], splitted[1])
	} else if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return nil, nil
	}

	return nil, nil
}

func downloadFromGS(bucket, object string) ([]byte, error) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	rc, err := client.Bucket(bucket).Object(object).NewReader(context.Background())
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