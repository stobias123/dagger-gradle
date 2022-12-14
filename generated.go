// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"

	"go.dagger.io/dagger/sdk/go/dagger"
)

func (r *query) gradle(ctx context.Context) (*Gradle, error) {

	return new(Gradle), nil

}

type gradle struct{}
type query struct{}

func main() {
	dagger.Serve(context.Background(), map[string]func(context.Context, dagger.ArgsInput) (interface{}, error){
		"Gradle.run": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			var fs dagger.FSID

			bytes, err = json.Marshal(fc.Args["fs"])
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(bytes, &fs); err != nil {
				return nil, err
			}

			var command string

			bytes, err = json.Marshal(fc.Args["command"])
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(bytes, &command); err != nil {
				return nil, err
			}

			return (&gradle{}).run(ctx,

				fs,

				command,
			)
		},
		"Query.gradle": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			return (&query{}).gradle(ctx)
		},
	})
}
