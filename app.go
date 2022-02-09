package main

import (
	"image/color"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/julienschmidt/httprouter"
	"github.com/notmahi/fnts/arts"
	"github.com/notmahi/fnts/common"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	r := httprouter.New()
	r.GET("/", drawHandlerNew)
	r.GET("/art/:id", generalizedHandler)

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func generalizedHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	log.Printf("id=%s", id)

	// First, build a map for handling routes
	m := map[string]func() ([]byte, error){
		"blackhole":    drawBlackhole,
		"circlegrid":   drawCircleGrid,
		"circleline":   drawCircleLine,
		"circleloop":   drawCircleLoop2,
		"circlemove":   drawCircleMove,
		"circlenoise":  drawCircleNoise,
		"colorcanva":   drawColorCanva,
		"colorcircle":  drawColorCircle,
		"colorcircle2": drawColorCircle2,
		"contourline":  drawContourLine,
		"domainwrap":   drawDomainWrap,
		// "domainwrap2":  drawDomainWrap2,
		"dotline":      drawDotLine,
		"dotswave":     drawDotsWave,
		"gridsquare":   drawGridSquare,
		"janus":        drawJanus,
		"julia":        drawJulia,
		"maze":         drawMaze,
		"noiseline":    drawNoiseLine,
		"oceanfish":    drawOceanFish,
		"perlinpearls": drawPerlinPearls,
		"pixelhole":    drawPixelHole,
		// "pointribbon":  drawPointRibbon,
		"randcircle":  drawRandCircle,
		"randomshape": drawRandomShape,
		"silkysky":    drawSilkySky,
		"silksmoke":   drawSilkSmoke,
		// "solarflare":   drawSolarFlare,
		"spiralsquare": drawSpiralSquare,
		// "swirl":        drawSwirl,
		"yarn": drawYarn,
	}

	// Get the function from the map
	f := m[id]
	b, err := f()
	if err != nil {
		log.Fatal(err)
	}

	// Set content headers
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))

	// Write image to response
	if _, err = w.Write(b); err != nil {
		log.Fatal("unable to write image.")
	}
}

func drawHandlerNew(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Log Requests
	log.Printf("method=%s path=%s ", r.Method, r.RequestURI)

	// Draw the image to bytes
	b, err := drawMyBytes()
	if err != nil {
		log.Fatal(err)
	}

	// Set content headers
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))

	// Write image to response
	if _, err = w.Write(b); err != nil {
		log.Fatal("unable to write image.")
	}
}

func drawMyBytes() ([]byte, error) {
	// Generate a new image
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)

	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema([]color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	})

	c.Draw(arts.NewRandomShape(150))

	// Return the image as []byte
	return c.ToBytes()
}

func drawBlackhole() ([]byte, error) {
	// Draw a black hole generative art
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(color.RGBA{R: 30, G: 30, B: 30, A: 255})
	c.FillBackground()
	c.SetLineWidth(1.0)
	c.SetLineColor(common.Tomato)
	c.Draw(arts.NewBlackHole(200, 400, 0.01))

	// Return the image as []byte
	return c.ToBytes()
}

func drawCircleGrid() ([]byte, error) {
	// Draw a circle line generative art

	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xED, 0x34, 0x41, 0xFF},
		{0xFF, 0xD6, 0x30, 0xFF},
		{0x32, 0x9F, 0xE3, 0xFF},
		{0x15, 0x42, 0x96, 0xFF},
		{0x00, 0x00, 0x00, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(color.RGBA{0xDF, 0xEB, 0xF5, 0xFF})
	c.FillBackground()
	c.SetColorSchema(colors)
	c.SetLineWidth(2.0)
	c.Draw(arts.NewCircleGrid(4, 6))
	return c.ToBytes()
}

func drawCircleLine() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(600, 600)
	c.SetBackground(common.Tan)
	c.SetLineWidth(1.0)
	c.SetLineColor(common.LightPink)
	c.FillBackground()
	c.Draw(arts.NewCircleLine(0.02, 600, 1.5, 2, 2))
	return c.ToBytes()
}

