package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// i := findOccurences("XMAS")
	i := findOccurences2()
	fmt.Println("Occurences: ", i)
}

func findOccurences(word string)(count int) {
	if word == "" {
		return;
	}


	file, err := os.Open("./day4/input.txt");

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	wordLen := len(word)
	lineLen := len(lines[0])

	for i := 0; i < len(lines); i++ {
		for j := 0; j < lineLen; j++ {
			if lines[i][j] != word[0] {
				continue
			}
			
			// check left 
			if j >= wordLen - 1 {
				valid := true
				k := 1
				for k < wordLen {
					if lines[i][j-k] != word[k] {
						valid = false
						break
					}
					k++
				} 
				if valid {
					count++
				}
			}

			// check right
			if j <= lineLen - wordLen {
				valid := true
				k := 1
				for k < wordLen {
					if lines[i][j+k] != word[k] {
						valid = false
						break
					}
					k++
				} 
				if valid {
					count++
				}
			}

			// top  
			if i >= wordLen - 1{

				// top left
				if j >= wordLen - 1 {
					valid := true
					k := 1
					for k < wordLen {
						if lines[i-k][j-k] != word[k] {
							valid = false
							break
						}
						k++
					} 
					if valid {
						count++
					}
				}

				// top right
				if j <= lineLen - wordLen {
					valid := true
					k := 1
					for k < wordLen {
						if lines[i-k][j+k] != word[k] {
							valid = false
							break
						}
						k++
					} 
					if valid {
						count++
					}
				}
				
				// top top 
				valid := true
				k := 1
				for k < wordLen {
					if lines[i-k][j] != word[k] {
						valid = false
						break
					}
					k++
				} 
				if valid {
					count++
				}
			} 

			// bottom  
			if i <= len(lines) - wordLen {

				// bottom left
				if j >= wordLen - 1 {
					valid := true
					k := 1
					for k < wordLen {
						if lines[i+k][j-k] != word[k] {
							valid = false
							break
						}
						k++
					} 
					if valid {
						count++
					}
				}

				// bottom right
				if j <= lineLen - wordLen {
					valid := true
					k := 1
					for k < wordLen {
						if lines[i+k][j+k] != word[k] {
							valid = false
							break
						}
						k++
					} 
					if valid {
						count++
					}
				}
				
				// bottom bottom 
				valid := true
				k := 1
				for k < wordLen {
					if lines[i+k][j] != word[k] {
						valid = false
						break
					}
					k++
				} 
				if valid {
					count++
				}
			} 
		} 
	}
	return;
}

func findOccurences2()(count int) {
	file, err := os.Open("./day4/input.txt");

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 1; i < len(lines) - 1; i++ {
		for j := 1; j < len(lines[0]) - 1; j++ {
			if lines[i][j] != 'A' {
				continue
			}
			if ((lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') || (lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M')) && 
			((lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S') || (lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M')) {
				count++
			}  
		}
	}
	return
}