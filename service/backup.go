package backupsvc

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/the-kaustubh/dynpkg/config"
	"github.com/the-kaustubh/dynpkg/logger"
)

var backupDir string

func init() {
	err := config.LoadConfig(".", "backupConfig.json", "")
	if err != nil {
		panic(err)
	}
	bDir := config.Get("backupDir").String()
	if len(bDir) == 0 {
		bDir = "backup"
	}
	backupDir = bDir
}

func GetBackupDir() string {
	return backupDir
}

func CreateBackup() {
	dt := time.Now().Format("20060102150405")
	fileName, err := filepath.Abs(fmt.Sprintf("%s/%s-%s.gz", backupDir, "backup", dt))
	if err != nil {
		logger.LogError(err)
		return
	}

	var oErr, oStd bytes.Buffer
	bkpCmd := exec.Command("mongodump", "--db=niv", fmt.Sprintf("--archive=%s", fileName), "--gzip")
	bkpCmd.Stderr = &oErr
	bkpCmd.Stdout = &oStd
	bkpCmd.Dir = backupDir
	err = bkpCmd.Run()
	if err != nil {
		logger.LogError(err)

	}
	logger.LogError("stderr: ", oErr.String())
	logger.LogError("stdout: ", oStd.String())
}
