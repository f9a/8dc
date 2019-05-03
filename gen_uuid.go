package main

import (
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

func main() {
	domain := "uuids.posts.0fd.x133.me"
	namespace := uuid.NewV5(uuid.NamespaceDNS, domain)
	fmt.Println("Namespace (UUIDv5): ", namespace)

	randomID, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Cannot create key for namespace: ", err)
		os.Exit(1)
	}
	fmt.Println("Bucket-Key (UUIDv4): ", randomID)

	id := uuid.NewV5(namespace, randomID.String())

	fmt.Println("Super-ID (UUIDv5): ", id)
}
