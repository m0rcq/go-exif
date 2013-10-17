// Copyright 2013 atanas "jack" argirov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package go-exif implements reading EXIF metadata as specified
// in EXIF format specification

// For more information see:
// http://www.media.mit.edu/pia/Research/deepview/exif.html
// http://www.exiv2.org/tags.html

package main

import (
	"flag"
	"fmt"
	exif "go-exif"
	"log"
	"math/big"
	"os"
	"strings"
)

var imageFile string

// PrintIfd prints the supplied IFD entries
func PrintIFD(ifds []exif.IfdEntries) {
	for _, v := range ifds {
		lval, ok := v.Values.([]interface{})
		var values string
		if ok {
			switch val := lval[0].(type) {
			case string:
				values = fmt.Sprintf("'%s'", val)
			case byte:
				values = fmt.Sprintf("%#x", val)
			case []uint8:
				var lstr []string
				for _, v := range lval {
					lstr = append(lstr, fmt.Sprintf("%#x", v))
				}
				values = strings.Join(lstr, ", ")
			case int16:
				values = fmt.Sprintf("%d", val)
			case int32:
				values = fmt.Sprintf("%d", val)
			case int64:
				values = fmt.Sprintf("%d", val)
			case uint16:
				values = fmt.Sprintf("%d", val)
			case uint32:
				values = fmt.Sprintf("%d", val)
			case uint64:
				values = fmt.Sprintf("%d", val)
			case *big.Rat:
				values = fmt.Sprintf("%s", val.RatString())
			default:
				values = fmt.Sprintf("%v", lval)
			}
		}
		fmt.Printf("[%s] (0x%04x) %s(%s) = [%s]\n", exif.IfdSeqMap[v.IfdSeq], v.Tag, v.TagDesc, exif.FormatType[int(v.Format)], values)
	}
}

func init() {
	flag.StringVar(&imageFile, "f", "", "Image File to process")
}

func main() {
	flag.Parse()
	fmt.Println("Image file to process: ", imageFile)
	f, err := os.Open(imageFile)
	if err != nil {
		log.Fatal(err)
	}

	exif := &exif.ExifData{}
	exif.ProcessExifStream(f)

	for k, v := range exif.IfdData {
		fmt.Printf("Section - %s:\n", k)
		PrintIFD(v)
	}

	// get specific tag
	if v, ok := exif.GetTagValues(0x010f); ok != false {
		fmt.Printf("Tag Value = %v\n", v)
		// you need to implement interface conversion as shown in PrintIFD()
	}

}
