package main

// First solution
//func main() {
//	t := time.Now()
//	ourStr := ""
//	highestLen := 0
//
//	for charIndex, char := range ourStr {
//		// initial the current loop
//		checkMap := map[string]bool{string(char): false}
//		currentLen := 1
//		//if currentLen > highestLen {
//		//	highestLen = currentLen
//		//}
//
//		for index := charIndex + 1; index < len(ourStr); index++ {
//			checkStr := string(ourStr[index])
//			_, ok := checkMap[checkStr]
//			if ok {
//				if currentLen > highestLen {
//					highestLen = currentLen
//				}
//				break
//			}
//			checkMap[checkStr] = false
//			currentLen++
//			if currentLen > highestLen {
//				highestLen = currentLen
//			}
//			fmt.Println(checkMap)
//		}
//
//	}
//	fmt.Println(time.Since(t))
//	fmt.Println(highestLen)
//}

// second solution
//func main() {
//	t := time.Now()
//	ourStr := "dvdf"
//	highestLen := 0
//	currentLen := 0
//
//	// map to save the character and its index
//	checkMap := make(map[string]int)
//
//	for index := 0; index <= len(ourStr)-1; {
//		character := string(ourStr[index])
//		charIndex, ok := checkMap[character]
//		if !ok {
//			checkMap[character] = index
//			currentLen++
//			index++
//			if currentLen > highestLen {
//				highestLen = currentLen
//			}
//		} else {
//			if currentLen > highestLen {
//				highestLen = currentLen
//			}
//			index = charIndex + 1
//			// empty the map for the new check
//			checkMap = make(map[string]int)
//			currentLen = 0
//		}
//
//	}
//	fmt.Println(time.Since(t))
//}
