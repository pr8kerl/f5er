package f5

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmcvetta/napping"
	"log"
	"net/http"
	"net/url"
)

var (
	//sessn   napping.Session
	tsport  http.Transport
	clnt    http.Client
	headers http.Header
	debug   bool
)

const (
	GET = iota
	POST
	PUT
	PATCH
	DELETE
)

type httperr struct {
	Message string
	Errors  []struct {
		Resource string
		Field    string
		Code     string
	}
}

type Device struct {
	Hostname string
	Username string
	Password string
	Session  napping.Session
}

type Response struct {
	Status  int
	Message string
}

type LBEmptyBody struct{}

type LBTransaction struct {
	TransId int    `json:"transId"`
	Timeout int    `json:"timeoutSeconds"`
	State   string `json:"state"`
}

type LBTransactionState struct {
	State string `json:"state"`
}

func New(host string, username string, pwd string) *Device {
	f := Device{Hostname: host, Username: username, Password: pwd}
	f.InitSession()
	return &f
}

func (f *Device) InitSession() {

	// REST connection setup
	tsport = http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clnt = http.Client{Transport: &tsport}
	headers = make(http.Header)

	//
	// Setup HTTP Basic auth for this session (ONLY use this with SSL).  Auth
	// can also be configured on a per-request basis when using Send().
	//
	f.Session = napping.Session{
		Client:   &clnt,
		Log:      debug,
		Userinfo: url.UserPassword(f.Username, f.Password),
		Header:   &headers,
	}

}

func (f *Device) StartTransaction() (error, string) {

	u := "https://" + f.Hostname + "/mgmt/tm/transaction"
	empty := LBEmptyBody{}
	tres := LBTransaction{}
	err, _ := f.sendRequest(u, POST, &empty, &tres)
	if err != nil {
		return err, ""
	}

	tid := fmt.Sprintf("%d", tres.TransId)
	// set the transaction header
	f.Session.Header.Set("X-F5-REST-Coordination-Id", tid)
	return nil, tid

}

func (f *Device) CommitTransaction(tid string) error {

	// remove the transaction header first
	f.Session.Header.Del("X-F5-REST-Coordination-Id")

	u := "https://" + f.Hostname + "/mgmt/tm/transaction/" + tid
	body := LBTransaction{State: "VALIDATING"}
	tres := LBTransaction{}
	err, _ := f.sendRequest(u, PATCH, &body, &tres)
	if err != nil {
		return err
	}

	return nil

}

func (f *Device) sendRequest(u string, method int, pload interface{}, res interface{}) (error, *Response) {

	//
	// Send request to server
	//
	e := httperr{}
	var (
		err   error
		nresp *napping.Response
	)
	f.Session.Log = debug

	switch method {
	case GET:
		nresp, err = f.Session.Get(u, nil, &res, &e)
	case POST:
		nresp, err = f.Session.Post(u, &pload, &res, &e)
	case PUT:
		nresp, err = f.Session.Put(u, &pload, &res, &e)
	case PATCH:
		nresp, err = f.Session.Patch(u, &pload, &res, &e)
	case DELETE:
		nresp, err = f.Session.Delete(u, nil, &res, &e)
	}

	var resp = Response{Status: nresp.Status(), Message: e.Message}
	if err != nil {
		return err, &resp
	}
	if nresp.Status() == 401 {
		return errors.New("unauthorised - check your username and passwd"), &resp
	}
	if nresp.Status() >= 300 {
		return errors.New(e.Message), &resp
	} else {
		// all is good in the world
		return nil, &resp
	}
}

func (f *Device) PrintObject(input interface{}) {

	jsonresp, err := json.MarshalIndent(&input, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jsonresp))

}

// F5 Module data struct
// to show all available modules when using show without args
type LBModule struct {
	Link string `json:"link"`
}

type LBModuleRef struct {
	Reference LBModule `json:"reference"`
}

type LBModules struct {
	Items []LBModuleRef `json:"items"`
}
