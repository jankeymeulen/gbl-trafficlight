package main

type Colour struct {
	Name string
	R    uint8
	G    uint8
	B    uint8
}

func getColourMap() map[string]Colour {
	var m map[string]Colour
	m = make(map[string]Colour)
	m["aliceblue"] = Colour{"aliceblue", 240, 248, 255}
	m["antiquewhite"] = Colour{"antiquewhite", 250, 235, 215}
	m["aqua"] = Colour{"aqua", 0, 255, 255}
	m["aquamarine"] = Colour{"aquamarine", 127, 255, 212}
	m["azure"] = Colour{"azure", 240, 255, 255}
	m["beige"] = Colour{"beige", 245, 245, 220}
	m["bisque"] = Colour{"bisque", 255, 228, 196}
	m["black"] = Colour{"black", 0, 0, 0}
	m["blanchedalmond"] = Colour{"blanchedalmond", 255, 235, 205}
	m["blue"] = Colour{"blue", 0, 0, 255}
	m["blueviolet"] = Colour{"blueviolet", 138, 43, 226}
	m["brown"] = Colour{"brown", 165, 42, 42}
	m["burlywood"] = Colour{"burlywood", 222, 184, 135}
	m["cadetblue"] = Colour{"cadetblue", 95, 158, 160}
	m["chartreuse"] = Colour{"chartreuse", 127, 255, 0}
	m["chocolate"] = Colour{"chocolate", 210, 105, 30}
	m["coral"] = Colour{"coral", 255, 127, 80}
	m["cornflowerblue"] = Colour{"cornflowerblue", 100, 149, 237}
	m["cornsilk"] = Colour{"cornsilk", 255, 248, 220}
	m["crimson"] = Colour{"crimson", 220, 20, 60}
	m["cyan"] = Colour{"cyan", 0, 255, 255}
	m["darkblue"] = Colour{"darkblue", 0, 0, 139}
	m["darkcyan"] = Colour{"darkcyan", 0, 139, 139}
	m["darkgoldenrod"] = Colour{"darkgoldenrod", 184, 134, 11}
	m["darkgray"] = Colour{"darkgray", 169, 169, 169}
	m["darkgreen"] = Colour{"darkgreen", 0, 100, 0}
	m["darkgrey"] = Colour{"darkgrey", 169, 169, 169}
	m["darkkhaki"] = Colour{"darkkhaki", 189, 183, 107}
	m["darkmagenta"] = Colour{"darkmagenta", 139, 0, 139}
	m["darkolivegreen"] = Colour{"darkolivegreen", 85, 107, 47}
	m["darkorange"] = Colour{"darkorange", 255, 140, 0}
	m["darkorchid"] = Colour{"darkorchid", 153, 50, 204}
	m["darkred"] = Colour{"darkred", 139, 0, 0}
	m["darksalmon"] = Colour{"darksalmon", 233, 150, 122}
	m["darkseagreen"] = Colour{"darkseagreen", 143, 188, 143}
	m["darkslateblue"] = Colour{"darkslateblue", 72, 61, 139}
	m["darkslategray"] = Colour{"darkslategray", 47, 79, 79}
	m["darkslategrey"] = Colour{"darkslategrey", 47, 79, 79}
	m["darkturquoise"] = Colour{"darkturquoise", 0, 206, 209}
	m["darkviolet"] = Colour{"darkviolet", 148, 0, 211}
	m["deeppink"] = Colour{"deeppink", 255, 20, 147}
	m["deepskyblue"] = Colour{"deepskyblue", 0, 191, 255}
	m["dimgray"] = Colour{"dimgray", 105, 105, 105}
	m["dimgrey"] = Colour{"dimgrey", 105, 105, 105}
	m["dodgerblue"] = Colour{"dodgerblue", 30, 144, 255}
	m["firebrick"] = Colour{"firebrick", 178, 34, 34}
	m["floralwhite"] = Colour{"floralwhite", 255, 250, 240}
	m["forestgreen"] = Colour{"forestgreen", 34, 139, 34}
	m["fuchsia"] = Colour{"fuchsia", 255, 0, 255}
	m["gainsboro"] = Colour{"gainsboro", 220, 220, 220}
	m["ghostwhite"] = Colour{"ghostwhite", 248, 248, 255}
	m["gold"] = Colour{"gold", 255, 215, 0}
	m["goldenrod"] = Colour{"goldenrod", 218, 165, 32}
	m["gray"] = Colour{"gray", 128, 128, 128}
	m["grey"] = Colour{"grey", 128, 128, 128}
	m["green"] = Colour{"green", 0, 128, 0}
	m["greenyellow"] = Colour{"greenyellow", 173, 255, 47}
	m["honeydew"] = Colour{"honeydew", 240, 255, 240}
	m["hotpink"] = Colour{"hotpink", 255, 105, 180}
	m["indianred"] = Colour{"indianred", 205, 92, 92}
	m["indigo"] = Colour{"indigo", 75, 0, 130}
	m["ivory"] = Colour{"ivory", 255, 255, 240}
	m["khaki"] = Colour{"khaki", 240, 230, 140}
	m["lavender"] = Colour{"lavender", 230, 230, 250}
	m["lavenderblush"] = Colour{"lavenderblush", 255, 240, 245}
	m["lawngreen"] = Colour{"lawngreen", 124, 252, 0}
	m["lemonchiffon"] = Colour{"lemonchiffon", 255, 250, 205}
	m["lightblue"] = Colour{"lightblue", 173, 216, 230}
	m["lightcoral"] = Colour{"lightcoral", 240, 128, 128}
	m["lightcyan"] = Colour{"lightcyan", 224, 255, 255}
	m["lightgoldenrodyellow"] = Colour{"lightgoldenrodyellow", 250, 250, 210}
	m["lightgray"] = Colour{"lightgray", 211, 211, 211}
	m["lightgreen"] = Colour{"lightgreen", 144, 238, 144}
	m["lightgrey"] = Colour{"lightgrey", 211, 211, 211}
	m["lightpink"] = Colour{"lightpink", 255, 182, 193}
	m["lightsalmon"] = Colour{"lightsalmon", 255, 160, 122}
	m["lightseagreen"] = Colour{"lightseagreen", 32, 178, 170}
	m["lightskyblue"] = Colour{"lightskyblue", 135, 206, 250}
	m["lightslategray"] = Colour{"lightslategray", 119, 136, 153}
	m["lightslategrey"] = Colour{"lightslategrey", 119, 136, 153}
	m["lightsteelblue"] = Colour{"lightsteelblue", 176, 196, 222}
	m["lightyellow"] = Colour{"lightyellow", 255, 255, 224}
	m["lime"] = Colour{"lime", 0, 255, 0}
	m["limegreen"] = Colour{"limegreen", 50, 205, 50}
	m["linen"] = Colour{"linen", 250, 240, 230}
	m["magenta"] = Colour{"magenta", 255, 0, 255}
	m["maroon"] = Colour{"maroon", 128, 0, 0}
	m["mediumaquamarine"] = Colour{"mediumaquamarine", 102, 205, 170}
	m["mediumblue"] = Colour{"mediumblue", 0, 0, 205}
	m["mediumorchid"] = Colour{"mediumorchid", 186, 85, 211}
	m["mediumpurple"] = Colour{"mediumpurple", 147, 112, 219}
	m["mediumseagreen"] = Colour{"mediumseagreen", 60, 179, 113}
	m["mediumslateblue"] = Colour{"mediumslateblue", 123, 104, 238}
	m["mediumspringgreen"] = Colour{"mediumspringgreen", 0, 250, 154}
	m["mediumturquoise"] = Colour{"mediumturquoise", 72, 209, 204}
	m["mediumvioletred"] = Colour{"mediumvioletred", 199, 21, 133}
	m["midnightblue"] = Colour{"midnightblue", 25, 25, 112}
	m["mintcream"] = Colour{"mintcream", 245, 255, 250}
	m["mistyrose"] = Colour{"mistyrose", 255, 228, 225}
	m["moccasin"] = Colour{"moccasin", 255, 228, 181}
	m["navajowhite"] = Colour{"navajowhite", 255, 222, 173}
	m["navy"] = Colour{"navy", 0, 0, 128}
	m["oldlace"] = Colour{"oldlace", 253, 245, 230}
	m["olive"] = Colour{"olive", 128, 128, 0}
	m["olivedrab"] = Colour{"olivedrab", 107, 142, 35}
	m["orange"] = Colour{"orange", 255, 165, 0}
	m["orangered"] = Colour{"orangered", 255, 69, 0}
	m["orchid"] = Colour{"orchid", 218, 112, 214}
	m["palegoldenrod"] = Colour{"palegoldenrod", 238, 232, 170}
	m["palegreen"] = Colour{"palegreen", 152, 251, 152}
	m["paleturquoise"] = Colour{"paleturquoise", 175, 238, 238}
	m["palevioletred"] = Colour{"palevioletred", 219, 112, 147}
	m["papayawhip"] = Colour{"papayawhip", 255, 239, 213}
	m["peachpuff"] = Colour{"peachpuff", 255, 218, 185}
	m["peru"] = Colour{"peru", 205, 133, 63}
	m["pink"] = Colour{"pink", 255, 192, 203}
	m["plum"] = Colour{"plum", 221, 160, 221}
	m["powderblue"] = Colour{"powderblue", 176, 224, 230}
	m["purple"] = Colour{"purple", 128, 0, 128}
	m["red"] = Colour{"red", 255, 0, 0}
	m["rosybrown"] = Colour{"rosybrown", 188, 143, 143}
	m["royalblue"] = Colour{"royalblue", 65, 105, 225}
	m["saddlebrown"] = Colour{"saddlebrown", 139, 69, 19}
	m["salmon"] = Colour{"salmon", 250, 128, 114}
	m["sandybrown"] = Colour{"sandybrown", 244, 164, 96}
	m["seagreen"] = Colour{"seagreen", 46, 139, 87}
	m["seashell"] = Colour{"seashell", 255, 245, 238}
	m["sienna"] = Colour{"sienna", 160, 82, 45}
	m["silver"] = Colour{"silver", 192, 192, 192}
	m["skyblue"] = Colour{"skyblue", 135, 206, 235}
	m["slateblue"] = Colour{"slateblue", 106, 90, 205}
	m["slategray"] = Colour{"slategray", 112, 128, 144}
	m["slategrey"] = Colour{"slategrey", 112, 128, 144}
	m["snow"] = Colour{"snow", 255, 250, 250}
	m["springgreen"] = Colour{"springgreen", 0, 255, 127}
	m["steelblue"] = Colour{"steelblue", 70, 130, 180}
	m["tan"] = Colour{"tan", 210, 180, 140}
	m["teal"] = Colour{"teal", 0, 128, 128}
	m["thistle"] = Colour{"thistle", 216, 191, 216}
	m["tomato"] = Colour{"tomato", 255, 99, 71}
	m["turquoise"] = Colour{"turquoise", 64, 224, 208}
	m["violet"] = Colour{"violet", 238, 130, 238}
	m["wheat"] = Colour{"wheat", 245, 222, 179}
	m["white"] = Colour{"white", 255, 255, 255}
	m["whitesmoke"] = Colour{"whitesmoke", 245, 245, 245}
	m["yellow"] = Colour{"yellow", 255, 255, 0}
	m["yellowgreen"] = Colour{"yellowgreen", 154, 205, 50}
	return m
}
