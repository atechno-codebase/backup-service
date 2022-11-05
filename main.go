package main

import (
	"ates/services/backup/config"
	"ates/services/backup/handlers"
	backupsvc "ates/services/backup/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/the-kaustubh/dynpkg/logger"
)

var PORT string

func init() {
	config.Init()
	backupsvc.Init()
	fmt.Println(config.Get("@this"))
	logger.InitDefaultLogger("backup-service", config.Get("@this"))
	PORT = config.Get("port").String()
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
	logger.LogError(fmt.Sprintf(`Backup service started on port: "%s"`, PORT))
	logger.LogFatal(http.ListenAndServe(PORT, nil))
}
