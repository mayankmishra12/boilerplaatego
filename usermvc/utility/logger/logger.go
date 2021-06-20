
package logger

import (
"context"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"

	"github.com/gofrs/uuid"
"go.uber.org/zap"
)
var sugarLogger *zap.SugaredLogger

type key string

const (
	_callerKey = "caller"
	_traceKey  = "traceID"

	// skipOneCaller denotes the number of frames to be skipped
	// to get the caller two level above in the stack
	_skipOneCaller = 2
)

// GetLoggerWithContext returns a global logger with Proper CallerName and TranceID
func GetLoggerWithContext(ctx context.Context) *zap.SugaredLogger {
	if getTraceID(ctx) == "" {
		ctx = SetTraceID(ctx)
	}
	return zap.S().With(  _traceKey, getTraceID(ctx))
}

// getTraceID returns traceID from the context
func getTraceID(ctx context.Context) string {
	traceID := ctx.Value(key(_traceKey))
	if traceID != nil {
		return traceID.(string)
	}
	return ""
}

// SetTraceID sets trace ID to context for logging purpose
// TraceID is being set on the best effort basis, there can be scenario
// where we are not able to set a traceID and proceed with empty traceID
func SetTraceID(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, key(_traceKey), generateTraceID())
	return ctx
}

// generateTraceID generates and return empty trace ID
func generateTraceID() string {
	traceID, _ := uuid.NewV4()
	return traceID.String()

}

func InitLogger() *zap.Logger{
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	logFile :=  filepath.Join("", "appmanager.log")
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10,
		MaxAge:     10,
		MaxBackups: 10,
		LocalTime:  false,
		Compress:   false,
	})
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig),w, zap.InfoLevel)
	logger = zap.New(core,zap.AddCaller())
	return logger
}

