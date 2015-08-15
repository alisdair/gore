package link

import (
	"encoding/base64"
	"encoding/json"
)

type Link struct {
	Id     int    `json:"id"`
	Target string `json:"target"`
}

func Parse(encoded string) (*Link, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	var link Link
	err = json.Unmarshal(decoded, &link)
	return &link, err
}

func Build(id int, url string) (string, error) {
	link := Link{Id: id, Target: url}

	bytes, err := json.Marshal(link)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}
