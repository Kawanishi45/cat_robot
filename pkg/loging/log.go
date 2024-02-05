package loging

import (
  "fmt"
  "github.com/Kawanishi45/cat_robot/pkg/constants"
  "log"
  "os"
  "time"
)

func CreateLog() (fileName string, err error) {
  // ai_2024_0205_1455.log の形式で.logファイルを作成する
  currentTime := time.Now()
  fileName = fmt.Sprintf("ai_%d_%02d%02d_%02d%02d.log",
    currentTime.Year(), currentTime.Month(), currentTime.Day(),
    currentTime.Hour(), currentTime.Minute())
  var file *os.File
  file, err = os.Create(fileName)
  if err != nil {
    return
  }
  defer file.Close()
  return
}

func SaveLog(logText string) {
  // txtファイルに記録していく
  f, err := os.OpenFile(constants.AiLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    log.Println("Failed to open log file:", err)
    return
  }
  defer func(f *os.File) {
    err = f.Close()
    if err != nil {
      log.Println("Failed to close log file:", err)
    }
  }(f)
  if _, err = f.WriteString(logText + "\n"); err != nil {
    log.Println("Failed to write log file:", err)
  }
}
