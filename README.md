IPNotify
========

IPNotify is a simple and powerful open-source tool that monitors your network for IPv4 address changes and sends real-time email notifications. With support for Dynamic DNS updates with DNSPod and Cloudflare, IPNotify automates DNS updates and ensures your DNS records are always up-to-date. Keep yourself informed and stay on top of network changes with IPNotify.

[![Go Version](https://img.shields.io/badge/Go-v1.16-blue)](https://golang.org/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/PxGo/IPNotify)](https://goreportcard.com/report/github.com/PxGo/IPNotify)
[![Downloads](https://img.shields.io/github/downloads/PxGo/IPNotify/total)](https://github.com/PxGo/IPNotify/releases)
[![References](https://img.shields.io/github/forks/PxGo/IPNotify?label=references)](https://github.com/PxGo/IPNotify/network/members)
[![License](https://img.shields.io/github/license/PxGo/IPNotify)](https://github.com/PxGo/IPNotify/blob/main/LICENSE)


Usage
-----

### Installation

You can install IPNotify by compiling the source code or downloading the pre-compiled binary.

#### Compile from source code

1.  `git clone https://github.com/PxGo/IPNotify`
2.  `cd IPNotify`
3.  `go build .`

#### Download pre-compiled binary

Download the binary for your OS from the [releases page](https://github.com/PxGo/IPNotify/releases).

### Configuration

You can specify the path to your configuration file with the `-f` option:

bash
```bash
IPNotify -f /path/to/config.yaml
```

#### Configuration file

Here is an example of a configuration file with explanations:

yaml

```yaml
ip:
  # query_urls is a list of URLs used to query the external IP address.
  query_urls:
    - https://ipinfo.io/ip
    - https://icanhazip.com
    - https://checkip.amazonaws.com/
    - https://myexternalip.com/raw
    - https://ifconfig.me/ip
  # interval is a cron expression used to schedule the IP address query
  # at regular intervals.
  # Example:
  # Every minute: 0 * * * * *
  # Every 5 minutes: 0 */5 * * * *
  interval: "0 * * * * *"

email:
  # smtp_host is the hostname or IP address of the SMTP server used to send emails.
  smtp_host: mail.example.com
  # smtp_port is the port number of the SMTP server.
  smtp_port: 25
  # smtp_user is the username used to authenticate with the SMTP server.
  smtp_user: sender@example.com
  # smtp_passwd is the password used to authenticate with the SMTP server.
  smtp_passwd: passwd
  # from is the email address of the sender.
  from: sender@example.com
  # to is a list of email addresses to send the notifications to.
  to:
    - recipient1@example.com
    - recipient2@example.com
  
DDNS:
  # DNSPod is a DDNS provider that can be used to automatically update DNS records for a given domain and subdomain.
  DNSPod:
    # Set "enabled" to true to enable DDNS with DNSPod.
    enabled: false
    # "login_token" is the authentication token used to authenticate with the DNSPod API.
    login_token: "login token"
    # "records" is a list of DNS record information that needs to be updated.
    records:
      # Each record has the following fields:
      # - "domain_id": The domain ID for the domain to update the DNS record for.
      # - "record_id": The record ID for the DNS record to update.
      # - "sub_domain": The subdomain for the DNS record to update.
      # - "record_type": The type of the DNS record to update.
      # - "record_line": The line for the DNS record to update.
      # - "ttl": The time-to-live (TTL) for the DNS record in seconds.
      -
        domain_id: "your domain id"
        record_id: "your record id"
        sub_domain: "www"
        record_type: "A"
        record_line: "默认"
        ttl: 600  
        
  # Cloudflare is a DDNS provider that can be used to automatically update DNS records for a given domain and subdomain.
  Cloudflare:
    # Set "enabled" to true to enable DDNS with Cloudflare.
    enabled: false
    # "api_token" is the authentication token used to authenticate with the Cloudflare API.
    api_token: "your api token"
    # "records" is a list of DNS record information that needs to be updated.
    records:
      # Each record has the following fields:
      # - "zone_id": The ID for the zone to update the DNS record for.
      # - "record_id": The ID for the DNS record to update.
      # - "record_name": The name of the DNS record to update.
      # - "record_type": The type of the DNS record to update.
      # - "ttl": The time-to-live (TTL) for the DNS record in seconds.
      -
        zone_id: "your zone id"
        record_id: "your record id"
        record_name: "www.example.com"
        record_type: "A"
        ttl: 1
```

License
-------

IPNotify is released under the [MIT License](https://github.com/PxGo/IPNotify/blob/main/LICENSE). Feel free to use, modify, and distribute the software. Contributions are welcome!
