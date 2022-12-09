package middleware

import (
	"context"
	"dating/platform/logger"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GinLogger(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		id := uuid.NewV4().String()
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "x-request-id", id))
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "request-start-time", start))
		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)
		fields := []zapcore.Field{
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.Int64("request-latency", latency.Milliseconds()),
		}
		fields = append(fields, zap.String("time", end.Format(time.RFC3339)))
		log.Info(ctx.Request.Context(), "GIN", fields...)
	}
}
