package main

import (
	"ates/services/backup/config"
	"ates/services/backup/handlers"
	backupsvc "ates/services/backup/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var PORT string

func init() {
	config.Init()
	backupsvc.Init()
	logrus.Info(config.Configutaion)
	PORT = config.Configutaion.Port
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		endpoint := r.URL.String()
		if strings.HasPrefix(endpoint, "/delete") {
			handlers.DeleteBackup(w, r)
			return
		}
		if strings.HasPrefix(endpoint, "/create") {
			handlers.CreateBackup(w, r)
			return
		}
		if endpoint == "/" {
			handlers.ListAll(w, r)
			return
		}
		handlers.DownloadBackup(w, r)
	})
	logrus.Error(fmt.Sprintf(`Backup service started on port: "%s"`, PORT))
	logrus.Error(http.ListenAndServe(PORT, nil))
}
