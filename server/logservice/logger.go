package logservice

import (
	"fmt"

	"github.com/1x-eng/struct_laas/gen/logger"
)

func ProcessLog(logMsg *logger.LogMessage) {
	fmt.Printf("Received log: Level=%s, Message=%s, Service=%s, TraceID=%s, Timestamp=%s, AdditionalData=%v\n",
		logMsg.Level, logMsg.Message, logMsg.Service, logMsg.TraceId, logMsg.Timestamp, logMsg.AdditionalData)
}
