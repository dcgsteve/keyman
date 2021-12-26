package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

const GitDirName = ".git"

func main() {
	// Filesystem abstraction based on memory
	fs := memfs.New()
	// Git objects storer based on memory
	storer := memory.NewStorage()

	// Clones the repository into the worktree (fs) and stores all the .git
	// content into the storer
	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://github.com/git-fixtures/basic.git",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Prints the content of the CHANGELOG file from the cloned repository
	changelog, err := fs.Open("CHANGELOG")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, changelog)

}
