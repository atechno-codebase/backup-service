package backupsvc

import (
	"ates/services/backup/config"
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var backupDir string

func Init() {
	bDir := config.Configutaion.BackupDir
	logrus.Info("backupDir from config: ", bDir)
	if len(bDir) == 0 {
		bDir = "backup"
	}
	backupDir = bDir
}

func GetBackupDir() string {
	return backupDir
}

func CreateBackup() {
	uname, pwd := config.Configutaion.DbUserName, config.Configutaion.DbPassword
	dt := time.Now().Format("20060102150405")
	fileName, err := filepath.Abs(fmt.Sprintf("%s/%s-%s.gz", backupDir, "backup", dt))
	if err != nil {
		logrus.Error(err)
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
		logrus.Error(err)

	}
	logrus.Error("stderr: ", oErr.String())
	logrus.Error("stdout: ", oStd.String())
}
