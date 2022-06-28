package kvstore

import "testing"

func TestBucket_Set(t *testing.T) {
	kv := NewKVStore()
	kv.Set("bozo", "1")
	kv.Begin()
	kv.Set("zu", "2")
	kv.Set("bozo", "2")
	kv.Rollback()
	kv.Rollback()
	kv.Rollback()
	kv.AllValues()
	t.Logf("-- %s\n", kv.Get("bozo"))
	kv.Begin()
	kv.Set("beg1", "1")
	kv.Rollback()
	kv.Begin()
	kv.Set("zu", "2")
	kv.Set("bozo", "2")
	sresult := kv.Get("bozo")
	if sresult != "2" {
		t.FailNow()
	}
	kv.Begin()
	kv.Set("zu", "29")
	kv.Commit()
	result := kv.AllValues()
	t.Logf("\n\n...........\n")
	for k, v := range result {
		t.Logf("%s,%s\n", k, v)
	}
	if result["zu"] != "29" {
		t.FailNow()
	}

}
