package main

import (
	"context"

	"github.com/stobias123/dagger-gradle/ext/gen/core"
	"go.dagger.io/dagger/sdk/go/dagger"
)

func (r *gradle) run(ctx context.Context, fs dagger.FSID, command string) (*dagger.Filesystem, error) {
	buildFS, err := core.Image(ctx, "gradle:7.5.1-jdk11-focal")
	if err != nil {
		return nil, err
	}

	outputFS, err := core.Exec(ctx, buildFS.Core.Image.ID, core.ExecInput{
		Args:    []string{"gradle", "--no-daemon", "server:jib"},
		Workdir: "/app",
		Env: []core.ExecEnvInput{
			{
				Name:  "JAVA_HOME",
				Value: "/opt/java/openjdk/",
			},
		},
		Mounts: []core.MountInput{
			{
				Fs:   fs,
				Path: "/app",
			},
		},
	})

	return &outputFS.Core.Filesystem, err
}
