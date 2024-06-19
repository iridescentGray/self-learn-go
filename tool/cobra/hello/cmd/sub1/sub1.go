package sub1

import (
	"fmt"

	"github.com/spf13/cobra"
)

func doSub1() error {
	fmt.Println("sub1")
	return nil
}

var Sub1Cmd = &cobra.Command{
	Use:   "sub1",
	Short: "Here is sub1",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := doSub1(); err != nil {
			return err
		}
		return nil
	},
}
