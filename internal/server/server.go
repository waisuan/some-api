package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"some-api/internal/location"
	"some-api/utils/db"
)

type Server struct {
	dataStore db.DataStore
	Router *echo.Echo
}

func NewServer(dbClient db.DataStore) *Server {
	e := echo.New()
	a := &Server{
		dataStore: dbClient,
		Router: e,
	}
	e.GET("/location/:id", a.getLocation)
	return a
}

func (s *Server) getLocation(c echo.Context) error {
	personId := c.Param("id")
	out, err := location.Get(s.dataStore, personId)
	if err != nil {
		c.Error(err)
	}
	return c.JSON(http.StatusOK, out)
}