package trunk

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/malaow3/trunk/formatter"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger //nolint

func InitializeLogger() {
	var err error
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func InitLogger() {
	log.SetFormatter(&formatter.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05.000",
		CallerFirst:     true,
		ColorTimestamp:  true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
		},
	})
	log.SetReportCaller(true)
}

func InitLoggerFullpath() {
	log.SetFormatter(&formatter.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05.000",
		CallerFirst:     true,
		ColorTimestamp:  true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d][%s()]", f.File, f.Line, funcName)
		},
	})
	log.SetReportCaller(true)
}

func InitLoggerNoColors() {
	log.SetFormatter(&formatter.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05.000",
		CallerFirst:     true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
		},
		NoColors: true,
	})
	log.SetReportCaller(true)
}

func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
