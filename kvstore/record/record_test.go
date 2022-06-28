package record

import "testing"

/*
	Testing a combination of commands
*/
func Test_Combo_RecordOps(t *testing.T) {
	r := NewRecord()

	r.SetCommit("Commit")
	r.StartTrans(3, "value3")
	result := r.Get()
	if result != "value3" {
		t.FailNow()
	}
	r.Rollback(1)
	if result != "value3" {
		t.FailNow()
	}
	r.StartTrans(4, "value3")
	r.Rollback(3)
	result = r.Get()
	if result != "value3" {
		t.FailNow()
	}

	r.Unset()
	result = r.Get()
	if result != "" {
		t.FailNow()
	}
	r.Rollback(4)
	result = r.Get()
	if result != "Commit" {
		t.FailNow()
	}

	r.StartTrans(10, "value10")
	r.Commit()
	result = r.Get()
	if result != "value10" {
		t.FailNow()
	}

}
