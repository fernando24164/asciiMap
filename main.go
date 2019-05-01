package main

import mapObject "asciiMap/map"

func main() {
	mapObj := mapObject.New(20, 5)
	mapObj.FillMap()
	mapObj.Draw()
}
