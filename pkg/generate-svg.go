package pkg

import (
	"fmt"
)
func GenerateSVG(number int) string {
	const template = `<svg xmlns="http://www.w3.org/2000/svg">
	<!-- 外側の丸みのある長方形 -->
	<rect x="5" y="5" width="290" height="110" rx="20" ry="20" fill="pink" />

	<!-- 内側の四角形 -->
	<rect x="10" y="10" width="280" height="100" rx="20" ry="20" fill="none" stroke="white" stroke-width="5" />
	<!-- 皿画像 -->
	<image x="40" y="30" width="90" height="70" href="/assets/dish.png" />
	<!-- 皿に表示する文字 -->
	<text x="79" y="90">%d</text>
	<!-- 猫画像 -->
	<image x="170" y="20" width="90" height="80" href="/assets/neko.png" />
  </svg>` 
	svg := fmt.Sprintf(template, number)
	return svg
}