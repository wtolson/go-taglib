package taglib

import (
	"testing"
	"time"
)

func TestReadNothing(t *testing.T) {
	file, err := Read("doesnotexist.mp3")

	if file != nil {
		t.Fatal("Returned non nil file struct.")
	}

	if err == nil {
		t.Fatal("Returned nil err.")
	}

	if err != ErrInvalid {
		t.Fatal("Didn't return ErrInvalid")
	}
}

func TestReadDirectory(t *testing.T) {
	file, err := Read("/")

	if file != nil {
		t.Fatal("Returned non nil file struct.")
	}

	if err == nil {
		t.Fatal("Returned nil err.")
	}

	if err != ErrInvalid {
		t.Fatal("Didn't return ErrInvalid")
	}
}

func TestTagLib(t *testing.T) {
	file, err := Read("test.mp3")

	if err != nil {
		panic(err)
		t.Fatalf("Read returned error: %s", err)
	}

	defer file.Close()

	// Test the Tags
	if title := file.Title(); title != "The Title" {
		t.Errorf("Got wrong title: %s", title)
	}

	if artist := file.Artist(); artist != "The Artist" {
		t.Errorf("Got wrong artist: %s", artist)
	}

	if album := file.Album(); album != "The Album" {
		t.Errorf("Got wrong album: %s", album)
	}

	if comment := file.Comment(); comment != "A Comment" {
		t.Errorf("Got wrong comment: %s", comment)
	}

	if genre := file.Genre(); genre != "Booty Bass" {
		t.Errorf("Got wrong genre: %s", genre)
	}

	if year := file.Year(); year != 1942 {
		t.Errorf("Got wrong year: %d", year)
	}

	if track := file.Track(); track != 42 {
		t.Errorf("Got wrong track: %d", track)
	}

	// Test the properties
	if length := file.Length(); length != 42*time.Second {
		t.Errorf("Got wrong length: %s", length)
	}

	if bitrate := file.Bitrate(); bitrate != 128 {
		t.Errorf("Got wrong bitrate: %d", bitrate)
	}

	if samplerate := file.Samplerate(); samplerate != 44100 {
		t.Errorf("Got wrong samplerate: %d", samplerate)
	}

	if channels := file.Channels(); channels != 2 {
		t.Errorf("Got wrong channels: %d", channels)
	}
}
