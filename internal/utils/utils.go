package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func GetJson(url string, target interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error happend while closing body: %s", err.Error())
		}
	}(resp.Body)
	return json.NewDecoder(resp.Body).Decode(target)
}
