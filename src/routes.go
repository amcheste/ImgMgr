package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route {
        "HealthCheck",
        "GET",
        "/health",
        Health,
    },
    Route {
        "StateCheck",
        "GET",
        "/state",
        State,
    },
    Route{
        "ListAlbums",
        "GET",
        "/albums",
        ListAlbums,
    },
}
