package cmdpattern

type Command interface {
	Execute(...string) string
}

type Exe struct {
	cmd map[string]Command
}

func NewExe() *Exe {
	e := &Exe{}
	e.cmd = map[string]Command{}
	return e
}

func (e *Exe) Register(token string, cmd Command) *Exe {
	e.cmd[token] = cmd
	return e
}

func (e *Exe) ExecByToken(token string, args ...string) string {
	if command, ok := e.cmd[token]; ok {
		return command.Execute(args...)
	}
	return ""
}
