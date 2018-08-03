//custom package function names first letter needs to be capital letter for it to be accessible when imported

package flames

import (
	"github.com/novalagung/gubrak"
	"fmt"
	)

func FindUniqueLetters(args []string) []string {
	chars1 := make([]string, 0)
	chars2 := make([]string, 0)

	for _, name1 := range args[0] {
		chars1 = append(chars1, string(name1))
	}
	for _, name2 := range args[1] {
		chars2 = append(chars2, string(name2))
	}

	c1, _ := gubrak.Difference(chars1, chars2)
	c2, _ := gubrak.Difference(chars2, chars1)
	result, _ := gubrak.Union(c1, c2)

	return result.([]string)
}

func FindUniqueLettersOfThreeNames(args []string) []string {
	chars := FindUniqueLetters(args[:2])
	chars3 := make([]string, 0)
	//var diff []string

	for _, nameChar := range args[2] {
		chars3 = append(chars3, string(nameChar))
	}
	c, _ := gubrak.Difference(chars, chars3)
	c3, _ := gubrak.Difference(chars3, chars)
	result, _ := gubrak.Union(c, c3)
	fmt.Println(result.([]string))
	return result.([]string)

	//for i := 0; i < 2; i++ {
	//	for _, s1 := range chars1 {
	//		found := false
	//		for _, s2 := range chars2 {
	//			if s1 == s2 {
	//				found = true
	//				break
	//			}
	//		}
	//		if !found {
	//			diff = append(diff, s1)
	//		}
	//	}
	//	if i == 0 {
	//		chars1, chars2 = chars2, chars1
	//	}
	//}

	//return diff
}

func Flames(diff , names []string) string {
	flames := []string{"F", "L", "A", "M", "E", "S"}
	count := 0
	response := ""

	for len(flames) == 6 {
		for i := range diff {
			if i + 1 == 1 {
				if count == len(flames) {
					count = 0
				} else if count == 5{
					count = 5
				} else {
					count = count
				}
			} else if count == 5 {
				count = 0
			} else {
				count = count + 1
			}
		}
		copy(flames[count:], flames[count+1:])
		flames[len(flames)-1] = ""
		flames = flames[:len(flames)-1]
		fmt.Println(flames)
	}

	for len(flames) == 5 {
		for i := range diff {
			if i + 1 == 1 {
				if count == len(flames) {
					count = 0
				} else if count == 4{
					count = 4
				} else {
					count = count
				}
			} else if count == 4 {
				count = 0
			} else {
				count = count + 1
			}
		}
		copy(flames[count:], flames[count+1:])
		flames[len(flames)-1] = ""
		flames = flames[:len(flames)-1]
		fmt.Println(flames)
	}

	for len(flames) == 4 {
		for i := range diff {
			if i + 1 == 1 {
				if count == len(flames) {
					count = 0
				} else if count == 3{
					count = 3
				} else {
					count = count
				}
			} else if count == 3 {
				count = 0
			} else {
				count = count + 1
			}
		}
		copy(flames[count:], flames[count+1:])
		flames[len(flames)-1] = ""
		flames = flames[:len(flames)-1]
		fmt.Println(flames)
	}

	for len(flames) == 3 {
		for i := range diff {
			if i + 1 == 1 {
				if count == len(flames) {
					count = 0
				} else if count == 2{
					count = 2
				} else {
					count = count
				}
			} else if count == 2 {
				count = 0
			} else {
				count = count + 1
			}
		}
		copy(flames[count:], flames[count+1:])
		flames[len(flames)-1] = ""
		flames = flames[:len(flames)-1]
		fmt.Println(flames)
	}

	for len(flames) == 2 {
		for i := range diff {
			if i + 1 == 1 {
				if count == len(flames) {
					count = 0
				} else if count == 1{
					count = 1
				} else {
					count = count
				}
			} else if count == 1 {
				count = 0
			} else {
				count = count + 1
			}
		}
		copy(flames[count:], flames[count+1:])
		flames[len(flames)-1] = ""
		flames = flames[:len(flames)-1]
		fmt.Println(flames)
	}

	acronyms := map[string]string{
		"F" : "Friends with benefits",
		"L" : "Lovers in bluefrog",
		"A" : "Adto sa kalibonan mag date",
		"M" : "Magsigeg sungog2x hantud ma fall",
		"E" : "Eut ra tirada",
		"S" : "Sex Buddies",
	}

	if len(names) == 2 {
		response = fmt.Sprintf("Si %v ug %v kay %v", names[0], names[1], acronyms[flames[0]])
	} else {
		response = fmt.Sprintf("Si %v, %v ug %v kay %v", names[0], names[1], names[2], acronyms[flames[0]])
	}

	return response
}