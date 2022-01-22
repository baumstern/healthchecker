package app

import (
	"encoding/json"
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
			w.Write([]byte("failed to get klaytn's latest block"))
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
