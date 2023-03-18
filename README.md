IPNotify
========

IPNotify is an open-source tool that monitors IPv4 address changes and sends email notifications in real-time. It's ideal for anyone who needs to stay informed about network changes and wants a simple, efficient way to receive alerts.

[![Go Version](https://img.shields.io/badge/Go-v1.16-blue)](https://golang.org/dl/)
[![Downloads](https://img.shields.io/github/downloads/PxGo/IPNotify/total)](https://github.com/PxGo/IPNotify/releases)
[![References](https://img.shields.io/github/forks/PxGo/IPNotify?label=references)](https://github.com/PxGo/IPNotify/network/members)


Usage
-----

### Installation

You can install IPNotify by compiling the source code, deploying it with Docker, or downloading the pre-compiled binary.

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
```

License
-------

IPNotify is released under the [MIT License](https://github.com/PxGo/IPNotify/blob/main/LICENSE). Feel free to use, modify, and distribute the software. Contributions are welcome!
