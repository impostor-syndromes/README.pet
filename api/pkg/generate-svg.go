package pkg

import (
	"fmt"
)

func GenerateSVG(number int) string {
	dishImage, charaImage, effectImage := chooseImages(number)

	const template = `<svg xmlns="http://www.w3.org/2000/svg">
	<!-- 外側の丸みのある長方形 -->
	<rect x="5" y="5" width="290" height="110" rx="20" ry="20" fill="pink" />

	<!-- 内側の四角形 -->
	<rect x="10" y="10" width="280" height="100" rx="20" ry="20" fill="none" stroke="white" stroke-width="5" />
	<!-- 皿画像 -->
	<image x="-5" y="5" width="180" height="150" href="/assets/dishes/%s" />
	<!-- 皿に表示する文字 -->
	<text x="78" y="100">%d</text>
	<!-- 猫画像 -->
	<image x="0" y="-135" width="400" height="400" href="/assets/charactors/%s" />
	<!-- キラキラ -->
	<image x="190" y="15" width="40" height="40" href="/assets/effects/%s" />
  </svg>`
	svg := fmt.Sprintf(template, dishImage, number, charaImage, effectImage)
	return svg
}

func chooseImages(number int) (string, string, string) {
	if number == 0 {
		// 0コミット
		return "baitdish-default.png", "cat-tears.png", "tears.gif"
	} else if number <= 5 {
		// 1~5コミット
		return "baitdish-small.png", "cat-null-phase-2.png", "clean.png"
	} else if number <= 10 {
		// 6~10コミット
		return "baitdish-medium.png", "cat-default.png", "stars1.gif"
	} else {
		// 11コミット以上
		return "baitdish-large.png", "cat-smile.png", "stars2.gif"
	}
}
