package main

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/gabrie30/ghorg/cmd"
	"github.com/joho/godotenv"
	homedir "github.com/mitchellh/go-homedir"
)

func init() {
	home, err := homedir.Dir()

	if err != nil {
		log.Fatal("Error trying to find users home directory")
	}

	err = godotenv.Load(home + "/.ghorg")
	if err != nil {
		log.Fatal("Error loading .ghorg file, create a .env from the sample and run Make install")
	}

	withTrailingSlash := ensureTrailingSlash(os.Getenv("ABSOLUTE_PATH_TO_CLONE_TO"))
	os.Setenv("ABSOLUTE_PATH_TO_CLONE_TO", withTrailingSlash)
}

func ensureTrailingSlash(path string) string {
	if string(path[len(path)-1]) == "/" {
		return path
	}

	return path + "/"
}

func main() {
	color.New(color.FgYellow).Println(
		`
 +-+-+-+-+ +-+-+ +-+-+-+-+-+
 |T|I|M|E| |T|O| |G|H|O|R|G|
 +-+-+-+-+ +-+-+ +-+-+-+-+-+
`)
	cmd.CloneAllReposByOrg()
}
