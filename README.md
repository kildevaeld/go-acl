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

type User struct {
  Id int
  Name string
}

// ACLType interface
func (u *User) ACLType () string {
  return "user"
}
// ACLIdentity interface
func (u *User) ACLIdentifier () string {
  return fmt.Sprintf("%d", b.Id)
}

u := &User{
  Id: 2, 
  Name: "Salty Mcgee",
}

b := &Blog{
  Id: 1,
  Title: "Blog title",
  Body: "Some body text",
}

a.Allow(u, "modify", b)
// same as: a.Allow("user:2", "modify", "blog:1")
a.Can(u, "modify", b)
// a.Can("user:2", "modify", "blog:1")

```
