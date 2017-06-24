package main
import (
  "net/http"
  "github.com/julienschmidt/httprouter"
  "log"
  "io/ioutil"
  "os"
  "fmt"
)

func upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  if r.Method == "POST" {
    file, _ := os.Create("/tmp/" + ps.ByName("filename"))
    defer file.Close()
    io.Copy(file, r.Body)
  }
}

func download(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  if r.Method == "GET" {
    http.ServeFile(w, r, "/tmp/" + ps.ByName("filename"))
  }
}

func list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  if r.Method == "GET" {
    io.ioutil.Readdir("/tmp/")
  }
}

func main() {
  router := httprouter.New()
  router.GET("/samples/:filename", download)
  router.POST("/samples/:filename", upload)
  router.GET("/samples", list)
  log.Fatal(http.ListenAndServe(":8888", router))
  out, _ := ioutil.Readdir("/tmp/")
  fmt.Println(out)
}
