# Binaural Beats Generator

This program generates binaural beats audio files using provided frequencies for the left and right audio channels. Binaural beats are a form of auditory illusion perceived when two different pure-tone sine waves are presented to a listener dichotically (one through each ear). They have been explored for their potential in various therapeutic settings.

## Features

- Generate custom binaural beats by specifying the frequencies for the left and right audio channels.
- Output the generated beats to a WAV file.
- The duration of the binaural beat audio is dynamically calculated based on the least common multiple of the two input frequencies, ensuring a seamless loop.

## Prerequisites

Before running this program, ensure you have Go installed on your system. You also need to have the `github.com/FridaFino/goalgorithms/math/lcm`, `github.com/go-audio/audio`, and `github.com/go-audio/wav` packages installed.

## Installation

First, clone the repository to your local machine:

```bash
git clone github.com/gbasilveira/binaural-beats-generator
```

Navigate to the directory containing the program:

```bash
cd binarual-beats-generator
```

## Usage

To use the program, run the following command in your terminal, replacing `<freqLeft>`, `<freqRight>`, and `<filename>` with your desired frequencies for the left and right channels (in Hz) and the output filename (including the `.wav` extension), respectively:

```bash
go run main.go <freqLeft> <freqRight> <filename>
```

Example:

```bash
go run main.go 100 108 binaural_beat.wav
```

This command generates a binaural beat with left and right frequencies of 440 Hz and 442 Hz, respectively, and saves it as `binaural_beat.wav`.

## Contributing

Contributions to improve the program are welcome. Please follow the standard pull request process to submit your changes.

## License

[MIT License](LICENSE.md)
