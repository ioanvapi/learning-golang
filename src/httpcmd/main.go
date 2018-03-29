package main

import (
	"io"
	"net/http"
	"os/exec"

	"github.com/Akagi201/light"
	log "github.com/sirupsen/logrus"
)

func main() {
	root := light.New()
	root.Post("/", func(w http.ResponseWriter, r *http.Request) {
    log.Info("receive push")
		args := []string{"pull"}
		cmd := exec.Command("git", args...)
		err := cmd.Run()
		if err != nil {
			log.Errorf("git pull failed, err: %v", err)
		}
		io.WriteString(w, "OK")
	})

	http.ListenAndServe(":2345", root)
}
