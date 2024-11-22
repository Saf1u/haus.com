package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var hausGVR = schema.GroupVersionResource{
	Group:    "haus.com",
	Version:  "v1",
	Resource: "hauses",
}

func main() {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Load the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}
	cl, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	resource, err := cl.Resource(hausGVR).Namespace("default").Get(context.Background(), "sample-haus", v1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Print(resource.GetGeneration())

}
