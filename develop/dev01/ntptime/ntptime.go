package ntptime

import (
	"time"

	"github.com/beevik/ntp"
)

type NtpTime struct {
	ntpServer string
	timeout   time.Duration
}

func (t NtpTime) CurrentTime() (time.Time, error) {
	return ntp.Time(t.ntpServer)
}

func NewNtpTime(ntpServer string, timeout time.Duration) *NtpTime {
	return &NtpTime{ntpServer, timeout}
}
