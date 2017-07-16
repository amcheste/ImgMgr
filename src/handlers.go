package main

import (
    "encoding/json"
    "net/http"
    "log"
    "os"
    "io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request){

    paths := []string{
        "/",
        "/health",
        "/state",
        "/albums",
    }

    json.NewEncoder(w).Encode(paths)
}

func Health(w http.ResponseWriter, r *http.Request){
    //Health only retuns 200 if the service is reachable
}


func State(w http.ResponseWriter, r *http.Request){
    //State only retuns 200 if the service is reachable
}

func ListAlbums(w http.ResponseWriter, r *http.Request){
    log.Println("Listing Albums")

    albumPath := os.Getenv("CAM_ALBUM_PATH")
    if albumPath == "" {
        log.Println("CAM_ALBUM_PATH not set using default...")
        albumPath = "/Media/Pictures"
    }

    files, _ := ioutil.ReadDir(albumPath)
    var albums []Album
    for _, f := range files {
        album := Album{Name: f.Name()}
        albums = append(albums, album)
    }

    json.NewEncoder(w).Encode(albums)
}

