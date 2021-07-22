package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)
var (
	_ *os.FileInfo
	_ error
)

var (
	dir     string
	file    string
	rootCmd = &cobra.Command{
		Use:   "del",
		Short: "del is a fast deleting application",
		Long: `A fast command line application that can delete
	files and folders with simple steps, del is built with love by Kevin.`,
		Run: func(cmd *cobra.Command, args []string) {
			if dir != "" {
				delDir(dir)
			} else if file != "" {
				delFile(file)
			} else {
				fmt.Println("parse in a flag, Example: \n[--d to delete directories and --f to delete files]")
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func main() {
	err := Execute()
	if err != nil {
		return
	}
}

func delDir(dir string) {

	if _, err := os.Stat(dir); err != nil {
		fmt.Printf("'%s' does not exist\n", dir)
	} else {
		fmt.Printf("successfully deleted '%s'\n", dir)
	}
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
}
func delFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Printf("'%s' does not exist\n", file)
	} else {
		fmt.Printf("Deleted '%s' successfully\n", file)
	}
}
func init() {
	rootCmd.PersistentFlags().StringVar(&dir, "d", "", "force delete a directory with it's contents.")
	rootCmd.PersistentFlags().StringVar(&file, "f", "", "force delete a file.")
}