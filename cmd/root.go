package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "goconsole",
	Short: "Goconsole is a project to simmulate some of the basic linux commands in Go",
	Long: `		A golang implementation of some of the basic linux commands 
		to interact with files and directories . I built this application to practice Go while
		I am learning about go libraries.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go Console")
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", "Ahmed Abouzied", "copyrights to Ahmed Abouzied")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Ahmed Abouzied <ahmedaabouzied44@gmail.com>")
	viper.SetDefault("license", "apache")
}

func initConfig() {

	// Search config in home directory with name ".cobra" (without extension).
	viper.AddConfigPath("./")
	viper.SetConfigName(".cobra")
	// }

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
