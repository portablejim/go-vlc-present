//go:generate goversioninfo -icon=vlc-present.ico -description="Vlc Present"
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)
import "github.com/kardianos/osext"

// Settings holds settings for application.
type Settings struct {
	VlcPath   string
	OtherArgs []string
	FromLeft  int
	FromTop   int
	Width     int
	Height    int
}

func main() {
	if len(os.Args) > 1 {
		file := os.Args[1]

		// Get program directory
		folderPath, err := osext.ExecutableFolder()
		if err != nil {
			log.Fatal(err)
		}

		// Get settings file local to the executable
		settingsFile := path.Join(folderPath, "vlc-present.settings")

		var settingsJSON Settings

		// Read the file and parse to json
		raw, _ := ioutil.ReadFile(settingsFile)
		json.Unmarshal(raw, &settingsJSON)
		if settingsJSON.VlcPath == "" {
			// If path to VLC is missing (e.g. either json parse fail)
			// Use some sane defaults
			settingsJSON = Settings{VlcPath: "C:\\Program\\ Files\\VideoLAN\\VLC\\vlc.exe", OtherArgs: []string{"--fullscreen", "--no-video-title-show", "--no-embedded-video", "--no-qt-fs-controller"}, FromLeft: 0, FromTop: 0, Width: 800, Height: 600}
			settingsString, _ := json.Marshal(settingsJSON)
			ioutil.WriteFile(settingsFile, settingsString, 0664)
		}

		argX := fmt.Sprintf("--video-x=%d", settingsJSON.FromLeft)
		argY := fmt.Sprintf("--video-y=%d", settingsJSON.FromTop)
		argWidth := fmt.Sprintf("--width=%d", settingsJSON.Width)
		argHeight := fmt.Sprintf("--height=%d", settingsJSON.Height)

		vlcArgs0 := []string{file, argX, argY, argWidth, argHeight}
		vlcArgs := append(vlcArgs0, settingsJSON.OtherArgs...)

		//fmt.Println("Vlc path:", settingsJSON.VlcPath)
		//fmt.Println("Vlc args:", vlcArgs)
		exec.Command(settingsJSON.VlcPath, vlcArgs...).Run()

	} else {
		fmt.Println("Usage:", os.Args[0], "<filename>")
	}

}
