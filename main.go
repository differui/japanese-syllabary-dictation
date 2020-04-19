package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Vowel rune
type Consonant rune

const (
	a Vowel = 'a'
	i Vowel = 'i'
	u Vowel = 'u'
	e Vowel = 'e'
	o Vowel = 'o'
)

const (
	empty           = 0
	k     Consonant = 'k'
	s     Consonant = 's'
	t     Consonant = 't'
)

func locate(vowel Vowel, consonant Consonant) (pathname string) {
	prefix := "assets/"
	if consonant == empty {
		pathname = fmt.Sprintf(prefix+"%c.mp3", vowel)
	} else {
		pathname = fmt.Sprintf(prefix+"%c%c.mp3", consonant, vowel)
	}
	return
}

func play(vowel Vowel, consonant Consonant) {
	// read asset
	b, err := Asset(locate(vowel, consonant))
	if err != nil {
		log.Fatal(err)
		return
	}

	// convert bytes to io reader
	r := ioutil.NopCloser(bytes.NewReader(b))

	// decode map file
	streamer, format, err := mp3.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// speak
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	// wait for voice ends
	time.Sleep(time.Second * 3)
}

func main() {
	play(a, empty)
	play(i, empty)
	play(u, empty)
	play(e, empty)
	play(o, empty)
}
