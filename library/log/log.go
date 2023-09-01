package log

import (
	"context"

	rules2 "oceanlearn.teach/ginessential/library/log/rules"
	"os"

	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

var logLevelMap = map[string]string{
	"DEBUG": "DEBUG",
	"INFO":  "INFO",
	"WARN":  "WARN",
	"ERROR": "ERROR",
}

const (
	LOG_LEVEL_DEBUG = 1
	LOG_LEVEL_INFO  = 2
	LOG_LEVEL_WARN  = 3
	LOG_LEVEL_ERROR = 4
	LOG_LEVEL_PANIC = 5
	LOG_LEVEL_FATAL = 6
)

var logFuncMap = map[int]func(*zap.Logger, string, ...zap.Field){
	LOG_LEVEL_DEBUG: (*zap.Logger).Debug,
	LOG_LEVEL_INFO:  (*zap.Logger).Info,
	LOG_LEVEL_WARN:  (*zap.Logger).Warn,
	LOG_LEVEL_ERROR: (*zap.Logger).Error,
	LOG_LEVEL_PANIC: (*zap.Logger).Panic,
	LOG_LEVEL_FATAL: (*zap.Logger).Fatal,
}

const (
	ENCODING_JSON    = "json"
	ENCODING_CONSOLE = "console"
)

func getLogLevel() string {
	result := logLevelMap[GetLogLevel()]
	if result == "" {
		if IsDevelop() {
			return "DEBUG"
		} else {
			return "INFO"
		}
	}
	return result
}

func getEncoding() string {
	result := GetLogEncoding()
	if result == "" {
		if IsDevelop() || IsTest() || IsStage() {
			return ENCODING_CONSOLE
		} else {
			return ENCODING_JSON
		}
	}
	return result
}

//func InitLogger() {
//	lumberConfig := &lumberjackConfig{}
//
//	env := os.Getenv("env")
//
//	if env == "" {
//		err := os.Setenv("env", "dev")
//		if err != nil {
//			return
//		}
//		env = "dev"
//	}
//	pathInfo := "conf/" + env + "/log.ini"
//	err := loadIni(pathInfo, "log", lumberConfig)
//	if err != nil {
//		Error("配置文件不存在")
//		return
//	}
//
//	logger = NewLogger(getLogLevel(), nil, lumberConfig)
//	logger = logger.WithOptions(zap.AddCallerSkip(2))
//	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.LstdFlags)
//}

type lumberjackConfig struct {
	FileName   string `json:"fileName"` // fileName 日志写入的文件 Path，空则表示不写入文件
	MaxSize    int    `json:"maxSize"`
	MaxBackUps int    `json:"maxBackUps"`
	MaxAge     int    `json:"maxAge"`
}

// NewLogger
// level 强制指定日志基本 level，空则表示从环境变量中读取
func NewLogger(level string, encoderConfig *zapcore.EncoderConfig, lumberjackConfig *lumberjackConfig) *zap.Logger {
	// 日志写入文件配置
	var hook *lumberjack.Logger
	if lumberjackConfig != nil {
		hook = &lumberjack.Logger{
			Filename:   lumberjackConfig.FileName,
			MaxSize:    lumberjackConfig.MaxSize,
			MaxBackups: lumberjackConfig.MaxBackUps,
			MaxAge:     lumberjackConfig.MaxAge,
			LocalTime:  true,
			Compress:   true,
		}
	}

	zapConfig := &zap.Config{
		EncoderConfig:    zapcore.EncoderConfig{},
		OutputPaths:      nil,
		ErrorOutputPaths: nil,
	}

	var encoder zapcore.Encoder

	zapConfig.EncoderConfig = zap.NewProductionEncoderConfig()
	// 外部配置 zap encoder
	if encoderConfig != nil {
		zapConfig.EncoderConfig = *encoderConfig
	}

	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	switch getEncoding() {
	case ENCODING_CONSOLE:
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
	default:
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
	}

	var writeSyncer zapcore.WriteSyncer
	if hook != nil {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)) // 打印到控制台和文件
	} else {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)) // 打印到控制台
	}

	// 兼容大小写
	level = strings.ToLower(level)

	core := zapcore.NewCore(
		encoder, // 编码器配置
		writeSyncer,
		zap.NewAtomicLevelAt(func() zapcore.Level {
			if level, ok := map[string]zapcore.Level{
				"debug":  zapcore.DebugLevel,
				"info":   zapcore.InfoLevel,
				"warn":   zapcore.WarnLevel,
				"error":  zapcore.ErrorLevel,
				"dpanic": zapcore.DPanicLevel,
				"panic":  zapcore.PanicLevel,
				"fatal":  zapcore.FatalLevel,
			}[level]; ok {
				return level
			} else {
				return zapcore.ErrorLevel
			}
		}()),
	)
	//return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).With(zap.String("app", app))
	return zap.New(core, zap.AddCaller())
}

