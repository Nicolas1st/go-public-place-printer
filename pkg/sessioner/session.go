package sessioner

import "time"

type SessionToken string
type UID uint

type Session struct {
	UID        UID
	ExpiryTime time.Time
}
