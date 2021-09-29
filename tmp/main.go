package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/cobra/cmd"
	"os"
)

func init() {
	cobra.
}

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var versionCmd = cobra.Command{
		Use:                    "version",
		Aliases:                nil,
		SuggestFor:             nil,
		Short:                  "print version",
		Long:                   "all software version",
		Example:                "",
		ValidArgs:              nil,
		ValidArgsFunction:      nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: "",
		Deprecated:             "",
		Annotations:            nil,
		Version:                "",
		PersistentPreRun:       nil,
		PersistentPreRunE:      nil,
		PreRun:                 nil,
		PreRunE:                nil,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version is 1.0.0")
		},
		RunE:                       nil,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
		CompletionOptions:          cobra.CompletionOptions{},
		TraverseChildren:           false,
		Hidden:                     false,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
	}




	////r := gin.Default()
	////r.GET("/")
	//set_data(20)
	//
	//fmt.Println("everything is ok")

}

//func set_data(x int)  {
//	defer func() {
//		if err:=recover(); err != nil {
//			fmt.Println(err)
//
//		}
//	}()
//	var arr [10]int
//	arr[x]=88
//	flag.Parse()
//
//
//}
