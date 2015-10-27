package main

import (
	"fmt"

	"github.com/coreos/rkt/api/v1alpha"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// Create client.
	fmt.Println("Connecting to api service...")

	clientConn, err := grpc.Dial("localhost:15441", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer clientConn.Close()
	client := v1alpha.NewPublicAPIClient(clientConn)
	fmt.Println("Successfully connected!")

	// Get version infos.
	fmt.Println("\nGet info")
	getInfo, err := client.GetInfo(context.Background(), &v1alpha.GetInfoRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(getInfo)

	// Get pod list.
	fmt.Println("\nGet pod list")
	podInfos, err := client.ListPods(context.Background(), &v1alpha.ListPodsRequest{&v1alpha.PodFilter{States: []v1alpha.PodState{}}})
	if err != nil {
		panic(err)
	}
	for i, p := range podInfos.Pods {
		fmt.Printf("pod #%d: %v\n", i, p)
	}

	// Inspect pod.
	fmt.Println("\nInspect pod")
	podInfo, err := client.InspectPod(context.Background(), &v1alpha.InspectPodRequest{Id: podInfos.Pods[0].Id})
	if err != nil {
		panic(err)
	}
	fmt.Println(podInfo)

	// Get image list.
	fmt.Println("\nGet image list")
	imgInfos, err := client.ListImages(context.Background(), &v1alpha.ListImagesRequest{})
	if err != nil {
		panic(err)
	}
	for i, m := range imgInfos.Images {
		fmt.Printf("image #%d: %v\n", i, m)
	}

	// Inspect image
	fmt.Println("\nInspect image")
	imgInfo, err := client.InspectImage(context.Background(), &v1alpha.InspectImageRequest{Id: imgInfos.Images[0].Id})
	if err != nil {
		panic(err)
	}
	fmt.Println(imgInfo)
}
