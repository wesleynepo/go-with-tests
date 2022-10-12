package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"testing/fstest"
)

type Post struct {
    Title string
    Description string
    Tags []string
    Body string
}

const (
    titleSeparator = "Title: "
    descriptionSeparator = "Description: "
    tagsSeparator = "Tags: "
)

func NewPostsFromFS(fileSystem fstest.MapFS) ([]Post, error) {
    dir, err := fs.ReadDir(fileSystem, ".")

    if err != nil {
        return nil, err
    }
    var posts []Post

    for _, f := range dir {
        post, err := getPost(fileSystem, f.Name())
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    return posts, nil
}

func getPost(fileSystem fs.FS, f string) (Post, error) {
    postFile, err := fileSystem.Open(f)

    if err != nil {
        return Post{}, err
    }

    defer postFile.Close()

    return newPost(postFile)
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

