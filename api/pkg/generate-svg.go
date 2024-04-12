package pkg

import (
	"fmt"
	"encoding/base64"
    "os"
	"io"
)

func GenerateSVG(number int) string {
	dishImage, charaImage, effectImage := chooseImages(number)

	const template = `<svg xmlns="http://www.w3.org/2000/svg">
	<!-- 外側の丸みのある長方形 -->
	<rect x="5" y="5" width="290" height="110" rx="20" ry="20" fill="pink" />

	<!-- 内側の四角形 -->
	<rect x="10" y="10" width="280" height="100" rx="20" ry="20" fill="none" stroke="white" stroke-width="5" />
	<!-- 皿画像 -->
	<image x="-5" y="5" width="180" height="150" href="data:image/png;base64,%s" />
	<!-- 皿に表示する文字 -->
	<text x="78" y="100">%d</text>
	<!-- 猫画像 -->
	<image x="0" y="-135" width="400" height="400" href="data:image/png;base64,%s" />
	<!-- キラキラ -->
	<image x="190" y="15" width="40" height="40" href="data:image/gif;base64,%s" />
  </svg>`
	svg := fmt.Sprintf(template, dishImage, number, charaImage, effectImage)
	return svg
}

func chooseImages(number int) (string, string, string) {
    if number == 0 {
        return encodeImage("assets/dishes/baitdish-default.png"), encodeImage("assets/charactors/cat-tears.png"), encodeImage("assets/effects/tears.gif")
    } else if number <= 5 {
        return encodeImage("assets/dishes/baitdish-small.png"), encodeImage("assets/charactors/cat-null-phase-2.png"), encodeImage("assets/effects/clean.png")
    } else if number <= 10 {
        return encodeImage("assets/dishes/baitdish-medium.png"), encodeImage("assets/charactors/cat-default.png"), encodeImage("assets/effects/stars1.gif")
    } else {
        return encodeImage("assets/dishes/baitdish-large.png"), encodeImage("assets/charactors/cat-smile.png"), encodeImage("assets/effects/stars2.gif")
    }
}

func encodeImage(filePath string) string {
    file, err := os.Open(filePath)
    if err != nil {
        return ""
    }
    defer file.Close()

    fileBytes, err := io.ReadAll(file)
    if err != nil {
        return ""
    }

    return base64.StdEncoding.EncodeToString(fileBytes)
}