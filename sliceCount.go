package main

import "fmt"

func main() {

	count := 1
	content := []string{"实用", "专利", "专利", "专利", "专利", "专利", "专利", "实用", "实用", "实用"}
	for i := 0; i < len(content); i++ {
		if content[i] == "replace" {
			continue
		}

		for j := i + 1; j < len(content); j++ {
			if content[i] == content[j] {

				count += 1
				rear := append([]string{}, content[j+1:]...)
				content = append(content[0:j], "replace")
				content = append(content, rear...)

			}
		}

		fmt.Printf("%s %d \n", content[i], count)

		count = 1

	}

}
