package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetIPv4() (string, error) {
	config := GetConfig()
	for _, url := range config.Ip.QueryUrls {
		ip, err := GetIPv4ByUrl(url)
		if err != nil {
			Logger.Error(
				fmt.Sprintf(
					"Failed to get IP information from %s. Error message: <%s>",
					url,
					err.Error(),
				),
			)
			continue
		}
		Logger.Info(fmt.Sprintf("Successfully retrieved IP information. IP address: %s", ip))
		return ip, nil
	}
	return "", errors.New(fmt.Sprintf("Failed to retrieve IP information"))
}

func GetIPv4ByUrl(url string) (string, error) {
	Logger.Info(fmt.Sprintf("Getting IP information from %s", url))
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Logger.Error(err)
		}
	}(resp.Body)
	content, err := io.ReadAll(resp.Body)
	return string(content), nil
}
