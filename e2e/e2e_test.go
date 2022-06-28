package e2e

import (
	"fmt"
	"github.com/mchirico/gsprog/kvstore"
	"testing"
)

func checkResult(t *testing.T, result, expected string) {
	if expected != result {
		t.Fatalf("Got: ->%s<-  Wanted: ->%s<-\n", result, expected)
	}
}

func Test_Transaction_Simple(t *testing.T) {
	e := kvstore.RegisterKVStoreCommands()

	// Start with transaction
	e.ExecByToken("BEGIN")
	for i := 1; i <= 10; i++ {
		key := fmt.Sprintf("key%d", i)
		e.ExecByToken("SET", key, "count")
	}
	checkResult(t, e.ExecByToken("NUMEQUALTO", "count"), "10")

	// Next level
	e.ExecByToken("BEGIN")
	for i := 11; i <= 20; i++ {
		value := fmt.Sprintf("value_%d", i)
		e.ExecByToken("SET", "key0", value)

		key := fmt.Sprintf("key%d", i)
		e.ExecByToken("SET", key, "count")
	}
	checkResult(t, e.ExecByToken("NUMEQUALTO", "count"), "20")

	// Now start rolling back
	e.ExecByToken("ROLLBACK")
	checkResult(t, e.ExecByToken("NUMEQUALTO", "count"), "10")

	e.ExecByToken("ROLLBACK")
	checkResult(t, e.ExecByToken("NUMEQUALTO", "count"), "0")

	checkResult(t, e.ExecByToken("ROLLBACK"), "NO TRANSACTION")

}

func Test_Trasaction_with_Unset(t *testing.T) {
	e := kvstore.RegisterKVStoreCommands()
	for i := 1; i <= 100; i++ {
		key := fmt.Sprintf("key%d", i)
		e.ExecByToken("SET", key, "count")

		// Half way set a transaction
		if i == 50 {
			e.ExecByToken("BEGIN")
		}
	}
	checkResult(t, e.ExecByToken("NUMEQUALTO", "count"), "100")

	e.ExecByToken("BEGIN")
	for i := 51; i <= 100; i++ {
		key := fmt.Sprintf("key%d", i)
		e.ExecByToken("UNSET", key)
	}
	checkResult(t, e.ExecByToken("NUMEQUALTO", "count"), "50")
	e.ExecByToken("ROLLBACK")
	checkResult(t, e.ExecByToken("NUMEQUALTO", "count"), "100")

	checkResult(t, e.ExecByToken("COMMIT"), "")
	checkResult(t, e.ExecByToken("COMMIT"), "NO TRANSACTION")
	checkResult(t, e.ExecByToken("COMMIT"), "NO TRANSACTION")

}
