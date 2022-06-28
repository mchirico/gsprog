package cmdpattern

import (
	"fmt"
	"reflect"
	"testing"
)

type PingCommand struct {
	args []string
}

func NewPing() *PingCommand {
	p := &PingCommand{}
	p.args = []string{}
	return p
}
func (p *PingCommand) verify(t *testing.T, args ...string) {
	if !reflect.DeepEqual(p.args, args) && !(len(p.args) == 0 && len(args) == 0) {
		t.Errorf("Expected: %s, Got: %s\n", p.args, args)
	}
	p.args = []string{}
}
func (p *PingCommand) Execute(args ...string) string {
	p.args = args
	return fmt.Sprintf("args: %d\n", len(args))
}

func Test_execByName(t *testing.T) {
	expected := "args: 3\n"

	p := NewPing()

	e := NewExe()
	e.Register("ping", p)

	// Verified called with correct args
	result := e.ExecByToken("ping", "one", "two", "three")
	if result != expected {
		t.Errorf("Expected: ->%s<-, Got: ->%s<-\n", expected, result)
	}
	p.verify(t, "one", "two", "three")

	// Not expecting command to be called
	result = e.ExecByToken("unknown", "one", "two", "three")
	p.verify(t)
	if result != "" {
		t.Errorf("Expected: ->%s<-, Got: ->%s<-\n", "", result)
	}
}
