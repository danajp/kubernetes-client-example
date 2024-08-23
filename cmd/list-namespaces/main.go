// Simple program to list namespaces using the controller-runtime and
// client-go clientset clients
//
// This will read the kubernetes client config as you would expect,
// respecting ~/.kube/config, the KUBECONFIG env var or using service
// account crendentials when running as a pod in a cluster.
package main

import (
	"context"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var scheme = runtime.NewScheme()

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	// add schemes for custom resources here to use with
	// controller-runtime client
}

func listNamespacesControllerRuntime(cfg *rest.Config) error {
	fmt.Println("list namespaces using controller-runtime client")

	k8sClient, err := client.New(cfg, client.Options{})
	if err != nil {
		return err
	}

	nsList := v1.NamespaceList{}

	if err := k8sClient.List(context.TODO(), &nsList); err != nil {
		return err
	}

	fmt.Println("NAMESPACES")
	for _, ns := range nsList.Items {
		fmt.Println(ns.Name)
	}

	return nil
}

func listNamespacesClientset(cfg *rest.Config) error {
	fmt.Println("list namespaces using client-go clientset")

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	nsList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	fmt.Println("NAMESPACES")
	for _, ns := range nsList.Items {
		fmt.Println(ns.Name)
	}

	return nil
}

func main() {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := listNamespacesControllerRuntime(config); err != nil {
		log.Fatal(err)
	}

	if err := listNamespacesClientset(config); err != nil {
		log.Fatal(err)
	}
}
