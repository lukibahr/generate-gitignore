package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:          "generate-gitignore",
	Short:        "generate-gitignore - client helper tool to generate a .gitignore file",
	Long:         `generate-gitignore - client helper tool to generate a .gitignore file. use the shorthand items avialable at toptal as parameter arguments, comma separated. Example -i go -i visualstudiocode,python,go`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		content, err := generateGitignore(toptal, items)
		if err != nil {
			log.Fatalf("error getting contents for file %s", err.Error())
		}
		// nolint:staticcheck
		err2 := ioutil.WriteFile(*&gitignorefile, content, 0755)
		if err2 != nil {
			log.Fatal(err2)
		}
		// nolint:staticcheck
		log.Infof("successfully generated file %s with items %s", *&gitignorefile, *&items)

		return nil
	},
}

var toptal, items, gitignorefile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&toptal, "toptal", "t", "https://www.toptal.com/developers/gitignore/api/", "the url to toptal")
	rootCmd.PersistentFlags().StringVarP(&items, "items", "i", "go", "items to add to .gitignore. defaults to go as go is lit")
	rootCmd.PersistentFlags().StringVarP(&gitignorefile, "gitignorefile", "g", ".gitignore", "the output .gitignore file to write to. defaults to .gitignore")
}

func generateGitignore(toptal, params string) ([]byte, error) {

	url := fmt.Sprintf("%s%s", toptal, params)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyText, nil
}

// Execute executes the rootCmd
func Execute(version string) {

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
