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

	reproID := uuid.NewString()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}

	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		eg.Go(func() error {
			return doSomething(ctx, dag, reproID)
		})
	}
	return eg.Wait()
}

func doSomething(ctx context.Context, dag *dagger.Client, reproID string) error {
	_, err := dag.Container().
		From("alpine").
		WithExec([]string{"sleep", "5"}).             // cached across all repros
		WithEnvVariable("REPRO_ID", reproID).         // cache buster
		WithExec([]string{"sleep", "5"}).             // cached once within repro
		WithExec([]string{"echo", uuid.NewString()}). // never cached
		Sync(ctx)
	return err
}
