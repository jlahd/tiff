// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package tiff

import (
	"image"
	"image/color"
	"reflect"
)

type Gray16 struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewGray16 returns a new Gray16 with the given bounds.
func NewGray16(r image.Rectangle) *Gray16 {
	return new(Gray16).Init(make([]uint8, 2*r.Dx()*r.Dy()), 2*r.Dx(), r)
}

func (p *Gray16) Init(pix []uint8, stride int, rect image.Rectangle) *Gray16 {
	*p = Gray16{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *Gray16) Pix() []byte           { return p.M.Pix }
func (p *Gray16) Stride() int           { return p.M.Stride }
func (p *Gray16) Rect() image.Rectangle { return p.M.Rect }
func (p *Gray16) Channels() int         { return 1 }
func (p *Gray16) Depth() reflect.Kind   { return reflect.Uint16 }

func (p *Gray16) ColorModel() color.Model { return colorGray16Model }

func (p *Gray16) Bounds() image.Rectangle { return p.M.Rect }

func (p *Gray16) At(x, y int) color.Color {
	return p.Gray16At(x, y)
}

func (p *Gray16) Gray16At(x, y int) colorGray16 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorGray16{}
	}
	i := p.PixOffset(x, y)
	return pGray16At(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray16) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*2
}

func (p *Gray16) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorGray16Model.Convert(c).(colorGray16)
	pSetGray16(p.M.Pix[i:], c1)
	return
}

func (p *Gray16) SetGray16(x, y int, c colorGray16) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetGray16(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray16) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Gray16{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(Gray16).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray16) Opaque() bool {
	return true
}
