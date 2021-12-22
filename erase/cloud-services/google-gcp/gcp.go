// my-go-examples cloud-services/google-gcp

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func main() {

	project := "jeffs-project-174816"
	jsonPath := "/home/jeff/.config/gcloud/jeffs-service-account.json"
	zone := "us-west1-a"

	ctx := context.Background()

	// Reads credentials from the specified path.
	computeService, err := compute.NewService(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatal(err)
	}

	// Use the Service to get gce image list
	fmt.Printf("\n\n**** Use the Service to get gce image list *****************************\n\n")
	req := computeService.Images.List(project)

	if err := req.Pages(ctx, func(page *compute.ImageList) error {
		for _, image := range page.Items {
			// TODO: Change code below to process each `image` resource:
			b, err := json.MarshalIndent(image, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(b)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("----------------")

	// Use the service to get gce instances list
	fmt.Printf("\n\n**** Use the service to get gce instances list *****************************\n\n")
	req2 := computeService.Instances.List(project, zone)

	if err := req2.Pages(ctx, func(page *compute.InstanceList) error {
		for _, image := range page.Items {
			// TODO: Change code below to process each `image` resource:
			b, err := json.MarshalIndent(image, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(b)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

}
