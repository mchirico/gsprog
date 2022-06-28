package kvstore

import (
	"fmt"
	"os"
)

/*
	Commands are defined here

*/

// SET key value
type SetCommand struct {
	kv *KVStore
}

func NewSet(kv *KVStore) *SetCommand {
	p := &SetCommand{}
	p.kv = kv
	return p
}

func (p *SetCommand) Execute(args ...string) string {
	if len(args) < 2 {
		return "Two arguments required\n"
	}
	p.kv.Set(args[0], args[1])
	return ""

}

// UNSET key
type UnSetCommand struct {
	kv *KVStore
}

func NewUnSet(kv *KVStore) *UnSetCommand {
	p := &UnSetCommand{}
	p.kv = kv
	return p
}

func (p *UnSetCommand) Execute(args ...string) string {
	if len(args) != 1 {
		return "key value required\n"
	}
	p.kv.UnSet(args[0])
	return ""

}

// GET key
type GetCommand struct {
	kv *KVStore
}

func NewGet(kv *KVStore) *GetCommand {
	p := &GetCommand{}
	p.kv = kv
	return p
}

func (p *GetCommand) Execute(args ...string) string {
	if len(args) < 1 {
		return "Need key value\n"
	}
	return p.kv.Get(args[0])

}

// BEGIN
type Begin struct {
	kv *KVStore
}

func NewBegin(kv *KVStore) *Begin {
	p := &Begin{}
	p.kv = kv
	return p
}

func (p *Begin) Execute(args ...string) string {
	p.kv.Begin()
	return ""
}

// ROLLBACK       # Rollback transaction
type Roll struct {
	kv *KVStore
}

func NewRoll(kv *KVStore) *Roll {
	p := &Roll{}
	p.kv = kv
	return p
}

func (p *Roll) Execute(args ...string) string {
	return p.kv.Rollback()
}

// NUMEQUALTO value   # Number of keys with value <value>
type NumCommand struct {
	kv *KVStore
}

func NewNumCommand(kv *KVStore) *NumCommand {
	p := &NumCommand{}
	p.kv = kv
	return p
}

func (p *NumCommand) Execute(args ...string) string {
	if len(args) < 1 {
		return "Need key value\n"
	}
	count := 0
	if args[0] == "" {
		return "0"
	}
	for _, v := range p.kv.AllValues() {
		if v == args[0] {
			count += 1
		}

	}
	return fmt.Sprintf("%d", count)
}

// END                # End program
type EndCommand struct {
	kv *KVStore
}

func NewEndCommand(kv *KVStore) *EndCommand {
	p := &EndCommand{}
	p.kv = kv
	return p
}

func (p *EndCommand) Execute(args ...string) string {
	if len(args) > 0 {
		return "Did you want to end/exit? Try END with no arguments"
	}
	os.Exit(0)
	return ""
}

// COMMIT         # Commit transaction
type CommitCommand struct {
	kv *KVStore
}

func NewCommitCommand(kv *KVStore) *CommitCommand {
	p := &CommitCommand{}
	p.kv = kv
	return p
}

func (p *CommitCommand) Execute(args ...string) string {
	return p.kv.Commit()

}
