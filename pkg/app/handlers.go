package app

import (
	"encoding/json"
	"fmt"
	"healthchecker/pkg/collector"
	"html/template"
	"net/http"
)

func (s *Server) Watch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		network := r.URL.Query().Get("network")
		if network == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("target blockchain didn't specified"))
			return
		}

		latestBlock, err := s.watchService.GetLatestBlock(network)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("failed to get %s's latest block", network)))
			return
		}

		raw, err := json.Marshal(latestBlock)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to marshal response to byte slice"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	}
}

func (s *Server) ServeIndexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ethLatestBlock, err := s.watchService.GetLatestBlock("ethereum")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to get klaytn's latest block"))
			return
		}

		KlayLatestBlock, err := s.watchService.GetLatestBlock("klaytn")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to get klaytn's latest block"))
			return
		}

		t, err := template.ParseFiles("web/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to parse html file"))
		}

		t.Execute(w, map[string]struct {
			Network     string
			LatestBlock *collector.LatestBlock
		}{
			"ethereum": {
				Network:     "Ethereum",
				LatestBlock: ethLatestBlock,
			},
			"klaytn": {
				Network:     "Klaytn",
				LatestBlock: KlayLatestBlock,
			},
		})
	}
}
