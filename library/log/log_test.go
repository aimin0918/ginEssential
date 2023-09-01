package log

import (
	"go.uber.org/zap"
	"testing"
)

type ObjectLogger struct {
	Password string
	Pass     string
	Name     string
	Anyfield string
}

func TestInfo(t *testing.T) {
	Info("test",
		zap.Any("pass", &ObjectLogger{
			Password: "2333333",
			Pass:     "plaintext password",
			Name:     "cf",
			Anyfield: "ssss",
		}),
		zap.Any("pass", &map[string]string{
			"password":     "2333333",
			"name":         "cf",
			"hahahahfield": "ssss",
		}),
	)
	Info("test", zap.String("pass", "hdhj1998881919"))
	//Warn("yet another test2", zap.String("ordernum", "hahahahahahahaha"))
}

func TestFile(t *testing.T) {
	logger := NewLogger("DEBUG", nil, &lumberjackConfig{
		FileName: "test.log",
		MaxSize:  100,
	})

	logger.Error("test")
}
