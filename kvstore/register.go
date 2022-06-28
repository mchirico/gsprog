package kvstore

import (
	"github.com/mchirico/gsprog/cmdpattern"
)

func SummaryOfCommands() string {
	summary := `
SET key value     
GET key
BEGIN 	       # This will begin transaction
ROLLBACK       # Rollback transaction
COMMIT         # Commit transaction
UNSET key      
NUMEQUALTO value   # Number of keys with value <value>
END                # End program
------------------------------------------------
	
	`
	return summary
}

func RegisterKVStoreCommands() *cmdpattern.Exe {
	kv := NewKVStore()
	e := cmdpattern.NewExe()
	e.Register("SET", NewSet(kv))
	e.Register("GET", NewGet(kv))
	e.Register("BEGIN", NewBegin(kv))
	e.Register("ROLLBACK", NewRoll(kv))
	e.Register("UNSET", NewUnSet(kv))
	e.Register("NUMEQUALTO", NewNumCommand(kv))
	e.Register("END", NewEndCommand(kv))
	e.Register("COMMIT", NewCommitCommand(kv))
	return e
}
