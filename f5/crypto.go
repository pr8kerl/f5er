package f5

import (
	"errors"
	"strings"
)

type SSLCreate struct {
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	SourcePath string `json:"sourcePath"`
}

type SSLCertificate struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	Generation int    `json:"generation"`
	SelfLink   string `json:"selfLink"`
	CurveName  string `json:"certificateKeyCurveName"`
	KeySize    int    `json:"certificateKeySize"`
	Checksum   string `json:"checksum"`
	CreateTime string `json:"createTime"`
	CreatedBy  string `json:"createdBy"`
	Expiration int    `json:"expirationDate"`
	ExpireTime string `json:"expirationString"`
	IsBundle   string `json:"isBundle"`
	Issuer     string `json:"issuer"`
	KeyType    string `json:"keyType"`
	UpdateTime string `json:"lastUpdateTime"`
	Mode       int    `json:"mode"`
	Revision   int    `json:"revision"`
	SerialNum  string `json:"serialNumber"`
	Size       int    `json:"size"`
	Subject    string `json:"subject"`
	UpdatedBy  string `json:"updatedBy"`
	Version    int    `json:"version"`
}

type SSLCertificates struct {
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selfLink"`
	Items    []SSLCertificate `json:"items"`
}

func (f *Device) GetCertificate(partition string, name string) (error, *SSLCertificate) {
	if !strings.HasSuffix(name, ".crt") {
		name = name + ".crt"
	}
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/sys/file/ssl-cert/~" + partition + "~" + name
	res := SSLCertificate{}
	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}
}

func (f *Device) GetCertificates() (error, *SSLCertificates) {
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/sys/file/ssl-cert"
	res := SSLCertificates{}
	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}
}

func (f *Device) CreateCertificateFromLocalFile(name string, partition string, cert_file string) (error, *SSLCertificate) {
	if !strings.HasSuffix(name, ".crt") {
		name = name + ".crt"
	}
	b := SSLCreate{Name: name, Partition: partition, SourcePath: "file:///var/config/rest/downloads/" + cert_file}
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/sys/file/ssl-cert"
	res := SSLCertificate{}

	err, _ := f.sendRequest(u, POST, b, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) CreateKeyFromLocalFile(name string, partition string, key_file string) (error, *SSLCertificate) {
	if strings.HasSuffix(name, ".crt") {
		return errors.New("The name cannot contain a .crt suffix for keys."), nil
	}
	if !strings.HasSuffix(name, ".key") {
		name = name + ".key"
	}
	b := SSLCreate{Name: name, Partition: partition, SourcePath: "file:///var/config/rest/downloads/" + key_file}
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/sys/file/ssl-key"
	res := SSLCertificate{}

	err, _ := f.sendRequest(u, POST, &b, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}
}
