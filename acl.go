package acl

import "fmt"

type ACLType interface {
	ACLType() string
}

type ACLIdentity interface {
	ACLIdentity() string
}

type Permission struct {
	Action   string
	Resource string
}

func pcontains(ps []Permission, p *Permission) bool {
	for _, i := range ps {
		if i.Action == p.Action && i.Resource == p.Resource {
			return true
		}
	}
	return false
}

type Role struct {
	Name        string
	Parent      string
	Permissions []Permission
}

func (self *Role) hasPermission(store Store, perm *Permission) bool {

	if pcontains(self.Permissions, perm) {
		return true
	}

	if self.Parent != "" {
		parent := store.GetRole(self.Parent)
		if parent != nil {
			return parent.hasPermission(store, perm)
		}
	}

	return false
}

/*type ACLRoles interface {
	ACLRoles() []string
}*/

type Store interface {
	GetRole(name string) *Role
	AddRole(role *Role) error
	UpdateRole(role *Role) error
	RemoveRole(role *Role) error
}

type ACL struct {
	store        Store
	IgnoreErrors bool
}

func (self *ACL) Role(name string, inherits string) *ACL {

	role := self.store.GetRole(name)

	if role != nil {
		return self
	}

	role = &Role{
		Name:        name,
		Permissions: []Permission{},
	}

	if inherits != "" {
		parent := self.store.GetRole(inherits)
		if parent == nil {
			panic(fmt.Errorf("parent role '%s' does not exists", inherits))
		}
		role.Parent = parent.Name
	}

	err := self.store.AddRole(role)

	if err != nil && !self.IgnoreErrors {
		panic(err)
	}

	return self
}

func (self *ACL) Allow(roles []string, action string, resource interface{}) *ACL {

	resourceStr := resourceStringFromInterface(resource)

	if resourceStr == "" {
		panic(fmt.Errorf("resource should be a string, ACLType or ACLIdentity. Was: %v", resource))
	}

	permission := Permission{
		Action:   action,
		Resource: resourceStr,
	}

	for _, roleName := range roles {

		role := self.store.GetRole(roleName)

		if role == nil && !self.IgnoreErrors {
			panic(fmt.Errorf("role '%s' does not exists", roleName))
		}

		if !role.hasPermission(self.store, &permission) {
			role.Permissions = append(role.Permissions, permission)
			self.store.UpdateRole(role)
		}
	}

	return self
}

func (self *ACL) HasRole(name string) bool {
	return self.store.GetRole(name) != nil
}

func (self *ACL) RemoveRole(name string) error {
	role := self.store.GetRole(name)
	if role != nil {
		return self.store.RemoveRole(role)
	}
	return nil
}

func (self *ACL) Can(roles []string, action string, resource interface{}) bool {

	resourceStr := resourceStringFromInterface(resource)

	if resourceStr == "" {
		return false
	}

	permission := Permission{
		Action:   action,
		Resource: resourceStr,
	}

	for _, roleName := range roles {

		role := self.store.GetRole(roleName)

		if role == nil && !self.IgnoreErrors {
			panic(fmt.Errorf("role '%s' does not exists", roleName))
		}

		if role.hasPermission(self.store, &permission) {
			return true
		}
	}

	return false
}

func resourceStringFromInterface(resource interface{}) string {
	resourceStr := ""

	if s, ok := resource.(string); ok {
		resourceStr = s
	} else {

		if i, ok := resource.(ACLType); ok {
			resourceStr = i.ACLType()
		}

		if i, ok := resource.(ACLIdentity); ok {
			if len(resourceStr) > 0 {
				resourceStr += ":"
			}
			resourceStr += i.ACLIdentity()
		}

	}

	return resourceStr

}

func New(store Store) *ACL {
	return &ACL{
		store: store,
	}
}
