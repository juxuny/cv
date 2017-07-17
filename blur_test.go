package cv

import (
	"image"
	"image/jpeg"
	"os"
	"testing"
)

func _saveJpg(img image.Image, fileName string) (e error) {
	out, e := os.Create(fileName)
	if e != nil {
		return
	}
	defer out.Close()
	e = jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
	return
}

func TestBlur(t *testing.T) {
	f, e := os.Open("./res/1.jpg")
	if e != nil {
		t.Fatal(e)
	}
	defer f.Close()
	img, _, e := image.Decode(f)
	if e != nil {
		t.Fatal(e)
	}
	dst := GaussianBlur(img, 5)
	e = _saveJpg(dst, "res/blur.jpg")
	if e != nil {
		t.Fatal(e)
	}
}
