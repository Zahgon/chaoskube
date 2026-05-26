package terminator

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeletePodTerminator simply asks k8s to delete the victim pod.
type DeletePodTerminator struct {
	client      kubernetes.Interface
	logger      log.FieldLogger
	gracePeriod time.Duration
}

// NewDeletePodTerminator creates and returns a DeletePodTerminator object.
func NewDeletePodTerminator(client kubernetes.Interface, logger log.FieldLogger, gracePeriod time.Duration) *DeletePodTerminator {
	_ = "STUB: not implemented"
	return nil
}

// Terminate sends a request to Kubernetes to delete the pod.
func (t *DeletePodTerminator) Terminate(ctx context.Context, victim v1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteOptions(gracePeriod time.Duration) metav1.DeleteOptions {
	_ = "STUB: not implemented"
	return *new(metav1.DeleteOptions)
}
