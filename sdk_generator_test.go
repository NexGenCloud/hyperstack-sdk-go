package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/NexGenCloud/hyperstack-sdk-go/lib/auth"
)

var (
	API_SERVER         = "https://infrahub-api.nexgencloud.com/v1"
	API_SERVER_STAGING = "https://infrahub-api-stg.ngbackend.cloud/v1"
)

type CustomTime struct {
	time.Time
}

type UserInfoResponse struct {
	Message *string     `json:"message,omitempty"`
	Status  *bool       `json:"status,omitempty"`
	User    *UserFields `json:"user,omitempty"`
}

const ctLayout = "2006-01-02T15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

// And use this in your structs:

type UserFields struct {
	CreatedAt *CustomTime `json:"created_at,omitempty"`
	Email     *string     `json:"email,omitempty"`
	Name      *string     `json:"name,omitempty"`
	Username  *string     `json:"username,omitempty"`
}

func createAuthClient() (*auth.ClientWithResponses, error) {
	apiKey := os.Getenv("HYPERSTACK_API_KEY")
	staging := os.Getenv("HYPERSTACK_STAGING") == "false"

	apiServer := API_SERVER
	if staging {
		apiServer = API_SERVER_STAGING
	}

	addHeaders := func(ctx context.Context, req *http.Request) error {
		req.Header.Add("api_key", apiKey)
		// TODO: do we need to support it?
		//req.Header.Add("Authorization", "Bearer "+token)
		return nil
	}

	authClient, err := auth.NewClientWithResponses(apiServer, auth.WithRequestEditorFn(addHeaders))
	if err != nil {
		return nil, err
	}

	return authClient, nil
}

func TestAuthUserInformation(t *testing.T) {
	authClient, err := createAuthClient()

	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}

	if err != nil {
		t.Fatalf("failed to create auth client: %s", err)
	}
	resp, err := authClient.RetrieveAuthenticatedUserDetailsWithResponse(context.Background())

	if err != nil {
		t.Fatalf("failed to get user information: %s", err)
	}

	if resp.HTTPResponse != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.HTTPResponse.Body)
		fmt.Printf("Raw Response from server: %s\n", string(bodyBytes))
	}

	if resp.JSON401 != nil {
		fmt.Printf("SDK Connection OK, Short test circuit due to lack of token")
	} else {

		if resp.JSON200 == nil {
			t.Fatalf("expected json response, got none")
		}

		user := resp.JSON200.User
		if user == nil {
			t.Fatalf("expected user information, got none")
		}

		fmt.Printf("User name: %s\n", derefString(user.Name))
		fmt.Printf("User email: %s\n", derefString(user.Email))
		fmt.Printf("User username: %s\n", derefString(user.Username))
		fmt.Printf("Created At: %s\n", user.CreatedAt)
		// Add more fields as necessary
	}
}

// derefString dereferences a string pointer and returns
// an empty string if the pointer is nil
func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
