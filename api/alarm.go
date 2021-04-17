package api

import "time"

// AlarmMode represents the mode for the alarm.
type AlarmMode int

// AlarmMode enums.
const (
	AlarmCreate  AlarmMode = 0
	AlarmDelete  AlarmMode = 2
	AlarmEnable  AlarmMode = 3
	AlarmDisable AlarmMode = 4
)

// Alarm for speaker.
type Alarm struct {
	ID         int    `json:"id"`        // When creating alarm, -1.
	Day        int    `json:"day"`       // Days when it is active, e.g. Monday + Tuesday = 0b1100000.
	DayRepeat  bool   `json:"dayrepeat"` // Part of request, significance?
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Type       int    `json:"type"`     // Always 0?
	Duration   int    `json:"duration"` // Duration in minutes.
	Volume     int    `json:"volume"`
	Enable     bool   `json:"enable"`
	Shuffle    bool   `json:"shuffle"`
	Title      string `json:"title"`      // "Default alarm sound"
	ModifiedID int    `json:"modifiedid"` // Part of request, significance?
	M2ID       string `json:"m2id"`       // Part of response, significance?
	IPAddr     string `json:"ipaddress"`  // Song server addr? Samba?
	SongPath   string `json:"songpath"`   // Samba path?

	Mode AlarmMode `json:"mode,omitempty"`
}

// AlarmDays enables one or more day(s) for an alarm.
func AlarmDays(days ...time.Weekday) int {
	var ad int
	for _, wd := range days {
		d := FromWeekday(wd)
		ad |= 1 << (6 - d)
	}
	return ad
}
