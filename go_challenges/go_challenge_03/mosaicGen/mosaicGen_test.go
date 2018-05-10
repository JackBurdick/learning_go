package mosaicGen

import (
	"fmt"
	"testing"
)

func Test_createMosaic(t *testing.T) {
	var cases = []struct {
		tarImgDir string
		tarImgP   string
		mosDir    string
	}{
		{
			"../input/target/",
			"abby_jack.png",
			"../input/mosaic/PCB_square_png",
		},
	}
	for _, c := range cases {
		fPath := c.tarImgDir + c.tarImgP
		outPath, err := createMosaic(fPath, c.mosDir)
		if err != nil {
			fmt.Printf("things aren't happening: %v\n", err)
		}
		fmt.Printf("OutPath: %v", outPath)
	}
}