func drawCircleLoop() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Black)
	c.SetLineWidth(1)
	c.SetLineColor(common.Orange)
	c.SetAlpha(30)
	c.SetIterations(1000)
	c.FillBackground()
	c.Draw(arts.NewCircleLoop(100))
	return c.ToBytes()
}

func drawCircleLoop2() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(color.RGBA{8, 10, 20, 255})
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewCircleLoop2(7))
	return c.ToBytes()
}

func drawCircleMove() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(1200, 500)
	c.SetBackground(common.White)
	c.FillBackground()
	c.Draw(arts.NewCircleMove(1000))
	return c.ToBytes()
}

func drawCircleNoise() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.White)
	c.SetAlpha(5)
	c.SetLineWidth(0.3)
	c.FillBackground()
	c.SetIterations(100)
	c.Draw(arts.NewCircleNoise(2000, 20, 80))
	return c.ToBytes()
}

func drawColorCanva() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetLineWidth(8)
	c.SetColorSchema(colors)
	c.Draw(arts.NewColorCanve(5))
	return c.ToBytes()
}

func drawColorCircle() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xFF, 0xC6, 0x18, 0xFF},
		{0xF4, 0x25, 0x39, 0xFF},
		{0x41, 0x78, 0xF4, 0xFF},
		{0xFE, 0x84, 0xFE, 0xFF},
		{0xFF, 0x81, 0x19, 0xFF},
		{0x56, 0xAC, 0x51, 0xFF},
		{0x98, 0x19, 0xFA, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewColorCircle(200))
	return c.ToBytes()
}

func drawColorCircle2() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewColorCircle2(15))
	return c.ToBytes()
}

func drawContourLine() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0x58, 0x18, 0x45, 0xFF},
		{0x90, 0x0C, 0x3F, 0xFF},
		{0xC7, 0x00, 0x39, 0xFF},
		{0xFF, 0x57, 0x33, 0xFF},
		{0xFF, 0xC3, 0x0F, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(color.RGBA{0x1a, 0x06, 0x33, 0xFF})
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewContourLine(200))
	return c.ToBytes()
}

func cmap(r, m1, m2 float64) color.RGBA {
	rgb := color.RGBA{
		R: uint8(common.Constrain(m1*200*r, 0, 255)),
		G: uint8(common.Constrain(r*200, 0, 255)),
		B: uint8(common.Constrain(m2*255*r, 70, 255)),
		A: 255,
	}
	return rgb
}

func drawDomainWrap() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.Draw(arts.NewDomainWrap(0.01, 4, 4, 20, cmap))

	return c.ToBytes()
}

func drawDomainWrap2() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Black)
	c.FillBackground()
	d := arts.NewDomainWrap(0.01, 4, 4, 20, cmap)
	d.SetDynamicParameter(0.005, 0, 100, "./temp")
	c.Draw(d)
	return c.ToBytes()
}

func drawDotLine() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(2080, 2080)
	c.SetBackground(color.RGBA{230, 230, 230, 255})
	c.SetLineWidth(10)
	c.SetIterations(15000)
	c.SetColorSchema(common.DarkPink)
	c.FillBackground()
	c.Draw(arts.NewDotLine(100, 20, 50, false))
	return c.ToBytes()
}

func drawDotsWave() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xFF, 0xBE, 0x0B, 0xFF},
		{0xFB, 0x56, 0x07, 0xFF},
		{0xFF, 0x00, 0x6E, 0xFF},
		{0x83, 0x38, 0xEC, 0xFF},
		{0x3A, 0x86, 0xFF, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewDotsWave(300))
	return c.ToBytes()
}

func drawGridSquare() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(600, 600)
	c.SetBackground(common.DarkPink[rand.Intn(5)])
	c.SetColorSchema(common.DarkPink)
	c.Draw(arts.NewGirdSquares(24, 10, 0.2))
	return c.ToBytes()
}

func drawJanus() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetColorSchema(common.DarkRed)
	c.SetForeground(common.LightPink)
	c.Draw(arts.NewJanus(10, 0.2))
	return c.ToBytes()
}

func julia1(z complex128) complex128 {
	c := complex(-0.1, 0.651)

	z = z*z + c

	return z
}

