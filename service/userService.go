package service

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/BohdanPatrash/indenticon/dto"

	"github.com/BohdanPatrash/indenticon/repo"
)

func GetAll() []dto.User {
	var users []dto.User = make([]dto.User, 0)
	for _, val := range repo.DB().Users {
		users = append(users, val)
	}
	return users
}

func AddUser(email string) {
	db := repo.DB()
	var hash Hash = GetHash(email)
	fmt.Println("Email: ", email)
	fmt.Println("Hash: ", hash.String())
	user := dto.User{
		Email: email,
		Hash:  hash,
	}
	fmt.Println("user: ", user)
	db.Users[hash.String()] = user
}

func GetByEmail(email string) dto.User {
	return GetByHash(GetHash(email))
}

func GetByHash(hash []byte) dto.User {
	var h Hash = hash[:]
	db := repo.DB()

	return db.Users[h.String()]
}

func CreateImage(hash []byte) error {
	if len(hash) < 32 {
		return fmt.Errorf("Wrong hash size")
	}

	width, height := 256, 256
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	mainColor := color.RGBA{hash[0], hash[1], hash[2], 255}
	secondaryColor := color.RGBA{hash[2], hash[3], hash[4], 255}

	//set indenticon background to soft milky color
	fillRect(image.Point{0, 0}, image.Point{width, height}, img, color.RGBA{247, 232, 217, 255})

	//draw 3 filled rectangles
	fillRect(image.Point{int(hash[4]), int(hash[5])}, image.Point{int(hash[6]), int(hash[7])}, img, secondaryColor)
	fillRect(image.Point{int(hash[8]), int(hash[9])}, image.Point{int(hash[10]), int(hash[11])}, img, secondaryColor)
	fillRect(image.Point{int(hash[12]), int(hash[13])}, image.Point{int(hash[14]), int(hash[15])}, img, secondaryColor)

	//draw 3 unfilled rectangles
	lineRect(image.Point{int(hash[16]), int(hash[17])}, image.Point{int(hash[18]), int(hash[19])}, img, mainColor)
	lineRect(image.Point{int(hash[20]), int(hash[21])}, image.Point{int(hash[22]), int(hash[23])}, img, mainColor)
	lineRect(image.Point{int(hash[24]), int(hash[25])}, image.Point{int(hash[26]), int(hash[27])}, img, mainColor)

	outputFile, err := os.Create("indenticon.png")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	png.Encode(outputFile, img)
	return nil
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
