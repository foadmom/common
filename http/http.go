package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	nl "nex/common/nexLogger"
	"sync"
)

type server struct {
	host string
	port string
	url  string
	// handler   func (http.ResponseWriter, *http.Request)
	processor func(string) (string, error)
}

func Init(host, port, url string, wg *sync.WaitGroup, pf func(string) (string, error)) {
	var _instance server = server{host, port, url, pf}
	_instance.listen()
	wg.Done()
}

// ============================================================================
//
// ============================================================================
func (m server) listen() {
	nl.Instance().Debug().Msg("Entering func listen()")
	http.HandleFunc(m.url, m.genericHandler)
	var _address string = ":" + m.port
	nl.Instance().Debug().Msg("calling http listen on port:" + m.port + " and url" + m.url)
	http.ListenAndServe(_address, nil)
	nl.Instance().Debug().Msg("Exiting func listen()")
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
		nl.Instance().Debug().Msg("Received json message: " + _jRequest)
		_jResponse, _err = m.processor(_jRequest)
		if _err == nil {
			sendJsonResponse(w, _jResponse)
		} else {
			nl.Instance().Error().Err(_err)
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
		log.Printf("error getting json payload from Body. Error = %v\n", _err)
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
