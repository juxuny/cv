package cv

import (
	"image"
	"image/jpeg"
	"os"
	"testing"
)

func _saveJpg(img image.Image, fileName string) (e error) {
	out, e := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	//	out, e := os.Create(fileName)
	if e != nil {
		return
	}
	defer out.Close()
	e = jpeg.Encode(out, img, &jpeg.Options{Quality: 70})
	return
}

func TestBlur(t *testing.T) {
	t.Skip()
	f, e := os.Open("./res/1.jpg")
	if e != nil {
		t.Fatal(e)
	}
	defer f.Close()
	t.Log("open file")
	img, _, e := image.Decode(f)
	if e != nil {
		t.Fatal(e)
	}
	dst := GaussianBlur(img, 3)
	e = _saveJpg(dst, "res/blur.jpg")
	if e != nil {
		t.Fatal(e)
	}
}
