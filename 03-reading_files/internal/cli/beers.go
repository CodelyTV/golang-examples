package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(),
	}

	return beersCmd
}

func runBeersFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {

		f, _ := os.Open("03-reading_files/data/beers.csv")
		reader := bufio.NewReader(f)
		for line := readLine(reader); line != nil; line = readLine(reader) {
			fmt.Printf("line (as bytes): %v\n", line)
			fmt.Printf("line (as string): %v\n", string(line))
		}

		fmt.Printf("\n\nEnd of file\n")

	}
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
