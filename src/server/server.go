package server

import (
	"configuration"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var urlRegExp = regexp.MustCompile("^([a-zA-Z0-9./_-]+)$")

// Server Represents the server's state.
type Server struct {
	config  configuration.Configuration
	logFile *os.File
}

// Initialize Initializes the server instance.
func (server *Server) Initialize(config configuration.Configuration) {

	validateConfiguration(config)
	server.config = config
}

// Start Starts the server.
func (server *Server) Start() {

	defer server.cleanUp()
	server.initializeLog()

	if server.config.UseBuiltInServer {
		server.startBuiltIn()
	} else {
		http.HandleFunc("/", server.handleRequest)
		http.ListenAndServe(":"+strconv.Itoa(server.config.Port), nil)
	}
}

func (server *Server) handleRequest(w http.ResponseWriter, r *http.Request) {

	logForRequest(r, "Incoming request.")

	if validateURL(w, r) {
		server.serveFile(w, r)
	}
}

func (server *Server) cleanUp() {

	if server.logFile != nil {
		server.logFile.Close()
	}
}

func (server *Server) initializeLog() {

	file, err := os.OpenFile(server.config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		server.logFile = nil
	}

	server.logFile = file
	log.SetOutput(server.logFile)
}

func (server *Server) serveFile(w http.ResponseWriter, r *http.Request) {

	sourcePath := server.config.ContentDirectoryPath + r.URL.Path
	if !checkIfFileExists(sourcePath) {
		logForRequest(r, "[ERROR 404] File not found.")
		http.NotFound(w, r)
		return
	}

	bs, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		logForRequest(r, "[ERROR 500] File cannot be read.")
		http.Error(w, "Internal server error", 500)
		return
	}

	mimeType := http.DetectContentType(bs)
	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(sourcePath)))

	// Use chunked transfer whenever possible. Go handles this itself, no need to set the "Transfer-Encoding" header.
	// Only set the content length if the request is using an earlier version of HTTP than 1.1.
	if r.ProtoMajor < 1 || r.ProtoMinor < 1 {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(bs)))
	}

	w.Write(bs)

	logForRequest(r, "Request served successfully.")
}

func logForRequest(r *http.Request, message string) {

	log.Println(fmt.Sprintf("%s | %s | %s", r.RemoteAddr, r.URL.Path, message))
}

func validateConfiguration(config configuration.Configuration) {

	if !checkIfDirExists(config.ContentDirectoryPath) {
		fmt.Println("The provided content directory does not exist.")
		os.Exit(1)
	}
}

func checkIfDirExists(path string) bool {

	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if !fileInfo.IsDir() {
		return false
	}

	return true
}

func checkIfFileExists(path string) bool {

	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if fileInfo.IsDir() {
		return false
	}

	return true
}

func validateURL(w http.ResponseWriter, r *http.Request) bool {

	match := urlRegExp.FindStringSubmatch(r.URL.Path)
	if match == nil {
		logForRequest(r, "[ERROR 404] Invalid URL.")
		http.NotFound(w, r)
		return false
	}

	return true
}

func (server *Server) startBuiltIn() {

	// router := http.NewServeMux()
	router := http.FileServer(http.Dir(server.config.ContentDirectoryPath))

	s := &http.Server{
		Addr:    ":" + strconv.Itoa(server.config.Port),
		Handler: server.wrapBuiltInHandler(router),
	}

	log.Fatal(s.ListenAndServe())
}

func (server *Server) wrapBuiltInHandler(f http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		logForRequest(r, "Incoming request.")
		record := &LogRecord{
			ResponseWriter: w,
		}

		f.ServeHTTP(record, r)

		logForRequest(r, strconv.Itoa(record.status))
	}
}
