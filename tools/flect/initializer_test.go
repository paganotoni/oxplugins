package flect

import (
	"context"
	"os"
	"testing"
)

func TestInitializer(t *testing.T) {
	t.Run("InflectionsFileDoesNotExist", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{})

		if err != nil {
			t.Fatalf("error should be nill, got %v", err)
		}

		_, err = os.Stat(root)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create file ")
		}

	})
	t.Run("inflectionFileExist", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		rootYml := root + "/inflections.yml"
		_, err = os.Create(rootYml)
		if err != nil {
			t.Fatalf("Problem creating file, %v", err)
		}

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{})

		if err != nil {
			t.Fatalf("error should be nill, got %v", err)
		}

		_, err = os.Stat(root)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create file ")
		}

	})
}
