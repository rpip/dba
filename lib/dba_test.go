package dba

import "testing"

func TestMeta(t *testing.T) {
	table := NewTable("users")
	table.SetMeta("primary_key", "user_id")
	if table.MustGetMeta("primary_key") != "user_id" {
		t.Error("expected primary_key meta field to be user_id")
	}

	if table.meta.Name != "users" {
		t.Error("expected table name to be users")
	}
}
