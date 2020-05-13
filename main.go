package main

import (
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/rakyll/statik/fs"
	"github.com/robotn/gohook"

	_ "gummi/statik"
)

var (
	lKeyPressed bool
	mLock       sync.RWMutex
	keys        = []rune{108, 76}
)

func main() {
	EvChan := hook.Start()

	go playsound()

	for ev := range EvChan {
		if !lKeyPressed {
			if ev.Kind == hook.KeyDown && contains(keys, ev.Keychar) {
				mLock.Lock()
				lKeyPressed = true
				mLock.Unlock()
				continue
			}
		}

		if ev.Kind == hook.KeyUp && ev.Keycode == 38 {
			lKeyPressed = false
		}
	}

	hook.End()
}

func contains(hystack []rune, needle rune) bool {
	for _, item := range hystack {
		if item == needle {
			return true
		}
	}

	return false
}

func playsound() {
	for {
		mLock.RLock()
		//if modifier && lKeyPressed {
		if lKeyPressed {
			statikFS, err := fs.New()
			if err != nil {
				panic(err)
			}

			// Access individual files by their paths.
			mp3file, err := statikFS.Open("/laugh.mp3")
			if err != nil {
				panic(err)
			}
			defer mp3file.Close()

			if err != nil {
				panic(err)
			}

			streamer, format, err := mp3.Decode(mp3file)
			if err != nil {
				panic(err)
			}

			err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

			if err != nil {
				panic(err)
			}

			done := make(chan bool)
			speaker.Play(beep.Seq(streamer, beep.Callback(func() {
				done <- true
			})))
			<-done

			speaker.Close()
			_ = streamer.Close()
		}
		mLock.RUnlock()
	}
}
