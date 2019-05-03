package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {

		f, _ := os.Open("03-reading_files/data/beers.csv")
		reader := bufio.NewReader(f)

		var beers = make(map[int]string)

		for line := readLine(reader); line != nil; line = readLine(reader) {
			values := strings.Split(string(line), ",")

			productID, _ := strconv.Atoi(values[0])
			beers[productID] = values[1]
		}

		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			fmt.Println(beers[i])
		} else {
			fmt.Println(beers)
		}
	}
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
