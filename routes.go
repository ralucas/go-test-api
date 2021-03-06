package main

import (
  "net/http"

  "github.com/gorilla/mux"
)

type Route struct {
  Name          string
  Method        string
  Pattern       string
  HandlerFunc   http.HandlerFun
}

type Routes []Route

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "TodoIndex",
    "GET",
    "/todos",
    TodoIndex
  },
  Route{
    "TodoShow",
    "GET",
    "/todos/{todoId}",
    TodoShow,
  },
}

