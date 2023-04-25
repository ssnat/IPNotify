package lib

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type IIPConfig struct {
	QueryUrls []string `yaml:"query_urls"`
	Interval  string   `yaml:"interval"`
}

type IEmailConfig struct {
	SMTPHost   string   `yaml:"smtp_host"`
	SMTPPort   int      `yaml:"smtp_port"`
	SMTPUser   string   `yaml:"smtp_user"`
	SMTPPasswd string   `yaml:"smtp_passwd"`
	From       string   `yaml:"from"`
	To         []string `yaml:"to"`
}

type IDNSPodConfig struct {
	Enabled    bool   `yaml:"enabled"`
	LoginToken string `yaml:"login_token"`
	Records    []struct {
		DomainId   string `yaml:"domain_id"`
		RecordId   string `yaml:"record_id"`
		SubDomain  string `yaml:"sub_domain"`
		RecordType string `yaml:"record_type"`
		RecordLine string `yaml:"record_line"`
		TTL        int
	} `yaml:"records"`
}

type ICloudflareConfig struct {
	Enabled  bool   `yaml:"enabled"`
	APIToken string `yaml:"api_token"`
	Records  []struct {
		ZoneId     string `yaml:"zone_id"`
		RecordId   string `yaml:"record_id"`
		RecordName string `yaml:"record_name"`
		RecordType string `yaml:"record_type"`
		TTL        int    `yaml:"ttl"`
		Proxied    bool   `yaml:"proxied"`
	} `yaml:"records"`
}

type IDDNSConfig struct {
	DNSPod     IDNSPodConfig     `yaml:"DNSPod"`
	Cloudflare ICloudflareConfig `yaml:"Cloudflare"`
}

type IConfig struct {
	Ip    IIPConfig    `yaml:"ip"`
	Email IEmailConfig `yaml:"email"`
	DDNS  IDDNSConfig  `yaml:"DDNS"`
}

var Config *IConfig = nil

func init() {
	configFilePath, err := GetConfigFilePath()
	if err != nil {
		Logger.PanicError(err)
	}
	buffer, err := os.ReadFile(configFilePath)
	if err != nil {
		Logger.PanicError(err)
	}

	err = yaml.Unmarshal(buffer, &Config)

	if err != nil {
		Logger.PanicError(err)
	}
}

func GetConfigFilePath() (string, error) {

	filename := "config.yaml"
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}

	defaultFilePath := filepath.Join(root, filename)

	Logger.Info(fmt.Sprintf("Default Configuration file path is %s", defaultFilePath))

	var filePath string

	flag.StringVar(&filePath, "f", defaultFilePath, "Path to the configuration file")

	flag.Parse()

	if !filepath.IsAbs(filePath) {
		filePath = filepath.Join(root, filePath)
	}

	Logger.Info(fmt.Sprintf("Configuration file path is %s", filePath))

	return filePath, nil
}

func GetConfig() *IConfig {
	return Config
}
