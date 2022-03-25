package cmd

import (
	"os"
	"reflect"
	"testing"

	"log"
)

func read(filename string) ([]byte, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func Test_generateGitignore(t *testing.T) {

	positive_gitignore, err := read("../test/.gitignore")
	if err != nil {
		log.Fatal("error reading file")
	}
	type args struct {
		toptal string
		params string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"positive_test", args{"https://www.toptal.com/developers/gitignore/api/", "go"}, positive_gitignore, false}, // compare with an existing .gitignore in test directory
		{"negative_test", args{"https://wwwwww.toptal.com/developers/gitignore/api/", "go"}, []byte("abc"), true},    // compare with an existing .gitignore in test directory
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateGitignore(tt.args.toptal, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateGitignore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateGitignore() = %v, want %v", got, tt.want)
			}
		})
	}
}
