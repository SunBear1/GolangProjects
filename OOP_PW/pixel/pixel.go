package pixel

import "fmt"

type Point struct {
	x float32
	y float32
}

type pixel struct {
	luminance int
	location  Point
}

func CreatePixel(_luminance int, _x float32, _y float32) pixel {
	return pixel{location: Point{x: _x, y: _y}, luminance: _luminance}
}

func (p *pixel) ToString() {
	fmt.Printf("Luminance: %d, X coordinate: %f, Y coordinate: %f\n", p.luminance, p.location.x, p.location.y)
}

func (p *pixel) GetLuminance() int {
	return p.luminance
}

func (p *pixel) SetLuminance(luminance int) {
	p.luminance = luminance
}
