package examples

import (
	"context"
	corev1 "k8s.io/api/core/v1"
)

//+cloudnativeflow:scaffold:flowstep=true
func CreatePod(context context.Context, podName string) (newPodName string, err error) {
	return "", nil
}

//+cloudnativeflow:scaffold:flowstep=true
func ListPod(context context.Context, newPodName string) (podList []*corev1.Pod, err error) {
	return nil, nil
}
