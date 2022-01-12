package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/Amovement/password-box/pkg/utils/loghook"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MyFormatter struct {
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	dayFormat := entry.Time.Format("2006-01-02")
	timeFormat := entry.Time.Format("15:04:05")
	var newLog string
	newLog = fmt.Sprintf("[%s]  [%s] |[%s]| [%s] :%s\n", dayFormat, timeFormat, entry.Level, entry.Data["line"], entry.Message)
	b.WriteString(newLog)
	return b.Bytes(), nil
}

var timeFormat = "02/Jan/2006:15:04:05 -0700"

var log = logrus.New()

func init() {
	log.SetOutput(os.Stdout)
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{
		file,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetFormatter(&MyFormatter{})
	log.AddHook(loghook.NewContextHook())
	if err == nil {
		log.SetOutput(fileAndStdoutWriter)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func GetLog() *logrus.Logger {
	//logrus.SetFormatter(&MyFormatter{})
	return log
}

// Logger is the logrus logger handler
func Logger(logger logrus.FieldLogger, notLogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		// param := c.Request.URL.
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		header := c.Request.Header
		raw_url := c.Request.URL
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
			"header":     header,
			"rawUrl":     raw_url,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)\n", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
