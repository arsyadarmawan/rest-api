package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/google"
	"log"
	"net/url"
	"os"
	"time"
)

type GCSObjectList struct {
	Items []struct {
		Kind                    string    `json:"kind"`
		ID                      string    `json:"id"`
		SelfLink                string    `json:"selfLink"`
		MediaLink               string    `json:"mediaLink"`
		Name                    string    `json:"name"`
		Bucket                  string    `json:"bucket"`
		Generation              string    `json:"generation"`
		Metageneration          string    `json:"metageneration"`
		ContentType             string    `json:"contentType"`
		StorageClass            string    `json:"storageClass"`
		Size                    string    `json:"size"`
		Md5Hash                 string    `json:"md5Hash"`
		Crc32C                  string    `json:"crc32c"`
		Etag                    string    `json:"etag"`
		TemporaryHold           bool      `json:"temporaryHold,omitempty"`
		EventBasedHold          bool      `json:"eventBasedHold,omitempty"`
		TimeCreated             time.Time `json:"timeCreated"`
		Updated                 time.Time `json:"updated"`
		TimeStorageClassUpdated time.Time `json:"timeStorageClassUpdated"`
	} `json:"items"`
}

const (
	keyFilePath = "/Users/arsyadthareeq/Personal/credentials/gcp_arsyad.json"
)

func main() {
	bucketName := "development-me"
	folderName := "arsyadthareq"

	isExist, err := folderExists(bucketName, folderName, keyFilePath)
	if err != nil {
		fmt.Sprintf("gada dong")
	}
	if !isExist {
		if errCreate := createFolder(bucketName, folderName, keyFilePath); errCreate != nil {
			fmt.Errorf(errCreate.Error())
		}
		//getRootFolderForCLient(bucketName, folderName)
	}
	//if errDelete := deleteFolder(bucketName, folderName, keyFilePath); errDelete != nil {
	//
	//	fmt.Errorf("Error deleting folder: %v", errDelete)
	//}
	fmt.Println("success")
}

func deleteFolder(bucketName, folderName, keyFilePath string) error {
	// Get the access token
	token, err := getAccessToken(keyFilePath)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	// Initialize Resty client
	client := resty.New()

	// Step 1: List all objects in the folder
	listURL := fmt.Sprintf("https://storage.googleapis.com/storage/v1/b/%s/o", bucketName)
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"prefix":    folderName + "/",
			"delimiter": "/",
			// Prefix for "folder"
		}).
		SetAuthToken(token).
		Get(listURL)

	if err != nil {
		return fmt.Errorf("failed to list objects in folder: %w", err)
	}
	if resp.IsError() {
		return fmt.Errorf("error listing objects: %s", resp.Status())
	}

	// Parse the response to get object names
	var result GCSObjectList
	if err = json.Unmarshal(resp.Body(), &result); err != nil {
		return fmt.Errorf("failed to parse object list: %w", err)
	}

	if len(result.Items) == 0 {
		fmt.Printf("No objects found in folder %s; nothing to delete.\n", folderName)
		return nil
	}

	// Step 2: Delete each object in the folder
	for _, item := range result.Items {
		// Construct URL for each object to delete
		encodedName := url.PathEscape(item.Name)
		deleteURL := fmt.Sprintf("https://storage.googleapis.com/storage/v1/b/%s/o/%s", bucketName, encodedName)

		deleteResp, err := client.R().
			SetAuthToken(token).
			Delete(deleteURL)
		if err != nil {
			return fmt.Errorf("failed to delete object %s: %w", item.Name, err)
		}
		if deleteResp.IsError() {
			return fmt.Errorf("error deleting object %s: %s", item.Name, deleteResp.Status())
		}
		fmt.Printf("Deleted object: %s\n", item.Name)
	}
	fmt.Printf("Folder %s and all contents deleted successfully\n", folderName)
	return nil
}

func getRootFolderForCLient(bucketName, folderName string) {
	token, err := getAccessToken(keyFilePath)
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	// Initialize Resty client
	client := resty.New()
	url := fmt.Sprintf("https://storage.googleapis.com/storage/v1/b/%s/o", bucketName)
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"prefix":    folderName + "/",
			"delimiter": "/",
		}).
		SetAuthToken(token).
		Get(url)

	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	// Parse and display response
	var objectList GCSObjectList
	if err := json.Unmarshal(resp.Body(), &objectList); err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}
}

func folderExists(bucketName, folderName, keyFilePath string) (bool, error) {
	// Get the access token
	token, err := getAccessToken(keyFilePath)
	if err != nil {
		return false, fmt.Errorf("failed to get access token: %w", err)
	}

	// Initialize Resty client
	client := resty.New()

	// Make a request to list objects with the folder prefix
	url := fmt.Sprintf("https://storage.googleapis.com/storage/v1/b/%s/o", bucketName)
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"prefix":     folderName + "/",
			"maxResults": "1", // Limit results for efficiency
		}).
		SetAuthToken(token).
		Get(url)

	if err != nil {
		return false, fmt.Errorf("request failed: %w", err)
	}
	if resp.StatusCode() == 200 {
		var objectList GCSObjectList
		if err := json.Unmarshal(resp.Body(), &objectList); err != nil {
			log.Fatalf("Failed to parse response: %v", err)
		}
		if len(objectList.Items) > 0 {
			return true, nil
		}
	}
	return false, nil
}

func getAccessToken(keyFilePath string) (string, error) {
	// Read the service account JSON key file
	jsonKey, err := os.ReadFile(keyFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read key file: %w", err)
	}

	// Parse JSON credentials to obtain token
	credentials, err := google.CredentialsFromJSON(context.Background(), jsonKey, "https://www.googleapis.com/auth/devstorage.full_control")
	if err != nil {
		return "", fmt.Errorf("failed to parse credentials: %w", err)
	}

	// Get token from the credentials
	tokenSource := credentials.TokenSource
	token, err := tokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token.AccessToken, nil
}

func createFolder(bucketName, folderName, keyFilePath string) error {
	// Get the access token
	token, err := getAccessToken(keyFilePath)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	// Initialize Resty client
	client := resty.New()

	// The GCS URL for uploading objects
	//url := fmt.Sprintf("https://storage.googleapis.com/upload/storage/v1/b/%s/o?uploadType=media&name=%s", bucketName, folderName+"/")
	url := fmt.Sprintf("https://storage.googleapis.com/upload/storage/v1/b/%s/o?uploadType=media&name=%s/", bucketName, folderName)
	// Make the request to create an empty object with the folder name
	"https://storage.googleapis.com/storage/v1/b/development-me/o?uploadType=media&name=arsyadthareq/"
	resp, err := client.R().
		SetAuthToken(token).
		SetHeader("Content-Type", "application/octet-stream").
		SetBody("").
		Post(url)

	if err != nil {
		return fmt.Errorf("failed to create folder: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("error response from GCS: %s", resp.Status())
	}

	fmt.Printf("Folder %s created successfully in bucket %s\n", folderName, bucketName)
	return nil
}
