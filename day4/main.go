package main

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
	// 创建clientset
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/<username>/.kube/config")
	clientset, _ := kubernetes.NewForConfig(config)

	factory := informers.NewSharedInformerFactory(clientset, 30*time.Second)

	informer := factory.Core().V1().Pods().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldPod := oldObj.(*corev1.Pod)
			newPod := newObj.(*corev1.Pod)
			if oldPod.ResourceVersion != newPod.ResourceVersion {
				fmt.Printf("Pod %s was updated\n", newPod.Name)
				msg := fmt.Sprintf("Pod %s has been updated\n", newPod.Name)
				fmt.Println(msg)
				//sendDingtalkMessage(msg)

			}
		},
	})

	stopper := make(chan struct{})
	defer close(stopper)
	informer.Run(stopper)
}
