package cmd

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"github.com/kyokomi/emoji/v2"
	"io/ioutil"
	"net/http"
	"os"
)

const ApiEndpoint = "https://user-api.coronatest.nl"
const UrlTemplate = "vaccinatie/programma/bepaalbaar/%s/NEE/NEE"

type ApiResponse struct {
	success string
}
var rootCmd = &cobra.Command{
	Use:   "canihasvaccine",
	Short: "Lookup if your year is allowed to register for a vaccine",
	Long: `A simple API reader to parse the output of coronatest.nl API server`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := http.Get(fmt.Sprintf("%s/%s", ApiEndpoint, fmt.Sprintf(UrlTemplate, args[0])))
		if err != nil {
			fmt.Printf("ERROR: API server returned an error: %s", err.Error())
			os.Exit(1)
		}

		respData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ERROR: could not read response from API server, error: %v", err)
			os.Exit(1)
		}

		//var data ApiResponse

		jsonObj, err := gabs.ParseJSON([]byte(respData))
		if err != nil {
			fmt.Printf("ERROR: data from API server is corrupted")
			os.Exit(1)
		}

		value := jsonObj.Path("success").String()
		if value == "true" {
			fmt.Println(emoji.Sprint(":syringe: Yes you can! HOORAY!!! :syringe:"))
		} else {
			fmt.Println(emoji.Sprint(":mask: Not yet. try again tomorrow... :mask:"))
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}