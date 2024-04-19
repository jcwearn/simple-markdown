package webserver

import (
	"flag"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jcwearn/simple-markdown/internal/simpleparser"
)

type Webserver struct {
	simpleParser *simpleparser.SimpleParser
	logger       *slog.Logger
}

func NewWebServer(parser *simpleparser.SimpleParser) *Webserver {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &Webserver{
		simpleParser: parser,
		logger:       logger,
	}
}

func (ws *Webserver) Start() error {
	addr := flag.String("addr", ":4000", "HTTP network address")

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/parse", ws.parseHandler)

	srv := &http.Server{
		Addr:         *addr,
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

func (ws *Webserver) parseHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	output := ws.simpleParser.ParseInput(string(buf))
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(output))

	return
}
