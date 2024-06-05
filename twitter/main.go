package twitter

import (
	"context"
	"fmt"
	"log"

	_ "unsafe"

	"github.com/gofrs/uuid"
)

//go:linkname Uint32 runtime.fastrand
func Uint32() uint32

func Twittermain() {

	///////
	id, err := uuid.NewV4()
	root := "http://44.234.134.96:8888/callback/x"
	redirectURI := fmt.Sprintf("%s/oauth", root)
	oAuth := NewTwitterOAuth(redirectURI, id.String()) // Change to NewGitOAuth() to use Github api
	// Generate authorization URL
	authURL, err := oAuth.GetAuthURI(context.Background())
	if err != nil {
		log.Printf("loginHandler Error: %v", err)
		// http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(authURL)
}
