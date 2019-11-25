package main

import (
	"fmt"
	"regexp"
)

func main() {
	f4()
}

func f1() {
	sourceStr := `my email is gerrylon@163.com`
	matched, _ := regexp.MatchString(`[\w-]+@[\w]+(?:\.[\w]+)+`, sourceStr)
	fmt.Printf("%v", matched)
}

func f2() {
	var sourceStr string = `
test text     lljflsdfjdskal
gerrylon@163.com
abc@gmail.com
someone@sina.com.cn`

	re := regexp.MustCompile(`[\w-]+@([\w]+(?:\.[\w]+)+)`)
	matched := re.FindAllStringSubmatch(sourceStr, -1)
	for _, match := range matched {
		fmt.Printf("email is: %s, domain is: %s\n", match[0], match[1])
	}
}

func f3() {
	var sourceStr string = `gerrylon@163.com, abc@gmail.com, someone@sina.com.cn`
	re := regexp.MustCompile(`[\w-]+@([\w]+(?:\.[\w]+)+)`)
	matched := re.FindAllStringIndex(sourceStr, -1)
	for _, pos := range matched {
		fmt.Printf("start is: %v, end is: %v\n", pos[0], pos[1])
	}
}

func f4() {
	text := `Hello 世界！123 Go. 哈哈`
	reg := regexp.MustCompile(`[\p{Han}]+`)        // 查找连续的汉字
	fmt.Printf("%q\n", reg.FindAllString(text, 1)) // ["世界"]
	fmt.Println(reg.FindString(text))

}
