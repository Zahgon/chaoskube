package chaoskube

import (
	"context"
	"errors"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"

	"github.com/linki/chaoskube/notifier"
	"github.com/linki/chaoskube/terminator"
	"github.com/linki/chaoskube/util"
)

// Chaoskube represents an instance of chaoskube
type Chaoskube struct {
	// a kubernetes client object
	Client kubernetes.Interface
	// a label selector which restricts the pods to choose from
	Labels labels.Selector
	// an annotation selector which restricts the pods to choose from
	Annotations labels.Selector
	// a kind label selector which restricts the kinds to choose from
	Kinds labels.Selector
	// a namespace selector which restricts the pods to choose from
	Namespaces labels.Selector
	// a namespace label selector which restricts the namespaces to choose from
	NamespaceLabels labels.Selector
	// a regular expression for pod names to include
	IncludedPodNames *regexp.Regexp
	// a regular expression for pod names to exclude
	ExcludedPodNames *regexp.Regexp
	// a list of weekdays when termination is suspended
	ExcludedWeekdays []time.Weekday
	// a list of time periods of a day when termination is suspended
	ExcludedTimesOfDay []util.TimePeriod
	// a list of days of a year when termination is suspended
	ExcludedDaysOfYear []time.Time
	// the timezone to apply when detecting the current weekday
	Timezone *time.Location
	// minimum age of pods to consider
	MinimumAge time.Duration
	// an instance of logrus.StdLogger to write log messages to
	Logger log.FieldLogger
	// a terminator that terminates victim pods
	Terminator terminator.Terminator
	// dry run will not allow any pod terminations
	DryRun bool
	// grace period to terminate the pods
	GracePeriod time.Duration
	// event recorder allows to publish events to Kubernetes
	EventRecorder record.EventRecorder
	// a function to retrieve the current time
	Now func() time.Time

	MaxKill int
	// chaos events notifier
	Notifier notifier.Notifier
	// namespace scope for the Kubernetes client
	ClientNamespaceScope string
}

var (
	// errPodNotFound is returned when no victim could be found
	errPodNotFound = errors.New("pod not found")
	// msgVictimNotFound is the log message when no victim was found
	msgVictimNotFound = "no victim found"
	// msgWeekdayExcluded is the log message when termination is suspended due to the weekday filter
	msgWeekdayExcluded = "weekday excluded"
	// msgTimeOfDayExcluded is the log message when termination is suspended due to the time of day filter
	msgTimeOfDayExcluded = "time of day excluded"
	// msgDayOfYearExcluded is the log message when termination is suspended due to the day of year filter
	msgDayOfYearExcluded = "day of year excluded"
)

// New returns a new instance of Chaoskube. It expects:
// * a Kubernetes client to connect to a Kubernetes API
// * label, annotation and/or namespace selectors to reduce the amount of possible target pods
// * a list of weekdays, times of day and/or days of a year when chaos mode is disabled
// * a time zone to apply to the aforementioned time-based filters
// * a logger implementing logrus.FieldLogger to send log output to
// * what specific terminator to use to imbue chaos on victim pods
// * whether to enable/disable dry-run mode
func New(client kubernetes.Interface, labels, annotations, kinds, namespaces, namespaceLabels labels.Selector, includedPodNames, excludedPodNames *regexp.Regexp, excludedWeekdays []time.Weekday, excludedTimesOfDay []util.TimePeriod, excludedDaysOfYear []time.Time, timezone *time.Location, minimumAge time.Duration, logger log.FieldLogger, dryRun bool, terminator terminator.Terminator, maxKill int, notifier notifier.Notifier, clientNamespaceScope string) *Chaoskube {
	_ = "STUB: not implemented"
	return nil
}

// Run continuously picks and terminates a victim pod at a given interval
// described by channel next. It returns when the given context is canceled.
func (c *Chaoskube) Run(ctx context.Context, next <-chan time.Time) {
	_ = "STUB: not implemented"
	return
}

// TerminateVictims picks and deletes a victim.
// It respects the configured excluded weekdays, times of day and days of a year filters.
func (c *Chaoskube) TerminateVictims(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Victims returns up to N pods as configured by MaxKill flag
func (c *Chaoskube) Victims(ctx context.Context) ([]v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Candidates returns the list of pods that are available for termination.
// It returns all pods that match the configured label, annotation and namespace selectors.
func (c *Chaoskube) Candidates(ctx context.Context) ([]v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeletePod deletes the given pod with the selected terminator.
// It will not delete the pod if dry-run mode is enabled.
func (c *Chaoskube) DeletePod(ctx context.Context, victim v1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// return early if we're running in dryRun mode.

// filterByKinds filters a list of pods by a given kind selector.
func filterByKinds(pods []v1.Pod, kinds labels.Selector) ([]v1.Pod, error) {
	_ = "STUB: not implemented"
	// empty filter returns original list
	return nil, nil
}

// split requirements into including and excluding groups

// if there aren't any including requirements, we're in by default

// Check owner reference

// convert the pod's owner kind to an equivalent label selector

// include pod if one including requirement matches

// exclude pod if it is filtered out by at least one excluding requirement

// filterByNamespaces filters a list of pods by a given namespace selector.
func filterByNamespaces(pods []v1.Pod, namespaces labels.Selector) ([]v1.Pod, error) {
	_ = "STUB: not implemented"
	// empty filter returns original list
	return nil, nil
}

// split requirements into including and excluding groups

// if there aren't any including requirements, we're in by default

// convert the pod's namespace to an equivalent label selector

// include pod if one including requirement matches

// exclude pod if it is filtered out by at least one excluding requirement

// filterPodsByNamespaceLabels filters a list of pods by a given label selector on their namespace.
func filterPodsByNamespaceLabels(ctx context.Context, pods []v1.Pod, labels labels.Selector, client kubernetes.Interface) ([]v1.Pod, error) {
	_ = "STUB: not implemented"
	// empty filter returns original list
	return nil, nil
}

// find all namespaces matching the label selector

// include pod if its in one of the matched namespaces

// filterByAnnotations filters a list of pods by a given annotation selector.
func filterByAnnotations(pods []v1.Pod, annotations labels.Selector) []v1.Pod {
	_ = "STUB: not implemented"
	// empty filter returns original list
	return nil
}

// convert the pod's annotations to an equivalent label selector

// include pod if its annotations match the selector

// filterByPhase filters a list of pods by a given PodPhase, e.g. Running.
func filterByPhase(pods []v1.Pod, phase v1.PodPhase) []v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// filterTerminatingPods removes pod which have a non nil DeletionTimestamp
func filterTerminatingPods(pods []v1.Pod) []v1.Pod { _ = "STUB: not implemented"; return nil }

// filterByMinimumAge filters pods by creation time. Only pods
// older than minimumAge are returned
func filterByMinimumAge(pods []v1.Pod, minimumAge time.Duration, now time.Time) []v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// filterByPodName filters pods by name.  Only pods matching the includedPodNames and not
// matching the excludedPodNames are returned
func filterByPodName(pods []v1.Pod, includedPodNames, excludedPodNames *regexp.Regexp) []v1.Pod {
	_ = "STUB: not implemented"
	// return early if neither included nor excluded regular expressions are given
	return nil
}

func filterByOwnerReference(pods []v1.Pod) []v1.Pod { _ = "STUB: not implemented"; return nil }

// Don't filter out pods with no owner reference

// Group remaining pods by their owner reference

// For each owner reference select a random pod from its group
