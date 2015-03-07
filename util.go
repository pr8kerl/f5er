package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmcvetta/napping"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	sessn   napping.Session
	tlscfg  tls.Config
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

func InitSession() {

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
	sessn = napping.Session{
		Client:   &clnt,
		Log:      debug,
		Userinfo: url.UserPassword(username, passwd),
		Header:   &headers,
	}

}

func SendRequest(u string, method int, sess *napping.Session, pload interface{}, res interface{}) (error, *napping.Response) {

	//
	// Send request to server
	//
	e := httperr{}
	var (
		err  error
		resp *napping.Response
	)
	sess.Log = debug

	switch method {
	case GET:
		resp, err = sess.Get(u, nil, &res, &e)
	case POST:
		resp, err = sess.Post(u, &pload, &res, &e)
	case PUT:
		resp, err = sess.Put(u, &pload, &res, &e)
	case PATCH:
		resp, err = sess.Patch(u, &pload, &res, &e)
	case DELETE:
		resp, err = sess.Delete(u, &res, &e)
	}

	if err != nil {
		return err, resp
	}
	if resp.Status() == 401 {
		return errors.New("unauthorised - check your username and passwd"), resp
	}
	if resp.Status() >= 300 {
		return errors.New(e.Message), resp
	} else {
		// all is good in the world
		return nil, resp
	}
}

/*
func prettifyString(string input) (string, error) {

	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	tabs := 0

	for scanner.Scan() {
		tok := scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", tabs)

}
*/
func prettifyScanner(input string) {

	printtabs := func(m int) {
		for i := 0; i < m; i++ {
			fmt.Printf("\t")
		}
	}

	tabs := 0
	open := false
	for _, tok := range input {

		switch {
		case tok == '"':
			if open {
				open = false
			} else {
				open = true
			}
		case tok == '{':
			if !open {
				fmt.Printf("\n")
				printtabs(tabs)
				tabs++
				fmt.Println(string(tok))
				printtabs(tabs)
			} else {
				fmt.Printf("%s", string(tok))
			}
		case tok == '}':
			if !open {
				fmt.Printf("\n")
				tabs--
				printtabs(tabs)
				fmt.Printf("%s", string(tok))
			} else {
				fmt.Printf("%s", string(tok))
			}
		case tok == '[':
			if !open {
				fmt.Printf("\n")
				printtabs(tabs)
				tabs++
				fmt.Println(string(tok))
				printtabs(tabs)
			} else {
				fmt.Printf("%s", string(tok))
			}
		case tok == ']':
			if !open {
				fmt.Printf("\n")
				tabs--
				printtabs(tabs)
				fmt.Printf("%s", string(tok))
			} else {
				fmt.Printf("%s", string(tok))
			}
		case tok == ',':
			fmt.Println(string(tok))
			printtabs(tabs)
		case tok == '\n':
		default:
			fmt.Printf("%s", string(tok))
		}
	}
	fmt.Println()

}

func prettifyBytes(input string) {

	f := func(c rune) bool {
		return (c == '{' || c == '}')
	}
	substrings := strings.FieldsFunc(input, f)
	for _, v := range substrings {
		fmt.Printf("pretty: %s\n", v)
	}

}

func printResponse(input interface{}) {

	jsonresp, err := json.MarshalIndent(&input, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonresp))

}

func bail(msg string) {
	log.SetFlags(0)
	log.Fatalf("\n%s\n\n", msg)
}
