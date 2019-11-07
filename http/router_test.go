package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertBaseRoute(t *testing.T, router *Router, method, path, name string, beforeHandlerLen int, afterHandlerLen int) {
	key := router.routeKey(method, path)
	assert.NotNil(t, router.routes[key])
	assert.IsType(t, &Route{}, router.routes[key])
	assert.Equal(t, beforeHandlerLen, len(router.routes[key].beforeHandlers))
	assert.Equal(t, afterHandlerLen, len(router.routes[key].afterHandlers))
	assert.Equal(t, name, router.routes[key].name)
}

func TestRouter_BaseRoute(t *testing.T) {
	router := New()
	router.GET("/gets/1", func(ctx *Context) {
		ctx.Write([]byte("Body"))
		ctx.Next()
	}).After(func(ctx *Context) {
		ctx.Write([]byte("After 1"))
		ctx.Next()
	}).After(func(ctx *Context) {
		ctx.Write([]byte("After 2"))
		ctx.Next()
	}).Before(func(ctx *Context) {
		ctx.Write([]byte("Before 1"))
		ctx.Next()
	}).Name("gets.1")

	assertBaseRoute(t, router, http.MethodGet, "/gets/1", "gets.1", 1, 2)

	router.POST("/posts", func(ctx *Context) {
		ctx.Write([]byte("Body"))
		ctx.Next()
	}).Name("posts.1")
	assertBaseRoute(t, router, http.MethodPost, "/posts", "posts.1", 0, 0)

	router.PUT("/resources/1/put", func(ctx *Context) {
		ctx.Write([]byte("Body"))
		ctx.Next()
	})
	assertBaseRoute(t, router, http.MethodPut, "/resources/1/put", "", 0, 0)

	router.DELETE("/1/delete", func(ctx *Context) {
		ctx.Write([]byte("Body"))
		ctx.Next()
	})
	assertBaseRoute(t, router, http.MethodDelete, "/1/delete", "", 0, 0)

	router.PATCH("/patch", func(ctx *Context) {
		ctx.Write([]byte("Body"))
		ctx.Next()
	}).Name("patch")
	assertBaseRoute(t, router, http.MethodPatch, "/patch", "patch", 0, 0)

	router.OPTIONS("/options", func(ctx *Context) {
		ctx.Write([]byte("Body"))
		ctx.Next()
	})
	assertBaseRoute(t, router, http.MethodOptions, "/options", "", 0, 0)

	//	http.ListenAndServe("127.0.0.1:28082",router)
}

func TestRouter_Group(t *testing.T) {
	router := New()
	v1 := router.Group("/v1").After(func(ctx *Context) {
		ctx.Write([]byte("Group v1 After"))
		ctx.Next()
	}).Before(ServerError, func(ctx *Context) {
		ctx.Write([]byte("Group v1 Before"))
		ctx.Next()
	})
	{
		v1.GET("/gets/1", func(ctx *Context) {
			ctx.Write([]byte("bdc"))
			ctx.Next()
		}).Name("gets.1")
		assertBaseRoute(t, router, http.MethodGet, "/v1/gets/1", "gets.1", 2, 1)

		v1.POST("/posts", func(ctx *Context) {
			ctx.Next()
		}).Name("v1.posts")
		assertBaseRoute(t, router, http.MethodPost, "/v1/posts", "v1.posts", 2, 1)

		//
		v1.DELETE("/delete", func(ctx *Context) {
		})
		assertBaseRoute(t, router, http.MethodDelete, "/v1/delete", "", 2, 1)

		v1.PUT("/put", func(ctx *Context) {
		})
		assertBaseRoute(t, router, http.MethodPut, "/v1/put", "", 2, 1)

		v1.PATCH("/patch", func(ctx *Context) {
		})
		assertBaseRoute(t, router, http.MethodPatch, "/v1/patch", "", 2, 1)

		v1.OPTIONS("/options", func(ctx *Context) {
		})
		assertBaseRoute(t, router, http.MethodOptions, "/v1/options", "", 2, 1)
	}

	v1Dep := v1.Group("/dep").Before(func(ctx *Context) {
		ctx.Write([]byte("Group v1--dep before"))
		ctx.Next()
	})
	{
		v1Dep.GET("/gets/1", func(ctx *Context) {

		})
	}
	assertBaseRoute(t, router, http.MethodGet, "/v1/dep/gets/1", "", 3, 1)
}

func TestRouter_Static(t *testing.T) {
	//http.Handle("/", http.FileServer(http.Dir("/tmp")))
	//http.ListenAndServe("127.0.0.1:28084", nil)
	router := New()
	router.Static("/file", "/tmp")
	router.GET("/gets/:name", func(ctx *Context) {
		ctx.Write([]byte(ctx.Param("name")))
		ctx.Next()
	})
	router.NotFound(func(ctx *Context) {
		ctx.Write([]byte("zzzz"))
		ctx.Next()
	})
	//router.GET("/gets/1", func(ctx *Context) {
	//	ctx.Write([]byte("Body"))
	//	ctx.Next()
	//}).After(func(ctx *Context) {
	//	ctx.Write([]byte("After 1"))
	//	ctx.Next()
	//}).After(func(ctx *Context) {
	//	ctx.Write([]byte("After 2"))
	//	ctx.Next()
	//}).Before(func(ctx *Context) {
	//	ctx.Write([]byte("Before 1"))
	//	ctx.Next()
	//}).Name("gets.1")
	//err := http.ListenAndServe("127.0.0.1:28084", router)
	//if err != nil {
	//	fmt.Println(err)
	//}
}