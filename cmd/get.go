/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	writeCounter "github.com/Lakshamana/go-gopher-cli/interfaces"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get gopher",
	Short: "This command will get the given gopher",
	Long:  "The get command will call GitHub repository gopher image",
  Args: cobra.ExactArgs(1),
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
  gopherName := args[0]

	url := fmt.Sprintf("https://github.com/scraly/gophers/raw/main/%s.png", gopherName)
  fmt.Printf("Trying to get gopher %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading the image", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		out, err := os.Create(gopherName + ".png")
		if err != nil {
			fmt.Println(err)
		}

		defer out.Close()

    counter := writeCounter.NewWriteCounter()

		_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("\r> Saved file in " + out.Name() + strings.Repeat(" ", 35))
	} else {
		fmt.Printf("\r> Error: file %s doesn't exist", gopherName)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
