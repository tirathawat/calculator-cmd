package main

import (
	"pair/calculator"
	"pair/commander"
	"pair/executor"
)

func main() {
	comm := commander.ProvideCommander(calculator.Calculator{})
	exec := executor.ProvideExecutor()
	Run(comm, exec)
}

func Run(comm *commander.Commander, exec *executor.Executor) {
	exec.AddCommand(
		"add",
		"this is add",
		comm.Add(),
	)

	exec.AddCommand(
		"minus",
		"this is minus",
		comm.Minus(),
	)

	exec.AddCommand(
		"multiply",
		"this is multiply",
		comm.Multiply(),
	)

	exec.AddCommand(
		"divide",
		"this is divide",
		comm.Divide(),
	)

	exec.Execute()
}
