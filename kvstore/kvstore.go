package kvstore

import (
	"fmt"
	"github.com/mchirico/gsprog/constants"
	"github.com/mchirico/gsprog/kvstore/record"
	"log"
)

type Record interface {
	Unset()
	SetCommit(value string)
	Commit()
	StartTrans(key int, value string)
	Rollback(rollbackey int) string
	Get() string
}

type KVStore struct {
	rec   map[string]Record
	stack []int
}

func NewKVStore() *KVStore {
	kv := &KVStore{}
	kv.rec = map[string]Record{}
	kv.stack = []int{}
	return kv

}

func (kv *KVStore) sizeCheck() error {
	if len(kv.rec) >= constants.KVMAP_WARN {
		log.Printf("kv.rec very large: %d", len(kv.rec))
		if len(kv.rec) >= constants.KVMAP_MAX {
			log.Printf("kv.rec too large: %d", len(kv.rec))
			return fmt.Errorf("kv.rec too large")
		}
	}
	return nil
}

func (kv *KVStore) Set(key, value string) {
	if err := kv.sizeCheck(); err != nil {
		return
	}
	r := kv.getRec(key)
	if tranNum, ok := kv.inTransaction(); ok {
		r.StartTrans(tranNum, value)
		return
	}
	r.SetCommit(value)
}

func (kv *KVStore) UnSet(key string) {
	value := ""
	if err := kv.sizeCheck(); err != nil {
		return
	}
	r := kv.getRec(key)
	if tranNum, ok := kv.inTransaction(); ok {
		r.StartTrans(tranNum, value)
		return
	}
	r.SetCommit(value)
}

func (kv *KVStore) AllValues() map[string]string {
	out := map[string]string{}
	for k, v := range kv.rec {
		out[k] = v.Get()
	}
	return out

}

func (kv *KVStore) Get(key string) string {
	if r, ok := kv.rec[key]; ok {
		if r.Get() == "" {
			return "Nil"
		}
		return r.Get()
	}
	return "Nil"
}

func (kv *KVStore) Commit() string {
	if _, ok := kv.inTransaction(); ok {
		for _, v := range kv.rec {
			v.Commit()
		}
		kv.stack = []int{}
		return ""
	}
	return "NO TRANSACTION"
}

func (kv *KVStore) Rollback() string {
	if tranNum, ok := kv.inTransaction(); ok {
		for _, v := range kv.rec {
			v.Rollback(tranNum)
		}
		kv.stack = kv.stack[0 : tranNum-1]
		return ""
	}
	return "NO TRANSACTION"
}

func (kv *KVStore) Begin() {
	if len(kv.stack) >= constants.KVSTACK_MAX {
		log.Printf("warning: kv.stack too large: %d\nUpdate KVSTACK_MAX", len(kv.stack))
		return
	}
	kv.stack = append(kv.stack, len(kv.stack))
}

func (kv *KVStore) inTransaction() (int, bool) {
	length := len(kv.stack)
	if length > 0 {
		return length, true
	}
	return 0, false
}

func (kv *KVStore) getRec(key string) Record {
	if r, ok := kv.rec[key]; ok {
		return r
	}
	r := record.NewRecord()
	kv.rec[key] = r
	return r
}
