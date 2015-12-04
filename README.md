# go-acl

# Basic Usage

```go
import "github.com/kildevaeld/go-acl"

a := acl.NewAcl(acl.NewMemoryStore())

a.Role("guest","")
a.Role("user", "guest") // user inherits guest
a.Role("admin", "user") // admin inherits user and guest

a.Allow([]string{"user", "admin"}, "comment", "blog") // multiple
a.Allow("guest", "read", "blog") 

a.Can("user", "comment", "blog") // true
a.Can([]string{"user"}, "read", "blog") // true

a.Can([]string{"guest"}, "comment", "blog") // false

```

Or conforming to interfaces 

```go 

type Blog struct {
  Id    int
  Title string
  Body  string
}
// ACLType interface
func (b *Blog) ACLType () string {
  return "blog"
}
// ACLIdentity interface
func (b *Blog) ACLIdentifier () string {
  return fmt.Sprintf("%d", b.Id)
}

b := &Blog{
  Id: 1,
  Title: "Blog title",
  Body: "Some body text",
}

a.Allow([]string{"user"}, "modify", b)
// same as: a.Allow([]string{"user"}, "modify", "blog:1")

```
