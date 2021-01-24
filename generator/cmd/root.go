package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/conradludgate/tfutils/generator/pkg"
)

func FooSchemaMap(w io.Writer) {
	s := map[string]interface{}{}
	s["StructName"] = "Foo"

	fields := []interface{}{}
	fields = append(fields, map[string]interface{}{
		"Template":    "simple",
		"Name":        "A",
		"Type":        "String",
		"Default":     strconv.Quote("foo"),
		"Description": strconv.Quote("Some value"),
	})
	fields = append(fields, map[string]interface{}{
		"Template":    "simple",
		"Name":        "B",
		"Type":        "Int",
		"Required":    true,
		"Description": strconv.Quote("Some other value"),
	})
	fields = append(fields, map[string]interface{}{
		"Template":    "simple_list",
		"Name":        "List",
		"Elem":        "String",
		"Description": strconv.Quote("Simple lists can be implemented"),
	})
	fields = append(fields, map[string]interface{}{
		"Template":     "complex_list",
		"Name":         "AnotherList",
		"ElemTypeName": "Baz",
		"Description":  strconv.Quote("Complex lists can too, as long as Baz\nimplements the 'Schema' interface"),
	})
	fields = append(fields, map[string]interface{}{
		"Template":    "map",
		"Name":        "Map",
		"Elem":        "Int",
		"Description": strconv.Quote("Maps can only be over simple types (terraform limitation)"),
	})
	fields = append(fields, map[string]interface{}{
		"Template":     "complex_set",
		"Name":         "Set",
		"ElemTypeName": "Bar",
		"Description":  strconv.Quote("map[int]... represents a Set.\nIf Bar implements the `Set` interface,\nthen that will be the Set function"),
	})

	s["Fields"] = fields

	pkg.IntoSchemaMapTemplate.Execute(w, s)
}

func FooUnmarshal(w io.Writer) {
	s := map[string]interface{}{}
	s["StructName"] = "Foo"

	fields := []interface{}{}
	fields = append(fields, map[string]interface{}{
		"Template": "simple",
		"Name":     "A",
		"Type":     "string",
	})
	fields = append(fields, map[string]interface{}{
		"Template": "simple",
		"Name":     "B",
		"Type":     "int",
	})
	fields = append(fields, map[string]interface{}{
		"Template": "simple_list",
		"Name":     "List",
		"ElemType": "string",
	})
	fields = append(fields, map[string]interface{}{
		"Template": "complex_list",
		"Name":     "AnotherList",
		"ElemType": "Baz",
	})
	fields = append(fields, map[string]interface{}{
		"Template": "map",
		"Name":     "Map",
		"ElemType": "int",
	})
	fields = append(fields, map[string]interface{}{
		"Template": "complex_set",
		"Name":     "Set",
		"ElemType": "Bar",
	})

	s["Fields"] = fields

	pkg.UnmarshalResourceTemplate.Execute(w, s)
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tfutils",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		FooUnmarshal(cmd.OutOrStdout())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tfutils.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".tfutils" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tfutils")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
