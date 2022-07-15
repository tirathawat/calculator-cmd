package executor

import (
	"fmt"
	"math/big"

	"pair/commander"

	"github.com/spf13/cobra"
)

type Executor struct {
	command cobra.Command
}

func (e *Executor) Execute() {
	e.command.Execute()
}

func (e *Executor) AddCommand(use, short string, fn commander.CommandFunc) {
	command := &cobra.Command{
		Use:   use,
		Short: short,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return ErrRequiredArguments
			}

			a, err := stringToBigFloat(args[0])
			if err != nil {
				return ErrInvalidArgument
			}

			b, err := stringToBigFloat(args[1])
			if err != nil {
				return ErrInvalidArgument
			}

			answer := fn(a, b)
			fmt.Println(answer.String())
			return nil
		},
	}
	e.command.AddCommand(command)
}

func ProvideExecutor() *Executor {
	return &Executor{}
}

func stringToBigFloat(value string) (*big.Float, error) {
	n := new(big.Float)
	n, ok := n.SetString(value)
	if !ok {
		return nil, ErrInvalidArgument
	}
	return n, nil
}
