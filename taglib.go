// Go wrapper for taglib
package taglib

// #cgo LDFLAGS: -ltag_c
// #include <stdlib.h>
// #include <taglib/tag_c.h>
import "C"

import (
	"errors"
	"time"
	"unsafe"
)

type File struct {
	fp    *C.TagLib_File
	tag   *C.TagLib_Tag
	props *C.TagLib_AudioProperties
}

var ErrInvalid = errors.New("invalid file")

// Reads and parses a music file. Returns an error if the provided filename is
// not a valid file.
func Read(filename string) (*File, error) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	fp := C.taglib_file_new(cs)
	if fp == nil || C.taglib_file_is_valid(fp) == 0 {
		return nil, ErrInvalid
	}

	return &File{
		fp:    fp,
		tag:   C.taglib_file_tag(fp),
		props: C.taglib_file_audioproperties(fp),
	}, nil
}

// Close and free the file.
func (file *File) Close() {
	C.taglib_file_free(file.fp)
	file.fp = nil
	file.tag = nil
	file.props = nil
}

func convertAndFree(cs *C.char) string {
	defer C.taglib_free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

// Returns a string with this tag's title.
func (file *File) Title() string {
	return convertAndFree(C.taglib_tag_title(file.tag))
}

// Returns a string with this tag's artist.
func (file *File) Artist() string {
	return convertAndFree(C.taglib_tag_artist(file.tag))
}

// Returns a string with this tag's album name.
func (file *File) Album() string {
	return convertAndFree(C.taglib_tag_album(file.tag))
}

// Returns a string with this tag's comment.
func (file *File) Comment() string {
	return convertAndFree(C.taglib_tag_comment(file.tag))
}

// Returns a string with this tag's genre.
func (file *File) Genre() string {
	return convertAndFree(C.taglib_tag_genre(file.tag))
}

// Returns the tag's year or 0 if year is not set.
func (file *File) Year() int {
	return int(C.taglib_tag_year(file.tag))
}

// Returns the tag's track number or 0 if track number is not set.
func (file *File) Track() int {
	return int(C.taglib_tag_track(file.tag))
}

// Returns the length of the file.
func (file *File) Length() time.Duration {
	length := C.taglib_audioproperties_length(file.props)
	return time.Duration(length) * time.Second
}

// Returns the bitrate of the file in kb/s.
func (file *File) Bitrate() int {
	return int(C.taglib_audioproperties_bitrate(file.props))
}

// Returns the sample rate of the file in Hz.
func (file *File) Samplerate() int {
	return int(C.taglib_audioproperties_samplerate(file.props))
}

// Returns the number of channels in the audio stream.
func (file *File) Channels() int {
	return int(C.taglib_audioproperties_channels(file.props))
}

func init() {
	C.taglib_set_string_management_enabled(0)
}
