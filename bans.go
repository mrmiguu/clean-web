package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("")))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	r.Header.Add("Content-Type", "text/html; charset=utf-8")
	// 	http.ServeFile(w, r, "index.html")
	// })

	http.ListenAndServe("localhost:80", nil)

	// dxweb.Width = js.Global.Get("window").Get("innerWidth").Int()
	// dxweb.Height = js.Global.Get("window").Get("innerHeight").Int()

	// bg := <-dxweb.LoadImage("assets/bg.png")
	// // bg.Resize(dxweb.Width, dxweb.Height)

	// logo := bg.NewText(`🚢`)
	// logo.Resize(204)
	// logo.Show(true)

	// bg.Show(true, 2500)
}
