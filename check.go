package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/vito/progrock"
)

const suspiciousDelta = 10 * 60 // 10 minutes

func check(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)

	var lastStarted int64

	for {
		var entry progrock.StatusUpdate
		if err := dec.Decode(&entry); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		for _, v := range entry.Vertexes {
			v := v
			started := v.Started.GetSeconds()
			if started != 0 {
				if lastStarted != 0 && started < (lastStarted-suspiciousDelta) {
					marshalled, err := json.Marshal(&v)
					if err != nil {
						return err
					}
					fmt.Fprintf(os.Stderr, "=== suspicious entry (started %d seconds before sibiling vertex )===\n", started-lastStarted)
					fmt.Fprintf(os.Stderr, "%s\n", string(marshalled))
					fmt.Fprintf(os.Stderr, "===========\n")
				}

				lastStarted = started
			}
		}
	}
}
