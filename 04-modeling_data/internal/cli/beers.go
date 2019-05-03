package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	beerscli "github.com/CodelyTV/golang-introduction/04-modeling_data/internal"
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

		var beers = make(map[int]beerscli.Beer)

		for line := readLine(reader); line != nil; line = readLine(reader) {
			values := strings.Split(string(line), ",")

			productID, _ := strconv.Atoi(values[0])
			price, _ := strconv.ParseFloat(values[3], 64)

			beer := beerscli.Beer{
				ProductID: productID,
				Name:      values[1],
				Category:  values[2],
				Price:     price,
				Type:      values[4],
				Brewer:    values[5],
				Country:   values[6],
			}

			beers[productID] = beer
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
