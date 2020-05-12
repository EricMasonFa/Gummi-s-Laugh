package main

import (
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/robotn/gohook"
)

var (
	modifier    bool
	lKeyPressed bool
	mLock       sync.RWMutex
)

func main() {
	EvChan := hook.Start()

	go playsound()

	for ev := range EvChan {
		if modifier && !lKeyPressed {
			if ev.Kind == hook.KeyDown && ev.Rawcode == 37 {
				mLock.Lock()
				lKeyPressed = true
				mLock.Unlock()
				continue
			}
		}
		if ev.Kind == hook.KeyUp && ev.Rawcode == 37 {
			lKeyPressed = false
		}
		if ev.Kind == hook.KeyHold && ev.Rawcode == 55 {
			mLock.Lock()
			modifier = true
			mLock.Unlock()
		}

		if ev.Kind == hook.KeyUp && ev.Rawcode == 55 {
			mLock.Lock()
			modifier = false
			mLock.Unlock()
		}
	}

	hook.End()
}

func playsound() {
	for {
		mLock.RLock()
		if modifier && lKeyPressed {
			mp3file, err := os.Open("./laugh.mp3")
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
