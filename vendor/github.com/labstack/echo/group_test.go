package echo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Fix me
func TestGroup(t *testing.T) {
	g := New().Group("/group")
	h := func(Context) error { return nil }
	g.CONNECT("/", h)
	g.Connect("/", h)
	g.DELETE("/", h)
	g.Delete("/", h)
	g.GET("/", h)
	g.Get("/", h)
	g.HEAD("/", h)
	g.Head("/", h)
	g.OPTIONS("/", h)
	g.Options("/", h)
	g.PATCH("/", h)
	g.Patch("/", h)
	g.POST("/", h)
	g.Post("/", h)
	g.PUT("/", h)
	g.Put("/", h)
	g.TRACE("/", h)
	g.Trace("/", h)
	g.Any("/", h)
	g.Match([]string{GET, POST}, "/", h)
	g.Static("/static", "/tmp")
	g.File("/walle", "_fixture/images//walle.png")
}

func TestGroupRouteMiddleware(t *testing.T) {
	// Ensure middleware slices are not re-used
	e := New()
	g := e.Group("/group")
	h := func(Context) error { return nil }
	m1 := func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			return next(c)
		}
	}
	m2 := func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			return next(c)
		}
	}
	m3 := func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			return next(c)
		}
	}
	m4 := func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			return c.NoContent(404)
		}
	}
	m5 := func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			return c.NoContent(405)
		}
	}
	g.Use(m1, m2, m3)
	g.GET("/404", h, m4)
	g.GET("/405", h, m5)

	c, _ := request(GET, "/group/404", e)
	assert.Equal(t, 404, c)
	c, _ = request(GET, "/group/405", e)
	assert.Equal(t, 405, c)
}
