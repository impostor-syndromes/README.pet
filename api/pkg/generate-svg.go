package pkg

import (
	"fmt"
)

func GenerateSVG(number int) string {
	number = 11
	dishImage, charaImage, effectImage := chooseImages(number)

	const template = `<svg xmlns="http://www.w3.org/2000/svg">
	<!-- 外側の丸みのある長方形 -->
	<rect x="5" y="5" width="290" height="110" rx="20" ry="20" fill="pink" />

	<!-- 内側の四角形 -->
	<rect x="10" y="10" width="280" height="100" rx="20" ry="20" fill="none" stroke="white" stroke-width="5" />
	<!-- 皿画像 -->
	<image x="40" y="30" width="90" height="70" href="/assets/dishes/%s" />
	<!-- 皿に表示する文字 -->
	<text x="79" y="90">%d</text>
	<!-- 猫画像 -->
	<image x="170" y="20" width="90" height="80" href="/assets/charactors/%s" />
	<!-- キラキラ -->
	<image x="240" y="15" width="40" height="40" href="/assets/effects/%s" />
  </svg>`
	svg := fmt.Sprintf(template, dishImage, number, charaImage, effectImage)
	return svg
}

func chooseImages(number int) (string, string, string) {
	if number == 0 {
		// 0コミット
		return "empty.png", "neko_cry.png", "clean.png"
	} else if number <= 5 {
		// 1~5コミット
		return "small.png", "neko_normal.png", "clean.png"
	} else if number <= 10 {
		// 6~10コミット
		return "medium.png", "neko_smile.png", "stars_blinking.gif"
	} else {
		// 11コミット以上
		return "large.png", "neko_happy.png", "stars.png"
	}
}
