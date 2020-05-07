package http

import (
	"github.com/firmeve/firmeve/kernel/contract"
	render2 "github.com/firmeve/firmeve/render"
	testing2 "github.com/firmeve/firmeve/testing"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

var configPath = "../testdata/config/config.yaml"

type MockResponseWriter struct {
	mock.Mock
	Bytes      []byte
	StatusCode int
	Headers    http.Header
}

func (m *MockResponseWriter) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}
func (m *MockResponseWriter) Header() http.Header {
	return m.Headers
}

func (m *MockResponseWriter) Write(p []byte) (int, error) {
	m.Bytes = p
	return len(p), nil
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

//func TestSomethingElse(t *testing.T) {
//
//	// create an instance of our test object
//	testObj := new(MockResponseWriter)
//
//	// setup expectations with a placeholder in the argument list
//	testObj.On("DoSomething", mock.Anything).Return(true, nil)
//
//	// call the code we are testing
//	targetFuncThatDoesSomethingWithObj(testObj)
//
//	// assert that the expectations were met
//	testObj.AssertExpectations(t)
//
//}

func assertBaseRoute(t *testing.T, router *Router, method, path, name string, beforeHandlerLen int, afterHandlerLen int) {
	key := router.routeKey(method, path)
	assert.NotNil(t, router.routes[key])
	assert.IsType(t, &Route{}, router.routes[key])
	assert.Equal(t, beforeHandlerLen, len(router.routes[key].(*Route).beforeHandlers))
	assert.Equal(t, afterHandlerLen, len(router.routes[key].(*Route).afterHandlers))
	assert.Equal(t, name, router.routes[key].(*Route).name)
}

func TestRouter_BaseRoute(t *testing.T) {
	router := New(testing2.ApplicationDefault())
	router.GET("/gets/1", func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Body")
		ctx.Next()
	}).After(func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "After 1")
		ctx.Next()
	}).After(func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "After 2")
		ctx.Next()
	}).Before(func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Before 1")
		ctx.Next()
	}).Name("gets.1")

	assertBaseRoute(t, router.(*Router), http.MethodGet, "/gets/1", "gets.1", 1, 2)

	router.POST("/posts", func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Body")
		ctx.Next()
	}).Name("posts.1")
	assertBaseRoute(t, router.(*Router), http.MethodPost, "/posts", "posts.1", 0, 0)

	router.PUT("/resources/1/put", func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Body")
		ctx.Next()
	})
	assertBaseRoute(t, router.(*Router), http.MethodPut, "/resources/1/put", "", 0, 0)

	router.DELETE("/1/delete", func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Body")
		ctx.Next()
	})
	assertBaseRoute(t, router.(*Router), http.MethodDelete, "/1/delete", "", 0, 0)

	router.PATCH("/patch", func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Body")
		ctx.Next()
	}).Name("patch")
	assertBaseRoute(t, router.(*Router), http.MethodPatch, "/patch", "patch", 0, 0)

	router.OPTIONS("/options", func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Body")
		ctx.Next()
	})
	assertBaseRoute(t, router.(*Router), http.MethodOptions, "/options", "", 0, 0)

	router.Handler("GET", "/original", func(writer http.ResponseWriter, request *http.Request) {

	})
	assertBaseRoute(t, router.(*Router), http.MethodGet, "/original", "", 0, 0)

	//	http.ListenAndServe("127.0.0.1:28082",router)
}

func TestRouter_HttpRouter(t *testing.T) {
	router := New(testing2.ApplicationDefault())
	assert.IsType(t, &httprouter.Router{}, router.HttpRouter())
}

func TestRouter_Group(t *testing.T) {
	router := New(testing2.ApplicationDefault())
	v1 := router.Group("/v1").After(func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Group v1 After")
		ctx.Next()
	}).Before(Recovery, func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Group v1 Before")
		ctx.Next()
	})
	{
		v1.GET("/gets/1", func(ctx contract.Context) {
			ctx.RenderWith(200, render2.Plain, "bdc")
			ctx.Next()
		}).Name("gets.1")
		assertBaseRoute(t, router.(*Router), http.MethodGet, "/v1/gets/1", "gets.1", 2, 1)

		v1.POST("/posts", func(ctx contract.Context) {
			ctx.Next()
		}).Name("v1.posts")
		assertBaseRoute(t, router.(*Router), http.MethodPost, "/v1/posts", "v1.posts", 2, 1)

		//
		v1.DELETE("/delete", func(ctx contract.Context) {
		})
		assertBaseRoute(t, router.(*Router), http.MethodDelete, "/v1/delete", "", 2, 1)

		v1.PUT("/put", func(ctx contract.Context) {
		})
		assertBaseRoute(t, router.(*Router), http.MethodPut, "/v1/put", "", 2, 1)

		v1.PATCH("/patch", func(ctx contract.Context) {
		})
		assertBaseRoute(t, router.(*Router), http.MethodPatch, "/v1/patch", "", 2, 1)

		v1.OPTIONS("/options", func(ctx contract.Context) {
		})
		assertBaseRoute(t, router.(*Router), http.MethodOptions, "/v1/options", "", 2, 1)
	}

	v1Dep := v1.Group("/dep").Before(func(ctx contract.Context) {
		ctx.RenderWith(200, render2.Plain, "Group v1--dep before")
		ctx.Next()
	})
	{
		v1Dep.GET("/gets/1", func(ctx contract.Context) {

		})
	}
	assertBaseRoute(t, router.(*Router), http.MethodGet, "/v1/dep/gets/1", "", 3, 1)
}
