package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

const WHICH_OS string = runtime.GOOS

// todo => We need to learn how to read everything in between double quotes i.e., "description"
// todo => Go get the encoding/csv package to learn how to write the CSV file
// todo =>   - Two items cannot have the same description
// todo =>   - Locking the file for writes, to avoid race conditions
// todo => Figure out how we are going to handle the id's (prob have to read the last item in the file)
// todo => Things to consider
// todo =>   - Do we need to worry about the operating system using the CLI tool?
var addCommand = &cobra.Command{
	Use:     "add",
	Short:   "Adds a new item",
	Aliases: []string{"a"},
	Long:    `This command will add a item to your list`,
	Run: func(cmd *cobra.Command, args []string) {
		currUsr, err := user.Current()
		checkError(err)
		fmt.Printf("Curr user: %s\n", currUsr)

		if len(args) == 0 {
			PrintUsageMsg("add", "add_none")
		} else if len(args) > 1 {
			PrintUsageMsg("add", "add_to_many")
		}

		var fileExist bool = true
		var fileName string = "killahtask_" + currUsr.Username + ".csv"
		var filePath string = filepath.Join(currUsr.HomeDir, fileName)
		_, err = os.Stat(filePath)
		if err != nil {
			fileExist = false
		}

		if fileExist {
			fmt.Println("This is where we would get the last id")
		} else {
			file, err := os.Create(filePath)
			checkError(err)
			defer file.Close();

			records := [][]string{
				{"task_id", "description", "created", "completed"},
				{"0", args[0], "a few seconds ago", "false"},
			}
			w := csv.NewWriter(file)
			w.WriteAll(records)
			checkError(w.Error())
		}
		fmt.Println("Record added successfully. Run \"killahtask list\" to see your task.")
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
