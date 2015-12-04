# go-acl

# Basic Usage

```go
import "github.com/kildevaeld/go-acl"

a := acl.NewAcl(acl.NewMemoryStore())

a.Role("guest","")
a.Role("user", "guest") // user inherits guest
a.Role("admin", "user") // admin inherits user and guest

a.Allow([]string{"user"}, "comment", "blog")
a.Allow([]string{"guest"}, "read", "blog")

a.Can([]string{"user"}, "comment", "blog") // true
a.Can([]string{"user"}, "read", "blog") // true

a.Can([]string{"guest"}, "comment", "blog") // false

```
