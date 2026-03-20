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
// initializes the logger for this package.
// This will be called when the package is imported.
// ==================================================================
func init() {
	httpLogger = l.Instance()
}

// ==================================================================
// initializes the server with the given host and port.
// It also initializes the mapping for url handlers.
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
func AddHandler(url string, handler func(h.ResponseWriter, *h.Request)) {
	httpLogger.Printf(l.Trace, "AddHandler: adding handler for %s", url)
	cHttpServer.Server.Mapping[url] = handler
}

// ==================================================================
// takes all the handlers as a map and registers them
// ==================================================================
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
