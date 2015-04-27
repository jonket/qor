package admin

import (
	"testing"

	"github.com/qor/qor"
)

type User struct {
	Name string
}

func TestAddResource(t *testing.T) {
	admin := New(&qor.Config{})
	user := admin.AddResource(&User{})

	if user != admin.resources[0] {
		t.Error("resource not added")
	}

	if admin.GetMenus()[0].Name != "User" {
		t.Error("resource not added to menu")
	}
}

func TestAddResourceWithInvisibleOption(t *testing.T) {
	admin := New(&qor.Config{})
	user := admin.AddResource(&User{}, &Config{Invisible: true})

	if user != admin.resources[0] {
		t.Error("resource not added")
	}

	if len(admin.GetMenus()) != 0 {
		t.Error("invisible resource registered in menu")
	}
}
