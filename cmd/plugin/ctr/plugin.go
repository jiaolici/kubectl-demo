package main

import (
	"fmt"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func RunPlugin(configFlags *genericclioptions.ConfigFlags, outputCh chan string) error {
	//fmt.Println("RunPlugin start")
	//config, err := configFlags.ToRESTConfig()
	//if err != nil {
	//	return fmt.Errorf("failed to read kubeconfig: %w", err)
	//}

	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	return fmt.Errorf("failed to create clientset: %w", err)
	//}

	//namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	//if err != nil {
	//	return fmt.Errorf("failed to list namespaces: %w", err)
	//}
	//
	//for _, namespace := range namespaces.Items {
	//	outputCh <- fmt.Sprintf("Namespace %s", namespace.Name)
	//	time.Sleep(time.Millisecond * 100)
	//}
	fmt.Println("RunPlugin end")
	return nil
}
