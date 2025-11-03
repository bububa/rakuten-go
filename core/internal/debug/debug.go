package debug

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bububa/rakuten-go/util"
)

// PrintError print error with debug
func PrintError(err error, debug bool) {
	if !debug {
		return
	}
	log.Println("[TIKTOK] [ERROR]", err)
}

// PrintStringResponse print string response with debug
func PrintStringResponse(str string, debug bool) {
	if !debug {
		return
	}
	log.Println("[TIKTOK] [RESP]", str)
}

// PrintGetRequest print get request with debug
func PrintGetRequest(url string, debug bool) {
	if !debug {
		return
	}
	log.Println("[TIKTOK] [REQ] GET", url)
}

// PrintPostJSONRequest print json request with debug
func PrintPostJSONRequest(url string, body []byte, debug bool) {
	if !debug {
		return
	}
	const format = "[TIKTOK] [REQ] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := util.NewBuffer()
	defer util.ReleaseBuffer(buf)
	json.Indent(buf, body, "", "    ")
	log.Printf(format, url, buf.String())
}

// PrintPostFormRequest print json request with debug
func PrintPostFormRequest(url string, body string, debug bool) {
	if !debug {
		return
	}
	const format = "[TIKTOK] [REQ] FORM POST %s\n" +
		"http request body:\n%s\n"

	log.Printf(format, url, body)
}

// PrintJSONRequest print json request with debug
func PrintJSONRequest(method string, url string, header http.Header, body []byte, debug bool) {
	if !debug {
		return
	}

	const format = "[TIKTOK] [REQ] JSON %s %s\n" +
		"http request header:\n%s\n" +
		"http request body:\n%s\n"

	buf := util.NewBuffer()
	defer util.ReleaseBuffer(buf)
	json.Indent(buf, body, "", "\t")
	headers := make([]string, 0, len(header))
	for k := range header {
		headers = append(headers, util.StringsJoin(k, ": ", header.Get(k)))
	}

	log.Printf(format, method, url, strings.Join(headers, "\n"), buf.String())
}

// PrintFormRequest print json request with debug
func PrintFormRequest(method string, url string, header http.Header, body []byte, debug bool) {
	if !debug {
		return
	}

	const format = "[TIKTOK] [REQ] FORM %s %s\n" +
		"http request header:\n%s\n" +
		"http request body:\n%s\n"

	buf := util.NewBuffer()
	defer util.ReleaseBuffer(buf)
	json.Indent(buf, body, "", "\t")
	headers := make([]string, 0, len(header))
	for k := range header {
		headers = append(headers, util.StringsJoin(k, ": ", header.Get(k)))
	}

	log.Printf(format, method, url, strings.Join(headers, "\n"), buf.String())
}

// PrintPostMultipartRequest print multipart/form-data post request with debug
func PrintPostMultipartRequest(url string, mp map[string]string, debug bool) {
	if !debug {
		return
	}
	const format = "[TIKTOK] [REQ] multipart/form-data POST %s\n" +
		"http request body:\n%s\n"

	bs, _ := json.MarshalIndent(mp, "", "\t")
	log.Printf(format, url, bs)
}

// DecodeJSONHttpResponse decode json response with debug
func DecodeJSONHttpResponse(r io.Reader, w *bytes.Buffer, v any, debug bool) error {
	tee := io.TeeReader(r, w)
	if err := json.NewDecoder(tee).Decode(v); err != nil {
		return err
	}
	if !debug {
		return nil
	}
	debugBuf := util.NewBuffer()
	defer util.ReleaseBuffer(debugBuf)
	if err := json.Indent(debugBuf, w.Bytes(), "", "\t"); err != nil {
		return err
	}

	log.Println(util.StringsJoin("[TIKTOK] [RESP] http response body:\n", debugBuf.String()))
	return nil
}

// DecodeXMLHttpResponse decode json response with debug
func DecodeXMLHttpResponse(r io.Reader, w *bytes.Buffer, v any, debug bool) error {
	tee := io.TeeReader(r, w)
	if err := xml.NewDecoder(tee).Decode(v); err != nil {
		return err
	}
	if !debug {
		return nil
	}
	debugBuf := util.NewBuffer()
	defer util.ReleaseBuffer(debugBuf)
	enc := xml.NewEncoder(debugBuf)
	enc.Indent("", "\t")
	if err := enc.Encode(v); err != nil {
		return err
	}

	log.Println(util.StringsJoin("[TIKTOK] [RESP] http response body:\n", debugBuf.String()))
	return nil
}
