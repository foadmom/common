package http

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	l "github.com/foadmom/common/logger"
)

type server struct {
	host string
	port string
	url  string
	// handler   func (http.ResponseWriter, *http.Request)
	processor func(string) (string, error)
}

var httpLogger l.LoggerInterface

func Init(host, port, url string, wg *sync.WaitGroup, pf func(string) (string, error)) {
	httpLogger = l.Instance()
	var _instance server = server{host, port, url, pf}
	_instance.listen()
	wg.Done()
}

// ============================================================================
//
// ============================================================================
func (m server) listen() {
	httpLogger.Print(l.Info, "Entering func listen()")
	http.HandleFunc(m.url, m.genericHandler)
	var _address string = ":" + m.port
	httpLogger.Printf(l.Trace, "calling http listen on port: %s and url %s", m.port, m.url)
	http.ListenAndServe(_address, nil)
	httpLogger.Print(l.Trace, "Exiting func listen()")
}

func (m server) genericHandler(w http.ResponseWriter, req *http.Request) {
	m.httpHandler(w, req)
}

// ============================================================================
//
// ============================================================================
func (m server) httpHandler(w http.ResponseWriter, req *http.Request) {
	var _jRequest string
	var _err error
	var _jResponse string

	_jRequest, _err = getJsonFromReq(req)
	if _err == nil {
		httpLogger.Printf(l.Debug, "Received json message: %s", _jRequest)
		_jResponse, _err = m.processor(_jRequest)
		if _err == nil {
			sendJsonResponse(w, _jResponse)
		} else {
			httpLogger.Printf(l.Error, "Error processing request: %s", _err.Error())
		}
	}
}

// ============================================================================
//
// ============================================================================
func getJsonFromReq(req *http.Request) (string, error) {
	var _err error
	var _jsonRequest []byte
	_jsonRequest, _err = io.ReadAll(req.Body)
	if _err != nil {
		httpLogger.Printf(l.Error, "error getting json payload from Body. Error = %s", _err.Error())
	}

	return string(_jsonRequest), _err
}

func sendJsonResponse(w http.ResponseWriter, resp string) {
	fmt.Fprintf(w, "%v\n", resp)
}

// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================
