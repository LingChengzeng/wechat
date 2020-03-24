// exec
package tools

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// PathSeparator is the string of os.PathSeparator
var PathSeparator = string(os.PathSeparator)

// Cmd is a custom struch which 'implements' the *exec.Cmd
type Cmd struct {
	*exec.Cmd
}

// Arguments sets the command line arguments, including the command as Args[0]
// If the args parameter is empty or nil, Run uses {Path}.
// In typical use, both Path and args are set by calling Command.
func (this *Cmd) Arguments(args ...string) *Cmd {
	// we need the first argument which is the command.
	this.Cmd.Args = append(this.Cmd.Args[0:1], args...)
	return this
}

// AppendArguments appends the arguments to the exists.
func (this *Cmd) AppendArguments(args ...string) *Cmd {
	this.Cmd.Args = append(this.Cmd.Args, args...)
	return this
}

// ResetArguments resets the arguments.
func (this *Cmd) ResetArguments() *Cmd {
	// keep only the first because is the command.
	this.Args = this.Args[0:1]
	return this
}

// Directory sets the working directory of the command.
// If workingDirectory is the empty string, Run runs the command in the
// calling process's current directory.
func (this *Cmd) Directory(workingDirectory string) *Cmd {
	this.Cmd.Dir = workingDirectory
	return this
}

// CommandBuilder creates a Cmd object and returns it
// accepts 2 parameters, one is optionally
// first parameter is the command (string)
// second variatic parameter is the argument(s) (slice of string)
//
// the difference from the normal Command function is that you can re-use
// this Cmd, it doesn't execute until you  call its Command function
func CommandBuilder(command string, args ...string) *Cmd {
	return &Cmd{Cmd: exec.Command(command, args...)}
}

// the below is just for exec.Command
// Command executes a command in shell and returns it's ouput,
// it's block versioin.
func Command(command string, args ...string) (output string, err error) {
	var out []byte

	// if no args given, try to get them from the command.
	if len(args) == 0 {
		commandArgs := strings.Split(command, " ")
		for _, commandArg := range commandArgs {
			// if starts with - means that this is and argument, append it to
			// the arguments.
			if commandArg[0] == '-' {
				args = append(args, commandArg)
			}
		}
	}

	out, err = exec.Command(command, args...).Output()
	if err != nil {
		output = string(out)
	}

	return
}

// MustCommand executes a command in shell and returns it's output,
// it's block version. It panics on an error
func MustCommand(command string, args ...string) (output string) {
	var out []byte
	var err error
	if len(args) == 0 {
		commandArgs := strings.Split(command, " ")
		for _, commandArg := range commandArgs {
			// if starts with - means that this is an argument,
			// append it to the arguments
			if commandArg[0] == '-' {
				args = append(args, commandArg)
			}
		}
	}

	out, err = exec.Command(command, args...).Output()
	if err != nil {
		argsToString := strings.Join(args, " ")
		panic(fmt.Sprintf("\nError running the command %s", command+
			" "+argsToString))
	}

	output = string(out)

	return
}
