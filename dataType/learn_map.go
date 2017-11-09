/**
 * @Author : Yannick
 * @File   : learn_map.go
 * @Date   : 2017-11-08
 * @Desc   : This is a demo which I could learn golang through it.
 */

package main

import (
	"fmt"
)

// LearnMap :
func LearnMap() {
	fmt.Println("## LearnMap() called begin ..")

	// create a map called colors, type(key) is string, type(value) is string
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	colors["Blue"] = "#dx1023"
	fmt.Printf("colors => %v\n\n", colors)

	// create a nil map without assiagn value when created it.
	// httpHeaders :=map[string]string will lead to a runtime error.
	var httpHeaders map[string]string
	fmt.Printf("httpHeaders => %v", httpHeaders)
	// httpHeaders["Method"] = "POST" // panic: assignment to entry in nil map
	// a nil map is a readonly map, always used as default argument of a function.

	// map[key] returns 2 values : key-value, bool
	yellowColor, isInMap := colors["Yellow"]

	/*
		 The below code will lead a compile error, which means else must fellow the } of a if stament, cannot
		 start from a new line.

			if isInMap {
				fmt.Printf("Yellow is in map: colors, and colors[\"Yellow\"] => %s\n", yellowColor)
			}
			else {
				fmt.Println("Yellow is not in map: colors")
			}
	*/
	if isInMap {
		fmt.Printf("Yellow is in map: colors, and colors[\"Yellow\"] => %s\n", yellowColor)
	} else {
		fmt.Println("Yellow is not in map: colors")
	}

	// loop a map
	webColors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range webColors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}

	// remove an element from a map
	delete(webColors, "Coral")
	fmt.Println("after delete(webColors, \"Coral\", the webColors is:")
	for key, value := range webColors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}

	// delete a not exist key-value
	delete(webColors, "blue")

	// map used as function arguments

	fmt.Println("## LearnMap() called end ..")
}
