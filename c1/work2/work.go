package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	elseIndex  = 2 // next color in palette
)

func lissajous(out io.Writer, cycles float64) {
	const (
		// number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				elseIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, err := strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		cycles = 5
	}
	lissajous(w, float64(cycles))
}

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
