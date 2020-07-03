package service

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

func CreateIndenticon(hash []byte) ([]byte, error) {
	if len(hash) < 32 {
		return nil, fmt.Errorf("Wrong hash size")
	}

	width, height := 256, 256
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	mainColor := color.RGBA{hash[0], hash[1], hash[2], 255}
	secondaryColor := color.RGBA{hash[2], hash[3], hash[4], 255}

	//set indenticon background to soft milky color
	fillRect(image.Point{0, 0}, image.Point{width, height}, img, color.RGBA{247, 232, 217, 255})

	//draw 3 filled rectangles
	for i := 4; i < 15; i += 4 {
		fillRect(image.Point{int(hash[i]), int(hash[i+1])}, image.Point{int(hash[i+2]), int(hash[i+3])}, img, secondaryColor)
	}

	//draw 3 unfilled rectangles
	for i := 16; i < 32; i += 4 {
		lineRect(image.Point{int(hash[i]), int(hash[i+1])}, image.Point{int(hash[i+2]), int(hash[i+3])}, img, mainColor)
	}

	var buff bytes.Buffer
	png.Encode(&buff, img)

	return buff.Bytes(), nil
}

func drawLine(A, B image.Point, img *image.RGBA, col color.RGBA) {
	if A.X == B.X {
		if A.Y > B.Y {
			A, B = B, A
		}
		for y := A.Y; y <= B.Y; y++ {
			img.Set(A.X, y, col)
		}
		return
	}

	if A.X > B.X {
		A, B = B, A
	}
	dx := B.X - A.X
	dy := B.Y - A.Y

	for x := A.X; x <= B.X; x++ {
		y := A.Y + dy*(x-A.X)/dx
		img.Set(x, y, col)
	}
}

func lineRect(A, B image.Point, img *image.RGBA, col color.RGBA) {
	drawLine(A, image.Point{A.X, B.Y}, img, col)
	drawLine(image.Point{A.X, B.Y}, B, img, col)
	drawLine(B, image.Point{B.X, A.Y}, img, col)
	drawLine(image.Point{B.X, A.Y}, A, img, col)
}

func fillRect(A, B image.Point, img *image.RGBA, col color.RGBA) {
	if A.X > B.X {
		A, B = B, A
	}
	for i := A.X; i <= B.X; i++ {
		if A.Y > B.Y {
			for j := A.Y; j >= B.Y; j-- {
				img.SetRGBA(i, j, col)
			}
		} else {
			for j := A.Y; j <= B.Y; j++ {
				img.SetRGBA(i, j, col)
			}
		}
	}
}
