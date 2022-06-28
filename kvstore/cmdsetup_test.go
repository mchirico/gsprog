package kvstore

import (
	"testing"
)

func TestRegisterKVStoreCommands(t *testing.T) {
	e := RegisterKVStoreCommands()
	e.ExecByToken("SET", "key0", "value_3")
	e.ExecByToken("SET", "key1", "value_3")
	e.ExecByToken("SET", "key2", "value_3")
	result := e.ExecByToken("NUMEQUALTO", "value_3")
	if result != "3" {
		t.Errorf("Expected %s, got %s\n", "3", result)
	}
	// TRANS
	e.ExecByToken("BEGIN")
	e.ExecByToken("SET", "key4", "value_3")
	e.ExecByToken("SET", "key5", "value_3")
	result = e.ExecByToken("NUMEQUALTO", "value_3")
	if result != "5" {
		t.Errorf("Expected %s, got %s\n", "5", result)
	}
	e.ExecByToken("UNSET", "key5")
	if result := e.ExecByToken("GET", "key5"); result != "Nil" {
		t.Errorf("Expected Nil. Gog %s\n", result)
	}

	if result := e.ExecByToken("GET", "unknown"); result != "Nil" {
		t.Errorf("Expected Nil. Gog %s\n", result)
	}

	// ROLLBACK
	if result := e.ExecByToken("ROLLBACK"); result != "" {
		t.Errorf("Rollback work, should have returned empty string: %s\n", result)
	}
	result = e.ExecByToken("NUMEQUALTO", "value_3")
	if result != "3" {
		t.Errorf("Expected %s, got %s\n", "3", result)
	}

	if result := e.ExecByToken("ROLLBACK"); result != "NO TRANSACTION" {
		t.Errorf("Rollback work, should have returned empty string: %s\n", result)
	}
}
