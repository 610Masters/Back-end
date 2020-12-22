/*
 * Swagger Blog
 *
 * Simple Blog
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}

func init() {
	apiMethodbind("Index", "GET", "/", Index)
	apiMethodbind("ArticleIdCommentsGet", "GET", "/article/{id}/comments", ArticleIdCommentsGet)
	apiMethodbind("ArticleIdGet", "GET", "/article/{id}", ArticleIdGet)
	apiMethodbind("ArticlesGet", "GET", "/articles", ArticlesGet)
	apiMethodbind("ArticleIdCommentPost", "POST", "/article/{id}/comment", ArticleIdCommentPost)
	apiMethodbind("Options", "OPTIONS", "/article/{id}/comment", Options)
	apiMethodbind("UserLoginPost", "POST", "/user/login", UserLoginPost)
	apiMethodbind("UserRegisterPost", "POST", "/user/register", UserRegisterPost)
}

func apiMethodbind(name string, method string, pattern string, handlerFunc http.HandlerFunc) {
	Routes = append(Routes, Route{name, method, pattern, handlerFunc})
}

func Index(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("go/index.html")
	if err != nil {
		log.Fatal(err)
	}
	index, _ := ioutil.ReadAll(f)
	fmt.Fprintf(w, string(index))
}
