//The MIT License (MIT)
//Copyright (c) 2019 Hunter Tsai
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package kmp

// KMP search  string by kmp algorithm return true,start position if contains
func KMP(s, substr string) (bool, int) {
	if substr == "" {
		return true, 0
	}

	failureArr := failureArr(substr)
	return search(s, failureArr)
}

func search(s string, failureArr []failInfo) (bool, int) {
	byteArr := []rune(s)
	if len(byteArr) < len(failureArr) {
		return false, -1
	}

	i := 0
	j := 0
	for i < len(byteArr) && j < len(failureArr) {
		if byteArr[i] == failureArr[j].b {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = failureArr[j-1].value
			}
		}
	}

	if j < len(failureArr) {
		return false, -1
	}

	return true, i - len(failureArr)

}

func failureArr(s string) []failInfo {
	a := []failInfo{}
	byteArr := []rune(s)
	for i := range byteArr {
		var info failInfo
		value := 0

		preIndex := i - 1
		nowIndex := i

		tmp := i
		for nowIndex >= 0 && preIndex >= 0 {

			if byteArr[nowIndex] == byteArr[preIndex] {
				value++
				nowIndex--
			} else {
				value = 0
				nowIndex = tmp
			}
			preIndex--

		}
		info.b = byteArr[i]
		info.value = value
		a = append(a, info)

	}
	return a
}

type failInfo struct {
	b     rune
	value int
}
