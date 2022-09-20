package ntptime

import (
	"github.com/beevik/ntp"
	"time"
)

type TimeNtp struct {
	ntpServer string
	timeout   time.Duration
}

func (t TimeNtp) CurrentTime() (time.Time, error) {
	return ntp.Time(t.ntpServer)
}

func NewTimeLib(ntpServer string, timeout time.Duration) *TimeNtp {
	return &TimeNtp{ntpServer, timeout}
}
