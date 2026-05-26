package util

import (
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// a short time format; like time.Kitchen but with 24-hour notation.
	Kitchen24 = "15:04"
	// a time format that just cares about the day and month.
	YearDay = "Jan_2"
)

// TimePeriod represents a time period with a single beginning and end.
type TimePeriod struct {
	From time.Time
	To   time.Time
}

// NewTimePeriod returns a normalized TimePeriod given a start and end time.
func NewTimePeriod(from, to time.Time) TimePeriod {
	_ = "STUB: not implemented"
	return *new(TimePeriod)
}

// Includes returns true iff the given pointInTime's time of day is included in time period tp.
func (tp TimePeriod) Includes(pointInTime time.Time) bool { _ = "STUB: not implemented"; return false }

// String returns tp as a pretty string.
func (tp TimePeriod) String() string { _ = "STUB: not implemented"; return "" }

// ParseWeekdays takes a comma-separated list of abbreviated weekdays (e.g. sat,sun) and turns them
// into a slice of time.Weekday. It ignores any whitespace and any invalid weekdays.
func ParseWeekdays(weekdays string) []time.Weekday { _ = "STUB: not implemented"; return nil }

// ParseTimePeriods takes a comma-separated list of time periods in Kitchen24 format and turns them
// into a slice of TimePeriods. It ignores any whitespace.
func ParseTimePeriods(timePeriods string) ([]TimePeriod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseDays(days string) ([]time.Time, error) { _ = "STUB: not implemented"; return nil, nil }

// TimeOfDay normalizes the given point in time by returning a time object that represents the same
// time of day of the given time but on the very first day (day 0).
func TimeOfDay(pointInTime time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// FormatDays takes a slice of times and returns a slice of strings in YearDate format (e.g. [Apr 1, Sep 24])
func FormatDays(days []time.Time) []string { _ = "STUB: not implemented"; return nil }

// NewPod returns a new pod instance for testing purposes.
func NewPod(namespace, name string, phase v1.PodPhase) v1.Pod {
	_ = "STUB: not implemented"
	return *new(v1.Pod)
}

// NewPodWithOwner returns a new pod instance for testing purposes with a given owner UID
func NewPodWithOwner(namespace, name string, phase v1.PodPhase, owner types.UID) v1.Pod {
	_ = "STUB: not implemented"
	return *new(v1.Pod)
}

// NewNamespace returns a new namespace instance for testing purposes.
func NewNamespace(name string) v1.Namespace { _ = "STUB: not implemented"; return *new(v1.Namespace) }

// RandomPodSubSlice creates a shuffled subslice of the give pods slice
func RandomPodSubSlice(pods []v1.Pod, count int) []v1.Pod { _ = "STUB: not implemented"; return nil }
