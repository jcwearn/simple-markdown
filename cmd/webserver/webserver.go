package webserver

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jcwearn/simple-markdown/internal/parser/pegparser"
	"github.com/jcwearn/simple-markdown/internal/parser/simpleparser"
)

type (
	WebServerConfig struct {
		Address      string
		SimpleParser simpleparser.SimpleParser
		PegParser    pegparser.PegParser
	}
	WebServer struct {
		address      string
		simpleParser simpleparser.SimpleParser
		pegParser    pegparser.PegParser
		logger       *slog.Logger
	}
)

func NewWebServer(cfg WebServerConfig) *WebServer {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &WebServer{
		address:      cfg.Address,
		simpleParser: cfg.SimpleParser,
		pegParser:    cfg.PegParser,
		logger:       logger,
	}
}

func (ws *WebServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/parse", ws.parseHandler)

	srv := &http.Server{
		Addr:         ws.address,
		Handler:      mux,
		ErrorLog:     slog.NewLogLogger(ws.logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ws.logger.Info("starting server", "addr", srv.Addr)

	err := srv.ListenAndServe()
	ws.logger.Error(err.Error())

	return err
}

func (ws *WebServer) parseHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	queryParams := r.URL.Query()
	parserQueryParam := queryParams.Get("parser")

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var output string
	if parserQueryParam == "peg" {
		output, err = ws.pegParser.ParseInput(string(buf))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		output, _ = ws.simpleParser.ParseInput(string(buf))
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(output))

	return
}