func drawJulia() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetIterations(800)
	c.SetColorSchema(common.Citrus)
	c.FillBackground()
	c.Draw(arts.NewJulia(julia1, 40, 1.5, 1.5))
	return c.ToBytes()
}

func drawMaze() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(600, 600)
	c.SetBackground(common.Azure)
	c.SetLineWidth(3)
	c.SetLineColor(common.Orange)
	c.FillBackground()
	c.Draw(arts.NewMaze(20))
	return c.ToBytes()
}

func drawNoiseLine() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0x06, 0x7B, 0xC2, 0xFF},
		{0x84, 0xBC, 0xDA, 0xFF},
		{0xEC, 0xC3, 0x0B, 0xFF},
		{0xF3, 0x77, 0x48, 0xFF},
		{0xD5, 0x60, 0x62, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(color.RGBA{0xF0, 0xFE, 0xFF, 0xFF})
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewNoiseLine(500))
	return c.ToBytes()
}

func drawOceanFish() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetColorSchema(colors)
	c.Draw(arts.NewOceanFish(100, 8))
	return c.ToBytes()
}

func drawPerlinPearls() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.White)
	c.SetAlpha(120)
	c.SetLineWidth(0.3)
	c.FillBackground()
	c.SetIterations(200)
	c.Draw(arts.NewPerlinPerls(10, 200, 40, 80))
	return c.ToBytes()
}

func drawPixelHole() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := generativeart.NewCanva(600, 600)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetColorSchema(colors)
	c.SetIterations(800)
	c.Draw(arts.NewPixelHole(50))
	return c.ToBytes()
}

func drawPointRibbon() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Lavender)
	c.SetLineWidth(2)
	c.SetIterations(150000)
	c.FillBackground()
	c.Draw(arts.NewPointRibbon(50))
	return c.ToBytes()
}

func drawRandCircle() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.MistyRose)
	c.SetLineWidth(1.0)
	c.SetLineColor(color.RGBA{
		R: 122,
		G: 122,
		B: 122,
		A: 30,
	})
	c.SetColorSchema(common.DarkPink)
	c.SetIterations(4)
	c.FillBackground()
	c.Draw(arts.NewRandCicle(30, 80, 0.2, 2, 10, 30, true))
	return c.ToBytes()
}

func drawRandomShape() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema([]color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	})
	c.Draw(arts.NewRandomShape(150))
	return c.ToBytes()
}

func drawSilkySky() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(600, 600)
	c.SetAlpha(10)
	c.Draw(arts.NewSilkSky(15, 5))
	return c.ToBytes()
}

func drawSilkSmoke() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Black)
	c.SetLineWidth(1.0)
	c.SetLineColor(common.MediumAquamarine)
	c.SetAlpha(30)
	c.SetColorSchema(common.DarkPink)
	c.SetIterations(4)
	c.FillBackground()
	c.Draw(arts.NewSilkSmoke(400, 20, 0.2, 2, 10, 30, false))
	return c.ToBytes()
}

func drawSolarFlare() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetLineColor(color.RGBA{255, 64, 8, 128})
	c.Draw(arts.NewSolarFlare())
	return c.ToBytes()
}

func drawSpiralSquare() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.MistyRose)
	c.SetLineWidth(10)
	c.SetLineColor(common.Orange)
	c.SetColorSchema(common.DarkPink)
	c.SetForeground(common.Tomato)
	c.FillBackground()
	c.Draw(arts.NewSpiralSquare(40, 400, 0.05, true))
	return c.ToBytes()
}

func drawSwirl() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Azure)
	c.FillBackground()
	c.SetForeground(color.RGBA{113, 3, 0, 140})
	c.SetIterations(400000)
	c.Draw(arts.NewSwirl(0.970, -1.899, 1.381, -1.506, 2.4, 2.4))
	return c.ToBytes()
}

func drawYarn() ([]byte, error) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Orange)
	c.FillBackground()
	c.SetLineWidth(0.3)
	c.SetLineColor(color.RGBA{A: 60})
	c.Draw(arts.NewYarn(2000))
	return c.ToBytes()
}
