package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go_8_mage/week14/vblog/api/version"
	"os"
)

var (
	vers bool
)
// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "vblog",
	Short:"微博客系统后台API Server",
	Long:"微博客系统后台API Server" ,
	RunE: func(cmd *cobra.Command,args []string)error{
		if vers{
			fmt.Println(version.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(){
	if err:=RootCmd.Execute();err!=nil{
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init(){
	RootCmd.PersistentFlags().BoolVarP(&vers,"version","v",false,"the vblog api version")
}