package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/wesleynepo/blogposts"
)

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
    t.Helper()

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %+v, want %+v", got, want)
    }
}


func TestNewBlogPosts(t *testing.T) {
    const (
        firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
        secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
Last
Mode
`
    )

    fs := fstest.MapFS {
        "hello world.md": {Data: []byte(firstBody)},
        "hello-world2.md": {Data: []byte(secondBody)},
    }

    posts, _ := blogposts.NewPostsFromFS(fs)

    got := posts[0]
    want := blogposts.Post{
        Title: "Post 1",
        Description: "Description 1",
        Tags: []string{"tdd", "go"},
        Body: `Hello
World`,
    }
    assertPost(t, got, want)
}


type StubFailingFS struct {}

func (s StubFailingFS) Open(name string) (fs.File, error) {
    return nil, errors.New("oh no, i always fail")
}
