package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"dagger.io/dagger"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func repro() error {
	ctx := context.TODO()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}

	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		eg.Go(func() error {
			return doSomething(ctx, dag)
		})
	}
	return eg.Wait()
}

func doSomething(ctx context.Context, dag *dagger.Client) error {
	_, err := dag.Container().
		From("alpine").
		WithExec([]string{"sleep", "5"}).             // cached exec op
		WithExec([]string{"echo", uuid.NewString()}). // always not-cached exec op
		Sync(ctx)
	return err
}
