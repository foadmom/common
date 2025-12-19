// ==================================================================
// usage:
//
//	server instance = cHttp.Init (host, port)
//	loop and call AddHandler for all url mappings
//
// ==================================================================
package cHttp

type gHttp struct {
	Server server
}

var gHttpLogger gHttp
var gHttpServer gHttp = gHttp{}

// ==================================================================
//
// ==================================================================
// ==================================================================
// takes a single handler and adds it to the mapping and registers it
// ==================================================================
// ==================================================================
// Assuming Init has been called and handler(s) been established,
// this will
// ==================================================================
// func (s *server) listen() {
// 	httpLogger.Print(l.Trace, "listen: calling http listen")
// 	var _listeningTo string = fmt.Sprintf("%s:%s", s.host, s.Port)
// 	h.ListenAndServe(_listeningTo, nil)
// }

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
