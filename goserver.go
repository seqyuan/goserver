package main

import (
     "flag"
     "fmt"
     "net/http"
     "os"
     "log"
     "path/filepath"
     "strconv"
)

func main() {
     flag.Parse()
     server_name := filepath.Base(os.Args[0])
     cwd, _ := os.Getwd()

     port := 8090
     absdir, _ := filepath.Abs(cwd)

     switch len(flag.Args()) {
     case 0:
         fmt.Println(fmt.Sprintf("%s <port> <path(default pwd)>\n%s 8090 ./\n%s 8090\n",server_name,server_name,server_name))
     case 1:
         port, _ = strconv.Atoi(os.Args[1])
     case 2:
         port, _ = strconv.Atoi(os.Args[1])
         path, _ := filepath.Abs(os.Args[2])
         absdir = path
     }

     fmt.Println(fmt.Sprintf("%s %v %s", server_name, port, absdir))
     err := http.ListenAndServe(fmt.Sprintf(":%v", port), http.FileServer(http.Dir(absdir)))
     if err != nil{
         log.Fatal(err)
     }
}

