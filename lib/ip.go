package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
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
	ip := string(content)
	err = CheckIPv4Format(ip)
	if err != nil {
		return "", err
	}
	return ip, nil
}

func CheckIPv4Format(ip string) error {
	reg := regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
	if !reg.MatchString(ip) {
		return errors.New(fmt.Sprintf("Invalid IPv4 address format."))
	}
	return nil
}
