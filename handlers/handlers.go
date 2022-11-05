package handlers

import (
	backupsvc "ates/services/backup/service"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/the-kaustubh/dynpkg/logger"
)

func DeleteBackup(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) <= 2 {
		fmt.Fprintf(w, `{"error": "invalid path"}`)
		return
	}
	fileName := parts[2]
	filePath := fmt.Sprintf("%s/%s", backupsvc.GetBackupDir(), fileName)
	err := os.Remove(filePath)
	if err != nil {
		logger.LogError(err)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}
	fmt.Fprintf(w, `{"message": "deleted %s"}`, fileName)
}

func CreateBackup(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	logger.LogInfo(fmt.Sprintf("backup started: %s", now))
	go backupsvc.CreateBackup()
	fmt.Fprintf(w, `{"error": "backup started: %s"}`, now)
	return
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	entries, err := os.ReadDir(backupsvc.GetBackupDir())
	if err != nil {
		logger.LogError(err)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	backupsList := []string{}
	for _, entry := range entries {
		backupsList = append(backupsList, entry.Name())
	}

	jsonArray, err := json.Marshal(backupsList)
	if err != nil {
		logger.LogError(err)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}
	fmt.Fprintf(w, `{"files": %s}`, string(jsonArray))
}

func DownloadBackup(w http.ResponseWriter, r *http.Request) {
	fileServerPath := backupsvc.GetBackupDir()
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) <= 1 {
		fmt.Fprintf(w, `{"error": "invalid path"}`)
		return
	}
	fileName := parts[1]
	fullPath := fmt.Sprintf("%s/%s", fileServerPath, fileName)
	http.ServeFile(w, r, fullPath)
}
