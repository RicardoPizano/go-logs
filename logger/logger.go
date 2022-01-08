package logger

import (
	"os"
	"strconv"
	"time"
)

var (
	log      Logger
	hostname string
)

type Fields map[string]interface{}

const (
	//Debug has verbose message
	debugLevel = "debug"
	//Info is default log level
	infoLevel = "info"
	//Warn is for logging messages about possible issues
	warnLevel = "warn"
	//Error is for logging errors
	errorLevel = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	fatalLevel = "fatal"
)

//const (
//	pathName             = "logs/api-masterAccounts_%s"
//	logLevelEnv          = "LOG_LEVEL"           // Default 1 (info level)
//	enableFileLogsEnv    = "ENABLE_FILE_LOG"     // Default false
//	enableConsoleLogsEnv = "ENABLE_CONSOLE_LOGS" // Default true
//)

//Logger is our contract for the logger
type Logger interface {
	Debug(moduleName string, functionName string, text string)

	Info(moduleName string, functionName string, text string)

	Warn(moduleName string, functionName string, text string)

	Error(moduleName string, functionName string, text string)

	Fatal(moduleName string, functionName string, text string)

	Request(method string, statusCode int, uri string, start time.Time)
}

// Configuration stores the config for the logger
type Configuration struct {
	EnableConsole bool
	ConsoleLevel  string
	EnableFile    bool
	FileLevel     string
	FileLocation  string
}

//func init() {
//	fileLocation := fmt.Sprintf(pathName, getHostname())
//	logLevel := getLogLevel(os.Getenv(logLevelEnv))
//	enableFileLogs := getEnvVar(enableFileLogsEnv, false)
//	enableConsoleLogs := getEnvVar(enableConsoleLogsEnv, true)
//	config := Configuration{
//		EnableConsole: enableConsoleLogs,
//		ConsoleLevel:  logLevel,
//		EnableFile:    enableFileLogs,
//		FileLevel:     logLevel,
//		FileLocation:  fileLocation,
//	}
//	logger, err := newZapLogger(config)
//	if err != nil {
//		print(err.Error())
//	}
//	log = logger
//}

func CreateLogger(config Configuration) {
	logger, err := newZapLogger(config)
	if err != nil {
		print(err.Error())
	}
	log = logger
}

func GetEnvVar(env string, defaultEnv bool) bool {
	envVar, err := strconv.ParseBool(os.Getenv(env))
	if err != nil {
		return defaultEnv
	}
	return envVar
}

func GetLogLevel(logLevel string) string {
	switch logLevel {
	case "0":
		return debugLevel
	case "1":
		return infoLevel
	case "2":
		return warnLevel
	case "3":
		return errorLevel
	default:
		return infoLevel
	}
}

func GetHostname() string {
	if hostname == "" {
		hostname, _ := os.Hostname()
		return hostname
	} else {
		return hostname
	}
}

func Debug(moduleName string, functionName string, text string) {
	log.Debug(moduleName, functionName, text)
}

func Info(moduleName string, functionName string, text string) {
	log.Info(moduleName, functionName, text)
}

func Warn(moduleName string, functionName string, text string) {
	log.Warn(moduleName, functionName, text)
}

func Error(moduleName string, functionName string, text string) {
	log.Error(moduleName, functionName, text)
}

func Fatal(moduleName string, functionName string, text string) {
	log.Fatal(moduleName, functionName, text)
}

func Request(method string, statusCode int, uri string, start time.Time) {
	log.Request(method, statusCode, uri, start)
}
