package acl_test

import (
	"testing"

	acl "github.com/kildevaeld/go-acl"
)

func TestRole(t *testing.T) {

	store := acl.NewMemoryStore()
	a := acl.New(store)

	a.Role("user", "")
	
	if !a.HasRole("user") {
		t.Errorf("no role")
	} 

}


func TestAllow (t *testing.T) {
	
	store := acl.NewMemoryStore()
	a := acl.New(store)
	a.Role("guest", "")
	a.Role("user", "guest")
	
	a.Allow([]string{"user"}, "create", "resources")
	a.Allow("guest", "view", "resources")
	
	if !a.Can([]string{"user"}, "create", "resources") {
		t.Errorf("expected user to be able to create resources")
	} 
	if !a.Can([]string{"user"}, "view", "resources") {
		t.Errorf("expected user to be able to view resources")
	} 
	
	if a.Can([]string{"guest"}, "create", "resources") {
		t.Errorf("expected guest not to be able to create resources")
	}
	
	if !a.Can([]string{"guest"}, "view", "resources") {
		t.Errorf("expected guest to be able to view resources")
	}
	
	
	
	
}