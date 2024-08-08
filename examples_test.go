package dynamodblocal_test

import (
	"context"
	"fmt"
	"log"

	dynamodblocal "github.com/shasderias/dynamodb-local-testcontainers-go"
)

func ExampleRun() {
	ctx := context.Background()

	dynamodbContainer, err := dynamodblocal.Run(ctx, "amazon/dynamodb-local:2.2.1")
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := dynamodbContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err) // nolint:gocritic
		}
	}()

	state, err := dynamodbContainer.State(ctx)
	if err != nil {
		log.Fatalf("failed to get container state: %s", err)
	}

	fmt.Println(state.Running)

	// Output:
	// true
}
