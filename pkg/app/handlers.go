package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) Watch(c echo.Context) error {
	network := c.QueryParam("network")
	if network == "" {
		return c.String(http.StatusBadRequest, "target blockchain didn't specified")
	}

	latestBlock, err := s.watchService.GetLatestBlock(network)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get %s's latest block", network))
	}
	return c.JSON(http.StatusOK, latestBlock)
}
