package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type DnsRecord struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

func UpdateCloudflareRecords(ip string) error {
	config := GetConfig()

	if !config.DDNS.Cloudflare.Enabled {
		return nil
	}

	Logger.Info("Updating Cloudflare records ...")

	for _, record := range config.DDNS.Cloudflare.Records {

		url := "https://api.cloudflare.com/client/v4/zones/" + record.ZoneId + "/dns_records/" + record.RecordId

		DNSRecord := &DnsRecord{
			Type:    record.RecordType,
			Name:    record.RecordName,
			Content: ip,
			TTL:     record.TTL,
			Proxied: record.Proxied,
		}

		data, err := json.Marshal(DNSRecord)

		if err != nil {
			return nil
		}

		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))

		if err != nil {
			return nil
		}

		req.Header.Set("Authorization", "Bearer "+config.DDNS.Cloudflare.APIToken)

		req.Header.Set("Content-Type", "application/json")

		Logger.Info(fmt.Sprintf("Updating Cloudflare record: %s.%s - %s", record.ZoneId, record.RecordName, ip))

		err = UpdateCloudflareRecord(req)

		if err != nil {
			Logger.Error(err)
			continue
		}

		Logger.Info(fmt.Sprintf("Updated Cloudflare record: %s.%s - %s", record.ZoneId, record.RecordName, ip))

	}

	return nil
}

func UpdateCloudflareRecord(req *http.Request) error {
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Logger.Error(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Cloudflare API returned status code: %d", resp.StatusCode))
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	if result["success"] != true {
		return errors.New(fmt.Sprintf("Cloudflare API returned error: %s", result["errors"]))
	}
	return nil
}
