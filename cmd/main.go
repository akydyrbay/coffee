package main

import (
	"flag"
	"fmt"
	. "frappuccino/config"
	. "frappuccino/internal/handler"
	"net/http"
)

func main() {
	flag.IntVar(&PortN, "port", 8080, "Port number")
	flag.StringVar(&Path, "dir", "data", "Path to the directory")
	flag.Usage = func() {
		fmt.Print("Coffee Shop Management System\n\n",
			"Usage:\n",
			"\thot-coffee [-port <N>] [-dir <S>]\n\thot-coffee --help\n\n",
			"Options:\n",
			"- --help\tShow this screen.\n",
			"- --port N\tPort number\n",
			"- --dir S\tPath to the directory\n")
	}
	flag.Parse()
	fmt.Printf("Server Started\nPort: %d\tDirectory: %s\n", PortN, Path)
	Path = Path + "/"
	mux := http.NewServeMux()
	mux.Handle("/", &RequestHandler{})
	port := ":" + fmt.Sprint(PortN)

	http.ListenAndServe(port, mux)
}
