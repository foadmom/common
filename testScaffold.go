package main

import (
	"fmt"
	"net/http"
	"time"

	h "github.com/foadmom/common/cHttp"
	l "github.com/foadmom/common/logger"
)

var _logger l.LoggerInterface

func main() {
	TestcHttp()
}

func TestcHttp() {
	_logger = l.Instance()
	var _http *h.CommonHttp
	_http = _http.Init("localhost", "8001")
	h.AddHandler("/", localHttpHandler)
	_http.Listen()

	_logger.Print(l.Trace, "exiting testScaffold")
}

func localHttpHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s Hello\n", time.Now().String())
}
