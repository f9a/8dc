---
title: Zap Snippets
date: 2019-04-01T00:04:02+01:00
categories:
  - go
url: /post/uuid/185fdbc6-a530-5da5-9bd3-9d4af9f43ae9
---

```go
func makeFileLogger(logger *zap.Logger, importDir string) (log *zap.Logger) {
	logFileName := fmt.Sprintf("errors-%v.log", time.Now().Format(time.RFC3339))
	logFile := filepath.Join(importDir, logFileName)

	fhErrorLogs, err := os.Create(logFile)
	if err != nil {
		logger.Fatal("Cannot create error log file", zap.Error(err))
	}

	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	systemdEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	fileEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	systemdWriter := zapcore.Lock(os.Stdout)
	fileWriter := zapcore.AddSync(fhErrorLogs)

	core := zapcore.NewTee(
		zapcore.NewCore(systemdEncoder, systemdWriter, allPriority),
		zapcore.NewCore(fileEncoder, fileWriter, allPriority),
	)

	return zap.New(core)
}

```
