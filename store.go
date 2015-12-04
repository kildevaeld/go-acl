package acl

type MemoryStore struct {
	roles map[string]*Role
}

func (self *MemoryStore) GetRole(name string) *Role {
	return self.roles[name]
}

func (self *MemoryStore) 	AddRole(role *Role) error {
	self.roles[role.Name] = role
	return nil
}

func (self *MemoryStore) 	UpdateRole(role *Role) error {
	self.roles[role.Name] = role
	return nil
}

func (self *MemoryStore) 	RemoveRole(role *Role) error {
	delete(self.roles, role.Name)
	return nil
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{make(map[string]*Role)}
}