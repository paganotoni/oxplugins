package docker

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestInitilizer(t *testing.T) {
	t.Run("dockerFileDoesNotExist", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{})
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		rootDoc := root + "/.dockerignore"
		rootFile := root + "/Dockerfile"

		_, err = os.Stat(rootDoc)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create .dockerignore file , %v", err)
		}
		_, err = os.Stat(rootFile)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create  Dockerfile file , %v", err)
		}
	})

	t.Run("dockerFileDoesExist", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		rootFile := filepath.Join(root, "Dockerfile")
		_, err = os.Create(rootFile)

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{})
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}
		_, err = os.Stat(rootFile)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create  Dockerfile file , %v", err)
		}

	})
	t.Run("dockerIgnoreDoesExist", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		rootdoc := filepath.Join(root, ".dockerignore")
		_, err = os.Create(rootdoc)

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{})
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}
		_, err = os.Stat(rootdoc)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create .dockerignore file , %v", err)
		}
	})
}
