package e2e

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

// newFakeClientset wraps fake.NewClientset for unit tests.
func newFakeClientset(objects ...runtime.Object) *fake.Clientset {
	return fake.NewClientset(objects...)
}
