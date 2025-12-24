// ==================================================================
// usage:
//
//	server instance = cHttp.Init (host, port)
//	loop and call AddHandler for all url mappings
//
// ==================================================================
package http

import (
	"fmt"
	h "net/http"

	l "github.com/foadmom/common/logger"
)

type server struct {
	Host    string
	Port    string
	URL     string
	Mapping map[string]func(h.ResponseWriter, *h.Request)
}

type httpInterface interface {
	Init(host, port string) *httpInterface
	// AddHandler(url string, handler func(h.ResponseWriter, *h.Request))
	// AddHandlers(m map[string]func(h.ResponseWriter, *h.Request))
	Listen()
}

type CommonHttp struct {
	Server server
}

var httpLogger l.LoggerInterface
var cHttpServer CommonHttp = CommonHttp{}

// ==================================================================
//
// ==================================================================
func init() {
	httpLogger = l.Instance()
}

// ==================================================================
//
// ==================================================================
func (s *CommonHttp) Init(host, port string) *CommonHttp {
	cHttpServer.Server.Mapping = make(map[string]func(h.ResponseWriter, *h.Request))
	cHttpServer.Server.Host = host
	cHttpServer.Server.Port = port
	return &cHttpServer
}

// ==================================================================
// takes a single handler and adds it to the mapping and registers it
// ==================================================================
// func (s *CommonHttp) AddHandler(url string, handler func(h.ResponseWriter, *h.Request)) {
func AddHandler(url string, handler func(h.ResponseWriter, *h.Request)) {
	httpLogger.Printf(l.Trace, "AddHandler: adding handler for %s", url)
	cHttpServer.Server.Mapping[url] = handler
	// h.HandleFunc(url, handler)
}

// ==================================================================
// takes all the handlers as a map and registers them
// ==================================================================
// func (s *CommonHttp) AddHandlers(m map[string]func(h.ResponseWriter, *h.Request)) {
func AddHandlers(m map[string]func(h.ResponseWriter, *h.Request)) {
	for _url, _handler := range m {
		AddHandler(_url, _handler)
	}
}

// ==================================================================
// Assuming Init has been called and handler(s) been established,
// this will
// ==================================================================
func (s *CommonHttp) Listen() {
	httpLogger.Print(l.Trace, "listen: calling http listen")
	for _url, _handler := range s.Server.Mapping {
		h.HandleFunc(_url, _handler)
	}
	var _listeningTo string = fmt.Sprintf("%s:%s", s.Server.Host, s.Server.Port)
	h.ListenAndServe(_listeningTo, nil)
}

// ==================================================================
//
// ==================================================================
// func (s *server) genericHandler(w http.ResponseWriter, req *http.Request) {
// 	s.httpHandler(w, req)
// }

// ==================================================================
//
// ==================================================================
// func (s *server) httpHandler(w http.ResponseWriter, req *http.Request) {
// 	var _jRequest string
// 	var _err error
// 	var _jResponse string

// 	_jRequest, _err = getJsonFromReq(req)
// 	if _err == nil {
// 		httpLogger.Printf(l.Trace, "Received json message: %s", _jRequest)
// 		_jResponse, _err = s.processor(_jRequest)
// 		if _err == nil {
// 			sendJsonResponse(w, _jResponse)
// 		} else {
// 			httpLogger.Printf(l.Error, "Error processing request: %s", _err.Error())
// 		}
// 	}
// }

// ==================================================================
//
// ==================================================================
// func getJsonFromReq(req *http.Request) (string, error) {
// 	var _err error
// 	var _jsonRequest []byte
// 	_jsonRequest, _err = io.ReadAll(req.Body)
// 	if _err != nil {
// 		httpLogger.Printf(l.Error, "error getting json payload from Body. Error = %s", _err.Error())
// 	}

// 	return string(_jsonRequest), _err
// }

// func sendJsonResponse(w http.ResponseWriter, resp string) {
// 	fmt.Fprintf(w, "%v\n", resp)
// }

// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
