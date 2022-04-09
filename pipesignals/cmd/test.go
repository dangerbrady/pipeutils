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
		main()
	},
}

var readBufferSize int
var readBufferDelay int
var ctlFile string
var ctlFileContents string

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	testCmd.Flags().IntVarP(&readBufferSize, "read-buffer-size", "r", 1024, "Size of CopyN in Bytes")
	testCmd.Flags().IntVarP(&readBufferDelay, "read-buffer-delay", "d", 0, "Length of delay between reads in seconds")
	testCmd.Flags().StringVarP(&ctlFile, "control-file", "f", "./pipe.ctl", "File for controling pipe")

	print("Hello from Test Init!\n")

}

func readCtlFile() {

	for {
		dat, err := os.ReadFile(ctlFile)
		if err != nil {
			panic(err)
		}
		ctlFileContents = string(dat)
		time.Sleep(5 * time.Second)

	}
}

func main() {
	if isInputFromPipe() {
		print("> Input is from pipe\n")
		// This is where we will call the processing for the pipe
		delayPipe(readBufferSize, readBufferDelay)
		print(">> Goodby from pipeutil\n")
	} else {
		print("> Input is not from a pipe\n")
		print("\n")
		print("main says ", readBufferSize)
	}

}

func printFlags() {
	var rbs int
	rbs = readBufferSize
	print("\nPrintFlags says >>> ", rbs)
}

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}

func delayPipe(readBufferSize int, readBufferDelay int) {

	go readCtlFile()

	pr, pw := io.Pipe()

	go func() {
		//fmt.Fprint(pw, "some io.Reader stream to be read\n")
		reader := bufio.NewReader(os.Stdin)
		//io.Copy(pw, reader)

		for {
			_, err := io.CopyN(pw, reader, int64(readBufferSize))

			if err != nil {
				break
			}
			if ctlFileContents == "1" {
				time.Sleep(time.Duration(readBufferDelay) * time.Second)
			}
		}

		pw.Close()
	}()

	if _, err := io.Copy(os.Stdout, pr); err != nil {
		panic(err)
	}

}
