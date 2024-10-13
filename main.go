package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	clientset "github.com/saf1u/haus.com/pkg/client/clientset/versioned/typed/haus.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

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

	hausClient, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	cl := hausClient.Hauses("default")
	haus, err := cl.Get(context.Background(), "sample-haus", v1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(haus.Spec.Location)
	fmt.Println(haus.Spec.Count)

}
