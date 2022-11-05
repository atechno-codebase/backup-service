package main

import (
	"ates/services/backup/handlers"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/the-kaustubh/dynpkg/logger"
	"github.com/tidwall/gjson"
)

var PORT string

func init() {
	content, err := ioutil.ReadFile("backupConfig.json")
	if err != nil {
		panic(err)
	}
	svcConf := gjson.ParseBytes(content)
	logger.InitDefaultLogger("backup-service", svcConf)
	PORT = svcConf.Get("port").String()
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
