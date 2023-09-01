package log

import "os"

const (
	ENV_DEVELOP    = "develop"
	ENV_TEST       = "test"
	ENV_STAGE      = "stage"
	ENV_PRODUCTION = "production"
	ENV_MIRROR     = "mirror"
	ENV_LOCALHOST  = "localhost"
	ENV_UNITTEST   = "unittest"
	ENV_CRON       = "cron"
)

func GetLogLevel() string {
	return GetKeyString("LOG_LEVEL")
}

func GetEnv() string {
	return GetKeyString("ENVIRON")
}

func IsDevelop() bool {
	return GetEnv() == ENV_DEVELOP
}

func IsStage() bool {
	return GetEnv() == ENV_STAGE
}

func IsTest() bool {
	return GetEnv() == ENV_TEST
}

func GetKeyString(key string) string {
	return os.Getenv(key)
}

func GetLogEncoding() string {
	return GetKeyString("LOG_ENCODING")
}
