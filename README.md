# cv


### Gaussian Blur

There is an example in `blur_test.go`

such like this: 
```
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
```

##### Result

before:

![](res/1.jpg)

after:

![](res/blur.jpg)


### RGB - Grayscal

simple implement:

```golang
func RGB2Gray(img image.Image) (dst image.Image) {
	bounds := img.Bounds()
	tmp := image.NewGray(img.Bounds())
	for x := 0; x <= bounds.Max.X; x++ {
		for y := 0; y <= bounds.Max.Y; y++ {
			c := tmp.ColorModel().Convert(img.At(x, y))
			tmp.Set(x, y, c)
		}
	}
	dst = tmp.SubImage(bounds)
	return
}

```