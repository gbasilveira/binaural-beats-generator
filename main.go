package main

import (
	"math"
	"os"
	"strconv"

	"github.com/FridaFino/goalgorithms/math/lcm"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func main() {
	// Sample rate and duration
	sampleRate := 44100

	// Frequencies for the binaural beats
	freqLeftInt, _ := strconv.Atoi(os.Args[1])
	freqRightInt, _ := strconv.Atoi(os.Args[2])
	filename := os.Args[3]

	freqLeft := float64(freqLeftInt)
	freqRight := float64(freqRightInt)

	// Duration of the binaural is calculated by the least common multiple  of the two frequencies
	duration := int(lcm.Lcm(int64(freqLeft), int64(freqRight)))

	// Create a buffer for each channel
	bufLeft := make([]int, sampleRate*duration)
	bufRight := make([]int, sampleRate*duration)

	// Generate sine waves for left and right channels
	for i := 0; i < sampleRate*duration; i++ {
		t := float64(i) / float64(sampleRate)
		bufLeft[i] = int(math.Sin(2*math.Pi*freqLeft*t) * 32767)
		bufRight[i] = int(math.Sin(2*math.Pi*freqRight*t) * 32767)
	}

	// Output file
	outFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Create a WAV encoder and write the buffers to the WAV file
	encoder := wav.NewEncoder(outFile, sampleRate, 16, 2, 1)

	// Interleave the buffers for stereo
	interleaveBuf := make([]int, 2*sampleRate*duration)
	for i := 0; i < sampleRate*duration; i++ {
		interleaveBuf[2*i] = bufLeft[i]
		interleaveBuf[2*i+1] = bufRight[i]
	}

	pcm := audio.IntBuffer{Data: interleaveBuf, Format: &audio.Format{SampleRate: sampleRate, NumChannels: 2}}

	if err := encoder.Write(&pcm); err != nil {
		panic(err)
	}
	if err := encoder.Close(); err != nil {
		panic(err)
	}

	println("Binaural beats file created.")
}
