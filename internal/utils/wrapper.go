package utils

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func WrapperTime(t *time.Time) *timestamppb.Timestamp {
	if t != nil {
		tmp := timestamppb.New(*t)
		return tmp
	}
	return nil
}
