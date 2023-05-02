package encoder

import (
	"encoding/base64"
	"io/ioutil"
)

// ImgstoBase64 takes a string slice and ranges over it
// converting each to a base64 encoded string and adding to a resulting slice.
// If it succeeds it returns the slice and a nil err.
func ImgstoBase64(paths []string) (result []string, err error) {
	for _, path := range paths {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return []string{}, err
		}
		encoded := base64.StdEncoding.EncodeToString(file)
		result = append(result, encoded)
	}
	return result, nil
}
