package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/matrix-org/gomatrix"
)

func main() {
	BASE_URL := os.Getenv("MMM_BASE_URL") // for use with url shotner ("sub" mode in meain/sirus)
	if BASE_URL == "" {
		BASE_URL = "https://matrix.org/_matrix/media/r0/download/matrix.org"
	}

	fmt.Println("Starting with user", os.Getenv("MMM_MATRIX_USER"))
	cli, _ := gomatrix.NewClient(os.Getenv("MMM_MATRIX_SERVER"), os.Getenv("MMM_MATRIX_USER"), os.Getenv("MMM_MATRIX_PASSWORD"))
	syncer := cli.Syncer.(*gomatrix.DefaultSyncer)

	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		if ev.Sender == os.Getenv("MMM_MATRIX_USER") {
			return
		}
		mtype, ok := ev.MessageType()
		if !ok {
			fmt.Print("unable to figure out message type")
			return
		}
		if mtype == "m.image" || mtype == "m.video" || mtype == "m.file" {
			url := ev.Content["url"].(string)
			splits := strings.Split(url, "/")
			hash := splits[len(splits)-1]
			cli.SendNotice(ev.RoomID, "[URL]: "+BASE_URL+"/"+hash)
		}
	})

	if err := cli.Sync(); err != nil {
		fmt.Println("Sync() returned ", err)
	}
}
