package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/slipperystairs/killahtask/cowsay"
	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var cow bool

func Now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

var addCommand = &cobra.Command{
	Use:     "add",
	Short:   "Adds a new item",
	Aliases: []string{"a"},
	Long:    `Adds a task to your list of TODOs by creating or updating your ~/killahtask_<username>.csv file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		earlyExit := false
		if len(args) == 0 {
			command = "add"
			PrintMsg(&command, "missing_task")
			earlyExit = true
		}

		if earlyExit {
			os.Exit(1)
		}

		description := strings.TrimSpace(strings.Join(args, " "))
		successMsg := fmt.Sprintf("\"%s\" added successfully!", description)
		// os.OpenFile doesn't have a way of letting us know if the file already exist.
		fileInfo, err := os.Stat(CurrentUser.Filepath)
		fileExists := err == nil

		// We check the file size because I ran into a bug during development where my file existed but it was empty.
		// Treat it like a new file when the user falls into this case for whatever reason.
		hasData := fileExists && fileInfo.Size() > 0
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		file, err := task.LoadFile(CurrentUser.Filepath)
		if err != nil {
			return err
		}
		// File will get closed even in the event of an error.
		defer task.CloseFile(file)

		if !fileExists || !hasData {
			records := [][]string{
				{"task_id", "description", "created", "completed"},
				{"0", description, Now(), "false"},
			}

			err := task.WriteCSV(file, records)
			if err != nil {
				return err
			}
		} else {
			newId := "0"
			csvReader := csv.NewReader(file)
			records, err := csvReader.ReadAll()
			if err != nil {
				return err
			}

			// It is possible for the user to delete the only record and be left with just the headers
			// If that is the case we reset the ID to 0.
			if len(records) > 1 {
				// Get the last task_id used and increment it by one.
				lastId, err := strconv.Atoi(records[len(records)-1][0])
				if err != nil {
					return err
				}
				newId = strconv.Itoa(lastId + 1)
			}

			if !uniqueDescription(description, records) {
				errStr := "Task description isn't unique!\n\"" + description + "\" already exists.\n"
				return errors.New(errStr)
			}

			// Append the new record to the end of the slice.
			records = append(records, []string{newId, description, Now(), "false"})
			csvErr := task.WriteCSV(file, records) // Re-write the file with the new records
			if csvErr != nil {
				return csvErr
			}
		}

		if !cow {
			fmt.Printf("%s\n", successMsg)
		} else {
			lines := []string{successMsg}
			cowsay.CowSay(lines)
		}
		return nil
	},
}

func init() {
	addCommand.PersistentFlags().BoolVar(&cow, "cowsay", false, "Display output using Cowsay")
	rootCmd.AddCommand(addCommand)
}
