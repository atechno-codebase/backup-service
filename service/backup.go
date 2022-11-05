package backupsvc

import (
	"ates/services/backup/config"
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/the-kaustubh/dynpkg/logger"
)

var backupDir string

func Init() {
	bDir := config.Get("backupDir").String()
	fmt.Println("backupDir from config: ", bDir)
	if len(bDir) == 0 {
		bDir = "backup"
	}
	backupDir = bDir
}

func GetBackupDir() string {
	return backupDir
}

func CreateBackup() {
	uname, pwd := config.GetDBCreds()
	dt := time.Now().Format("20060102150405")
	fileName, err := filepath.Abs(fmt.Sprintf("%s/%s-%s.gz", backupDir, "backup", dt))
	if err != nil {
		logger.LogError(err)
		return
	}

	var oErr, oStd bytes.Buffer
	params := []string{
		"--db=niv",
		fmt.Sprintf("--archive=%s", fileName),
		"--gzip",
		"--username",
		uname,
		"--password",
		pwd,
		"--authenticationDatabase",
		"admin",
	}
	bkpCmd := exec.Command("mongodump", params...)

	fmt.Println("trying cmd: ", bkpCmd.String())
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
