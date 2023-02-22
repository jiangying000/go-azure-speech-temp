package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
)
func main() {
	fmt.Println("startx")
	speechKey := os.Getenv("SPEECH_KEY")
	speechRegion := os.Getenv("SPEECH_REGION")

	// file := "./ch_m.wav"
	file := "./OSR_us_000_0010_8k.wav"

	audioConfig, err := audio.NewAudioConfigFromWavFileInput(file)
	if err != nil {
		fmt.Println("Got an error1: ", err)
		return
	}
	defer audioConfig.Close()

	config, err := speech.NewSpeechConfigFromSubscription(speechKey, speechRegion)
	if err != nil {
		fmt.Println("Got an error:2 ", err)
		return
	}
	defer config.Close()

	speechRecognizer, err := speech.NewSpeechRecognizerFromConfig(config, audioConfig)
	if err != nil {
		fmt.Println("Got an error:3 ", err)
		return
	}
	defer speechRecognizer.Close()

	task := speechRecognizer.RecognizeOnceAsync()
	var outcome speech.SpeechRecognitionOutcome
	select {
	case outcome = <-task:
		fmt.Println("got", outcome)
	case <-time.After(20 * time.Second):
		fmt.Println("Timed out")
		return
	}
	fmt.Println("selected")
	defer outcome.Close()
	if outcome.Error != nil {
		fmt.Println("Got an error: ", outcome.Error)
	}
	fmt.Println("Got a recognition!")
	fmt.Println(outcome)
	fmt.Println(outcome.Result)
	fmt.Println(outcome.Result.Text)
	// fmt.Println(outcome.Result.Reason)
	// fmt.Println(outcome.Result.Text)
}
