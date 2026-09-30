package appsignal_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Drieam/appsignal-client-go"
)

func ExampleClient_GetLogSourceExport() {
	client := appsignal.NewClient("", os.Getenv("APPSIGNAL_API_TOKEN"))

	export, err := client.GetLogSourceExport(context.Background(), "6aad24d1ba6bc351255e7cb5", "6aad24d1ba6bc351255e7cba")
	if err != nil {
		log.Fatal(err)
	}
	if export == nil {
		log.Fatal("log source has no export")
	}

	fmt.Println(export.S3URL)
}

func ExampleClient_CreateLogSourceExport() {
	client := appsignal.NewClient("", os.Getenv("APPSIGNAL_API_TOKEN"))

	export, err := client.CreateLogSourceExport(context.Background(), appsignal.CreateLogSourceExportInput{
		AppID:       "6aad24d1ba6bc351255e7cb5",
		LogSourceID: "6aad24d1ba6bc351255e7cba",
		S3URL:       "https://s3.amazonaws.com/my-logs-bucket",
		S3AccessKey: "AKIAIOSFODNN7EXAMPLE",
		S3SecretKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		Compressed:  true,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(export.S3URL)
}

func ExampleClient_UpdateLogSourceExport() {
	client := appsignal.NewClient("", os.Getenv("APPSIGNAL_API_TOKEN"))

	s3Url := "https://s3.amazonaws.com/my-new-logs-bucket"

	export, err := client.UpdateLogSourceExport(context.Background(), appsignal.UpdateLogSourceExportInput{
		AppID:       "6aad24d1ba6bc351255e7cb5",
		LogExportID: "65d893877ccda67f2e172495",
		S3URL:       &s3Url,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(export.S3URL)
}

func ExampleClient_DeleteLogSourceExport() {
	client := appsignal.NewClient("", os.Getenv("APPSIGNAL_API_TOKEN"))

	export, err := client.DeleteLogSourceExport(context.Background(), appsignal.DeleteLogSourceExportInput{
		AppID:       "6aad24d1ba6bc351255e7cb5",
		LogExportID: "65d893877ccda67f2e172495",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted", export.S3URL)
}
