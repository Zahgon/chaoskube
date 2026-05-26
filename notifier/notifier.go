package notifier

import (
	v1 "k8s.io/api/core/v1"
)

type Notifier interface {
	NotifyPodTermination(pod v1.Pod) error
}

type Notifiers struct {
	notifiers []Notifier
}

func New() *Notifiers { _ = "STUB: not implemented"; return nil }

func (m *Notifiers) NotifyPodTermination(pod v1.Pod) error { _ = "STUB: not implemented"; return nil }

func (m *Notifiers) Add(notifier Notifier) { _ = "STUB: not implemented"; return }