func GetZapLog() *zap.SugaredLogger {
	return logger.Sugar()
}

func InfoWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	fields = extractCtxInfo(ctx, fields...)
	doLog(LOG_LEVEL_INFO, msg, fields...)
}

func WarnWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	fields = extractCtxInfo(ctx, fields...)
	doLog(LOG_LEVEL_WARN, msg, fields...)
}

func DebugWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	fields = extractCtxInfo(ctx, fields...)
	doLog(LOG_LEVEL_DEBUG, msg, fields...)
}

func ErrorWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	fields = extractCtxInfo(ctx, fields...)
	doLog(LOG_LEVEL_ERROR, msg, fields...)
}

func FatalWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	fields = extractCtxInfo(ctx, fields...)
	doLog(LOG_LEVEL_FATAL, msg, fields...)
}

func extractCtxInfo(ctx context.Context, fields ...zap.Field) []zap.Field {
	if userId := ctx.Value("user_id"); userId != nil {
		fields = append(fields, zap.Any("user_id", userId))
	}

	traceId := GetTraceId(ctx)
	fields = append(fields, zap.String(traceIdKey, traceId))
	return fields
}

func Info(msg string, fields ...zap.Field) {
	doLog(LOG_LEVEL_INFO, msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	doLog(LOG_LEVEL_WARN, msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	doLog(LOG_LEVEL_ERROR, msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	doLog(LOG_LEVEL_DEBUG, msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	doLog(LOG_LEVEL_PANIC, msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	doLog(LOG_LEVEL_FATAL, msg, fields...)
}

func doLog(level int, msg string, fields ...zap.Field) {
	// fieldList 储存过滤完毕的值,不影响field...里的值
	//fieldList := make([]zap.Field, 0)
	//for _, one := range fields {
	//	// reflect 需要传入ptr，否则不进行过滤
	//	if  one.Type == zapcore.ReflectType &&
	//		one.Interface != nil &&
	//		reflect.TypeOf(one.Interface).Kind() != reflect.Ptr {
	//
	//		fieldList = append(fieldList, one)
	//	} else {
	//		fieldList = append(fieldList, *filter(&one))
	//	}
	//}

	// 传入过滤完毕的值去实际的log函数记录日志
	//logFuncMap[level](logger, msg, fieldList...)
	logFuncMap[level](logger, msg, fields...)
}

func filter(field *zap.Field) *zap.Field {
	return rules2.Filter(field)
}

//func init() {
//	InitLogger()
//}

// Setup Initialize the utils
//func getCurrentAbPathByCaller() string {
//	var abPath string
//	_, filename, _, ok := runtime.Caller(0)
//	if ok {
//		abPath = path.Dir(filename)
//		index := strings.Index(abPath, "whgo_xjd/")
//		if index != -1 {
//			return abPath[0 : index+len("whgo_xjd/")]
//		} else {
//			Fatal("项目文件夹必须使用whgo_xjd命名", zap.String("path", abPath))
//		}
//		return abPath
//	}
//	return ""
//}
//
//func loadIni(path, section string, v interface{}) (err error) {
//	abPath := getCurrentAbPathByCaller()
//	cfg, err := ini.Load(abPath + path)
//	if err != nil {
//		Fatal("setting.Setup, fail to parse", zap.String("path", abPath+path), zap.Error(err))
//		return
//	}
//
//	err = cfg.Section(section).MapTo(v)
//	if err != nil {
//		Fatal("Cfg.MapTo err", zap.String("section", section), zap.Error(err))
//		return
//	}
//
//	return
//}
