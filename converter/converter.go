package converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type conv struct {
	from  string
	to    string
	dir   string
	files []string
}

func NewConv(from string, to string, dir string) (*conv, error) {
	// Check if the extension can be converted
	if from != "png" && from != "jpg" && from != "jpeg" {
		return nil, fmt.Errorf("error: %#v is not supported", from)
	}
	if to != "png" && to != "jpg" && to != "jpeg" {
		return nil, fmt.Errorf("error: %#v is not supported", to)
	}
	// Check if input and output are the same
	if from == to {
		return nil, fmt.Errorf("error: -from and -to are same")
	}
	// Check if the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("error: %#v is not exists", dir)
	}
	// Initialize and return a structure
	c := conv{
		from:  from,
		to:    to,
		dir:   dir,
		files: []string{},
	}
	return &c, nil
}

func (c *conv) getFilePaths() error {
	err := filepath.Walk(c.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error: Check if %#v is not exists", c.dir)
		}
		// In the case of a directory
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		// Collect only paths with the selected extension
		if ext == "."+c.from {
			c.files = append(c.files, path)
		}
		return nil
	})
	return err
}

func convert(dst, src string) error {
	// Open image files
	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("cannot open image file. %s", src)
	}
	defer sf.Close()
	// Create a file
	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create image file. %s", dst)
	}
	defer df.Close()
	// Decode images from input files into memory
	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}
	// Change the save format by extension
	switch strings.ToLower(filepath.Ext(dst)) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpeg", ".jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	default:
		err = os.Remove(dst)
	}
	if err != nil {
		return fmt.Errorf("could not export image file. %s", dst)
	}
	return nil
}

func (c *conv) Do() error {
	if err := c.getFilePaths(); err != nil {
		return err
	}
	// Handle one by one
	for _, src := range c.files {
		dst := src[:len(src)-len(filepath.Ext(src))] + "." + c.to
		err := convert(dst, src)
		if err != nil {
			return err
		}
	}
	return nil
}
