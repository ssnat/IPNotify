package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const DNSPodAPIUrl = "https://dnsapi.cn/Record.Modify"

func UpdateDNSPodRecords(ip string) error {

	config := GetConfig()

	if !config.DDNS.DNSPod.Enabled {
		return nil
	}

	Logger.Info("Updating DNSPod records ...")

	for _, record := range config.DDNS.DNSPod.Records {

		values := url.Values{
			"login_token": {config.DDNS.DNSPod.LoginToken},
			"format":      {"json"},
			"domain_id":   {record.DomainId},
			"record_id":   {record.RecordId},
			"sub_domain":  {record.SubDomain},
			"record_type": {record.RecordType},
			"record_line": {record.RecordLine},
			"value":       {ip},
			"ttl":         {fmt.Sprintf("%d", record.TTL)},
		}

		Logger.Info(fmt.Sprintf("Updating DNSPod record: %s.%s - %s", record.DomainId, record.SubDomain, ip))

		err := UpdateDNSPodRecord(values)

		if err != nil {
			return err
		}

		Logger.Info(fmt.Sprintf("Updated DNSPod record: %s.%s - %s", record.DomainId, record.SubDomain, ip))

	}
	return nil
}

func UpdateDNSPodRecord(values url.Values) error {
	res, err := http.PostForm(DNSPodAPIUrl, values)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Logger.Error(err)
		}
	}(res.Body)

	var resJson struct {
		Status struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"status"`
	}

	err = json.NewDecoder(res.Body).Decode(&resJson)

	if err != nil {
		return err
	}

	if resJson.Status.Code != "1" {
		return errors.New(fmt.Sprintf("code: %s, message: %s", resJson.Status.Code, resJson.Status.Message))
	}

	return nil
}
