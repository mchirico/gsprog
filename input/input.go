package input

import (
	"bufio"
	"fmt"
	"github.com/mchirico/gsprog/cmdpattern"
	"github.com/mchirico/gsprog/kvstore"
	"os"
	"strings"
)

func ProcessInput() {

	e := kvstore.RegisterKVStoreCommands()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(kvstore.SummaryOfCommands())

	for {
		fmt.Print("-> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Printf("%s\n", exe(e, command))
	}

}

func exe(e *cmdpattern.Exe, lineInput string) string {
	tokens := strings.Fields(lineInput)
	switch {
	case len(tokens) <= 0:
		return ""
	case len(tokens) == 1:
		return e.ExecByToken(tokens[0])
	case len(tokens) > 1:
		return e.ExecByToken(tokens[0], tokens[1:]...)

	}
	return ""
}
