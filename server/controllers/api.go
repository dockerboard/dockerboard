package controllers

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/go-libs/quest"
)

type DockerOptions struct {
	Url             *url.URL
	CertPath        string
	TLSClientConfig *tls.Config
}

const UNIX_SOCKET = "unix:///var/run/docker.sock"

var dockerOptions = &DockerOptions{}

func NewRequest(method quest.Method, endpoint string) (q *quest.Requester, err error) {
	q, err = quest.Request(method, dockerOptions.Url.String()+endpoint)
	if err != nil {
		return
	}
	q.TLSConfig(dockerOptions.TLSClientConfig)
	return
}

func NewContainers() *ContainersController {
	return &ContainersController{}
}

func NewImages() *ImagesController {
	return &ImagesController{}
}

func NewApps() *AppsController {
	return &AppsController{}
}

func GetTLSConfig(path string, insecure bool) (t *tls.Config, err error) {
	certPool := x509.NewCertPool()
	if !insecure {
		var file []byte
		file, err = ioutil.ReadFile(path + "/ca.pem")
		if err != nil {
			return
		}
		certPool.AppendCertsFromPEM(file)
	}
	cert, err := tls.LoadX509KeyPair(path+"/cert.pem", path+"/key.pem")
	if err != nil {
		return
	}
	t = &tls.Config{
		RootCAs:            certPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: insecure,
	}
	return
}

func init() {
	host := os.Getenv("DOCKER_HOST")
	certPath := os.Getenv("DOCKER_CERT_PATH")
	tlsVerify := os.Getenv("DOCKER_TLS_VERIFY")

	if host == "" {
		host = UNIX_SOCKET
	}

	u, err := url.Parse(host)
	if err == nil {
		dockerOptions.Url = u
		if u.Scheme == "tcp" {
			u.Scheme = "http"
		}
		if certPath != "" {
			u.Scheme = "https"
			dockerOptions.CertPath = certPath
			TLSClientConfig, _ := GetTLSConfig(certPath, tlsVerify == "1")
			if TLSClientConfig != nil {
				dockerOptions.TLSClientConfig = TLSClientConfig
			}
		}
	}
}
