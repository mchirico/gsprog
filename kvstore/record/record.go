package record

import (
	"github.com/mchirico/gsprog/constants"
	"log"
)

/*
	Record is a "record" of values

	The "key" part is done with a map in kvstore. Here,
	we're only working with the value part

*/

type record struct {
	commit   string
	transMap map[int]string
	active   string
	stack    []int
}

func NewRecord() *record {
	r := &record{}
	r.transMap = map[int]string{}
	r.stack = []int{}
	return r
}

func (r *record) Unset() {
	if len(r.stack) > 0 {
		r.active = ""
	} else {
		r.commit = ""
	}

}

func (r *record) SetCommit(value string) {
	r.commit = value
}

func (r *record) Commit() {
	if len(r.stack) > 0 {
		r.commit = r.active
		r.active = ""
		r.stack = []int{}
		r.transMap = map[int]string{}

	}

}

// StartTrans ... start a transation. Kicked off with BEGIN
func (r *record) StartTrans(key int, value string) {
	r.transMap[key] = value
	r.active = value
	if len(r.stack) >= constants.RSTACK_MAX {
		log.Printf("warning: r.stack too large: %d value: %s\nUpdate RSTACK_MAX", len(r.stack), value)
		return
	}
	r.stack = append(r.stack, key)
}

// Rollback transaction with key, if it exists
// if it doesn't exist, do nothing
func (r *record) Rollback(key int) string {
	if _, ok := r.transMap[key]; ok {
		delete(r.transMap, key)
		r.stack = remove(r.stack, key)
		if len(r.stack) > 0 {
			r.active = r.transMap[r.stack[len(r.stack)-1]]
		} else {
			r.active = ""
		}
		return ""
	}
	return "NO TRANSACTION"
}

// Get returns the active record, or if no
// active record, then, return a commited record.
func (r *record) Get() string {
	if len(r.stack) > 0 {
		return r.active
	}
	return r.commit
}

// remote is a helper function
func remove(slice []int, s int) []int {
	newSlice := []int{}
	for _, v := range slice {
		if v != s {
			newSlice = append(newSlice, v)
		}

	}
	return newSlice
}
