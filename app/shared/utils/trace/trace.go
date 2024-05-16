package trace

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"fmt"
	"strings"
	"ml-elizabeth/app/domain/entity"
	"ml-elizabeth/app/shared/utils/constants"
	log "github.com/sirupsen/logrus"
)

func New(spanID, traceparent string) *entity.Trace {
	var (
		traceID        string
		parentID       string
		traceparentExp = regexp.MustCompile(constants.TraceparentRegexp)
	)

	if traceparentExp.MatchString(traceparent) {
		splittedTrace := strings.Split(traceparent, "-")
		traceID = splittedTrace[1]
		parentID = splittedTrace[2]

		if !isValidTraceID(splittedTrace[1]) {
			traceID = getDefaultTraceID(spanID)
		}

		if !isValidParentID(splittedTrace[2]) {
			parentID = spanID
		}


		return &entity.Trace{
			SpanID:    spanID,
			TraceID:   traceID,
			ParentID:  parentID,
	
		}
	}

	
	return &entity.Trace{
		SpanID:    spanID,
		TraceID:   getDefaultTraceID(spanID),
		ParentID:  parentID,

	}
}

func NewSpanID(lenSpanID int) string {
	bytes := make([]byte, lenSpanID/constants.HalfSpanID)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Error("could not create SpanID")
		return "0000000000000000"
	}
	return hex.EncodeToString(bytes)
}


func isValidTraceID(traceID string) bool {
	return traceID != "" && traceID != "00000000000000000000000000000000"
}


func isValidParentID(parentID string) bool {
	return parentID != "" && parentID != "0000000000000000"
}

func getDefaultTraceID(spanID string) string {
	return fmt.Sprintf("0000000000000000%s", spanID)
}
