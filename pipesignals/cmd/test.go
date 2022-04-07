/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Used to test pipes",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("test called")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")s

	print("Hello from Test!\n")

	if isInputFromPipe() {
		print("> Input is from pipe\n")
		// This is where we will call the processing for the pipe
		inspectPipe()
		print(">> Goodby from pipeutil\n")
	} else {
		print("> Input is not from a pipe\n")
	}

}

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}

func inspectPipe() {

	pr, pw := io.Pipe()

	go func() {
		//fmt.Fprint(pw, "some io.Reader stream to be read\n")
		reader := bufio.NewReader(os.Stdin)
		//io.Copy(pw, reader)

		for {
			_, err := io.CopyN(pw, reader, 1024*100)

			if err != nil {
				break
			}
			time.Sleep(5 * time.Second)

		}

		pw.Close()
	}()

	if _, err := io.Copy(os.Stdout, pr); err != nil {
		panic(err)
	}

}
