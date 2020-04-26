package main

import "fmt"

func main() {

	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {

	if len(strs) == 0 {
		return nil
	}

	if len(strs) == 1 {
		return [][]string{strs}
	}

	helper := map[[256]int][]string{}

	for _, str := range strs {

		bitMap := [256]int{}
		for _, c := range str {
			bitMap[c-'a'] += 1
		}

		helper[bitMap] = append(helper[bitMap], str)
	}

	res := make([][]string, 0, len(helper))
	for _, v := range helper {
		res = append(res, v)
	}

	return res
}
