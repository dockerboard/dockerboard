package models

import (
	"crypto/tls"
	"net/url"
)

const (
	DEFAULT_HTTP_HOST   = "tcp://127.0.0.1"
	DEFAULT_UNIX_SOCKET = "unix:///var/run/docker.sock"
	DEFAULT_CA_FILE     = "ca.pem"
	DEFAULT_KEY_FILE    = "key.pem"
	DEFAULT_CERT_FILE   = "cert.pem"
)

type Host struct {
	Name        string
	Id          string
	URL         *url.URL
	TLSVerify   bool
	TLSCaFile   string
	TLSKeyFile  string
	TLSCertFile string
	TLSCertPath string
	TLSConfig   *tls.Config `json:"-"`
}

type Hosts []*Host