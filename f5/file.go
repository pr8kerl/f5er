package f5

import (
	"net/http"
	"crypto/tls"
	"fmt"
	"errors"
	"bytes"
	"strconv"
)

type StatusError struct {

}

func (f *Device) UploadFile(filename string, data []byte) (error) {
	//	err, res := f.sendRequest("/mgmt/shared/file-transfer/uploads/" + filename, POSTR, bytes.NewBuffer(data), resp)
	if(len(data) > 512 * 1024) {
		return errors.New("File size is too large, and we dont support chunked file sizes yet.")
	}

	var url string = f.Proto + "://" + f.Hostname + "/mgmt/shared/file-transfer/uploads/" + filename
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	request.SetBasicAuth(f.Username, f.Password)
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("Content-Range", "0-" + strconv.Itoa(len(data) - 1) + "/" + strconv.Itoa(len(data)))
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode > 299 || response.StatusCode < 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		s := buf.String()
		return errors.New("Unable to process request, returned status: " + response.Status + " " + s)
	}
	return nil
}
