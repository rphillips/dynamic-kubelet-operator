package stub

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/rphillips/dynamic-kubelet-operator/pkg/apis/dynamickubelet/v1alpha1"
	"github.com/rphillips/dynamic-kubelet-operator/pkg/reconciler"
)

func NewHandler(r *reconciler.Reconciler) sdk.Handler {
	return &Handler{
		Reconciler: r
	}
}

type Handler struct {
	Reconciler *reconciler.Reconciler
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.DynamicKubeletOperator:
		log.Println("Dynamic Kubelet Operator")
	case *corev1.Node:
		if event.Deleted {
			log.Printf("Got delete event %v/%v", o.GetNamespace(), o.GetName())
			return nil
		}
		log.Printf("Got event %v/%v", o.GetNamespace(), o.GetName())
		return h.Reconcile(ctx, o)
	}
	return nil
}
