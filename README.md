[![Go](https://github.com/mchirico/gsprog/actions/workflows/go.yml/badge.svg)](https://github.com/mchirico/gsprog/actions/workflows/go.yml)
# github.com/mchirico/gsprog

## Running Program

### Step 1

Clone repo, or run from codespace
```bash
git clone https://github.com/mchirico/gsprog.git
cd gsprog
go mod tidy
go run main.go

```

A menu similiar to the following, will prompt for input 


```bash
SET key value     
GET key
BEGIN          # This will begin transaction
ROLLBACK       # Rollback transaction
COMMIT         # Commit transaction
UNSET key      
NUMEQUALTO value   # Number of keys with value <value>
END                # End program
------------------------------------------------


-> 

```

## Step 2

Example entering input

```bash

-> SET key1 1

-> SET key2 1

-> NUMEQUALTO 1
2

```

## Step 3

To end the program..."END"

```bash
-> END
```


# Comprehensive Program Testing

Although the program can be run interactively,
it can also be driven from a program.

Reference: `e2e/e2e_test.go`

```go

func Test_Basic_Ops(t *testing.T) {
	e := kvstore.RegisterKVStoreCommands()
	e.ExecByToken("SET", "key0", "value0")
	checkResult(t, e.ExecByToken("GET", "key0"), "value0")
	e.ExecByToken("UNSET", "key0")
	checkResult(t, e.ExecByToken("GET", "key0"), "Nil")

}


```

