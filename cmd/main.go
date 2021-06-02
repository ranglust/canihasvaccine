package cmd

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/kyokomi/emoji/v2"
	"github.com/ranglust/canihasvaccine/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"io/ioutil"
	"net/http"
	"os"
)

const ApiEndpoint = "https://user-api.coronatest.nl"
const UrlTemplate = "vaccinatie/programma/bepaalbaar/%s/NEE/NEE"

var Config = config.Configuration{}

type ApiResponse struct {
	success string
}
var rootCmd = &cobra.Command{
	Use:   "canihasvaccine",
	Short: "Lookup if your year is allowed to register for a vaccine",
	Long: `A simple API reader to parse the output of coronatest.nl API server`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		yearFlag, err := cmd.Flags().GetString("year")
		if err != nil {
			fmt.Printf("ERROR: could not parse command line flags, error: %v\n", err)
			os.Exit(1)
		}

		var year string
		if yearFlag != "" {
			year = yearFlag
		} else if Config.Year != "" {
			year = Config.Year
		} else {
			fmt.Println("ERROR: you must specify the year with the --year command line flag")
			os.Exit(1)
		}

		if canIHasVaccine(year) {
			fmt.Println(emoji.Sprint(":syringe: Yes you can! HOORAY!!! :syringe:"))
		} else {
			fmt.Println(emoji.Sprint(":mask: Not yet. try again tomorrow... :mask:"))
		}
	},
}

func canIHasVaccine(year string) bool {
	resp, err := http.Get(fmt.Sprintf("%s/%s", ApiEndpoint, fmt.Sprintf(UrlTemplate, year)))
	if err != nil {
		fmt.Printf("ERROR: API server returned an error: %s", err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: could not read response from API server, error: %v", err)
		os.Exit(1)
	}

	jsonObj, err := gabs.ParseJSON([]byte(respData))
	if err != nil {
		fmt.Printf("ERROR: data from API server is corrupted")
		os.Exit(1)
	}

	value := jsonObj.Path("success").String()

	return value == "true"
}

func init() {
	initViper()
	viperReadConfig(&Config)

	rootCmd.Flags().StringP("year", "y", "", "Year (#### format, i.e. 1999)")
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func viperReadConfig(configuration *config.Configuration) {
	// read config
	// Ignore silently since it may raise an error when config file
	// is missing
	_ = viper.ReadInConfig()

	// load config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("ERROR: could not load config, %s\n", err)
		os.Exit(-1)
	}
}

func initViper() {
	viper.SetConfigName("canihasvaccine")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.canihasvaccine")
	viper.AddConfigPath(".")
}
