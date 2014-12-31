package controllers

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/dockerboard/dockerboard/app/models"
	"github.com/go-libs/quest"
)

var localhost = &models.Host{Name: "Local"}
var hosts = models.Hosts{localhost}

func NewRequest(method, endpoint, host string) (q *quest.Requester, err error) {
	var (
		h *models.Host
	)
	if host == "" {
		h = localhost
	} else {
		h, _, _, err = GetHost(host)
		if err != nil {
			return
		}
		if h == nil {
			h = localhost
		}
	}
	h.URL.Path = endpoint
	q, err = quest.Request(method, h.URL.String())
	if err != nil {
		return
	}
	// Must Overwrite Url for unix
	q.Url = h.URL
	if h.TLSConfig != nil {
		q.TLSConfig(h.TLSConfig)
	}
	return
}

func NewContainers() *ContainersController {
	return &ContainersController{}
}

func NewContainerActions() *ContainerActionsController {
	return &ContainerActionsController{}
}

func NewImages() *ImagesController {
	return &ImagesController{}
}

func NewImageActions() *ImageActionsController {
	return &ImageActionsController{}
}

func NewSystem() *SystemController {
	return &SystemController{}
}

func NewHosts() *HostsController {
	return &HostsController{}
}

func NewHostActions() *HostActionsController {
	return &HostActionsController{}
}

func NewApps() *AppsController {
	return &AppsController{}
}

func GetTLSConfig(path string, insecure bool) (t *tls.Config, err error) {
	t = &tls.Config{}
	if !insecure {
		var file []byte
		file, err = ioutil.ReadFile(path + "/ca.pem")
		if err != nil {
			return
		}
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(file)
		t.RootCAs = certPool
	}
	cert, err := tls.LoadX509KeyPair(path+"/cert.pem", path+"/key.pem")
	if err != nil {
		return
	}
	t.Certificates = []tls.Certificate{cert}
	t.InsecureSkipVerify = insecure
	// Avoid fallback to SSL protocols < TLS1.0
	t.MinVersion = tls.VersionTLS10
	return
}

func ParseURL(addr string) (u *url.URL, err error) {
	u, err = url.Parse(addr)
	if err != nil {
		return
	}
	if u.Scheme == "unix" {
		u.Host = u.Path
		u.Path = ""
	} else if u.Scheme == "tcp" {
		u.Scheme = "http"
	} else if u.Scheme == "" {
		u.Scheme = "http"
		u.Host = u.Path
		u.Path = ""
	}
	return
}

func GetHost(addr string) (*models.Host, int, *url.URL, error) {
	u, err := ParseURL(addr)
	if err != nil {
		return nil, 0, nil, err
	}
	for i, h := range hosts {
		if h.URL.Scheme == u.Scheme && h.URL.Host == u.Host {
			return h, i, u, nil
		}
	}
	return nil, 0, u, errors.New("Not Found Host.")
}

func init() {
	host := os.Getenv("DOCKER_HOST")
	certPath := os.Getenv("DOCKER_CERT_PATH")
	tlsVerify := os.Getenv("DOCKER_TLS_VERIFY")

	if host == "" {
		host = models.DEFAULT_UNIX_SOCKET
	}

	u, err := ParseURL(host)
	if err == nil {
		localhost.URL = u
		if certPath != "" && u.Scheme != "unix" {
			u.Scheme = "https"
			localhost.TLSVerify = tlsVerify == "1"
			localhost.TLSCertPath = certPath
			localhost.TLSCaFile = models.DEFAULT_CA_FILE
			localhost.TLSKeyFile = models.DEFAULT_KEY_FILE
			localhost.TLSCertFile = models.DEFAULT_CERT_FILE
			if TLSClientConfig, err := GetTLSConfig(certPath, localhost.TLSVerify); err == nil {
				localhost.TLSConfig = TLSClientConfig
			}
		}
	}
}
