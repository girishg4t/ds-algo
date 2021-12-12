package main

import (
	"fmt"
	"os"

	icpatest "github.com/girishg4t/blind-50/ICPATEst"
	canfinishcourseschedule "github.com/girishg4t/blind-50/canFinishcourseschedule"
	canjump "github.com/girishg4t/blind-50/canJump"
	characterreplacement "github.com/girishg4t/blind-50/characterReplacement"
	combinationsum4 "github.com/girishg4t/blind-50/combinationSum4"
	containerwithwater "github.com/girishg4t/blind-50/containerWithWater"
	countsubstrings "github.com/girishg4t/blind-50/countSubstrings"
	"github.com/girishg4t/blind-50/divbuy3"
	eraseoverlapintervals "github.com/girishg4t/blind-50/eraseOverlapIntervals"
	findmin "github.com/girishg4t/blind-50/findMin"
	"github.com/girishg4t/blind-50/findword"
	"github.com/girishg4t/blind-50/findwords2"
	groupanagram "github.com/girishg4t/blind-50/groupAnagram"
	"github.com/girishg4t/blind-50/insertintervals"
	lengthoflis "github.com/girishg4t/blind-50/lengthOfLIS"
	longestconsecutive "github.com/girishg4t/blind-50/longestConsecutive"
	"github.com/girishg4t/blind-50/longestsubstring"
	mergeinterval "github.com/girishg4t/blind-50/mergeInterval"
	"github.com/girishg4t/blind-50/minmumwindowsubstring"
	numdecodings "github.com/girishg4t/blind-50/numDecodings"
	"github.com/girishg4t/blind-50/numberofislands"
	pacificatlantic "github.com/girishg4t/blind-50/pacificAtlantic"
	"github.com/girishg4t/blind-50/rainwater"
	"github.com/girishg4t/blind-50/twosum"
	uniquepaths "github.com/girishg4t/blind-50/uniquePaths"
	wordbreak "github.com/girishg4t/blind-50/wordBreak"
)

func main() {
	arg := os.Args[1:]
	switch arg[0] {
	case "1":
		fmt.Println(twosum.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	case "2":
		fmt.Println(mergeinterval.Merge([][]int{
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		}))
		fmt.Println(mergeinterval.Merge([][]int{
			{1, 3},
			{2, 4},
			{5, 7},
			{6, 8},
		}))
		fmt.Println(mergeinterval.Merge([][]int{
			{1, 4},
			{0, 4},
		}))

		fmt.Println(mergeinterval.Merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
	case "3":
		fmt.Println(groupanagram.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	case "4":
		fmt.Println(containerwithwater.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
		fmt.Println(containerwithwater.MaxArea([]int{1, 2, 1}))
		fmt.Println(containerwithwater.MaxArea([]int{4, 3, 2, 1, 4}))
		fmt.Println(containerwithwater.MaxArea([]int{1, 1}))
		fmt.Println(containerwithwater.MaxArea([]int{2, 3, 4, 5, 18, 17, 6}))
	case "5":
		fmt.Println(findmin.FindMin([]int{3, 4, 5, 1, 2}))
		fmt.Println(findmin.FindMin([]int{4, 5, 6, 7, 0, 1, 2}))
		fmt.Println(findmin.FindMin([]int{11, 13, 15, 17}))
		fmt.Println(findmin.FindMin([]int{11, 13, 15, 17, 18, 19, 20, 21, 10}))
		fmt.Println(findmin.FindMin([]int{13, 15, 17, 18, 19, 20, 21, 10, 11}))
		fmt.Println(findmin.FindMin([]int{5, 1, 2, 3, 4}))
	case "6":
		// fmt.Println(characterreplacement.CharacterReplacement("AAAB", 0))
		// fmt.Println(characterreplacement.CharacterReplacement("BAAA", 0))
		// fmt.Println(characterreplacement.CharacterReplacement("AABABCC", 2))
		// fmt.Println(characterreplacement.CharacterReplacement("ABBB", 2))
		fmt.Println(characterreplacement.CharacterReplacement("AABABBA", 1))
	case "7":
		fmt.Println(longestsubstring.LengthOfLongestSubstring("abcabcbb"))
		fmt.Println(longestsubstring.LengthOfLongestSubstring("bbbbb"))
		fmt.Println(longestsubstring.LengthOfLongestSubstring("pwwkew"))
		fmt.Println(longestsubstring.LengthOfLongestSubstring(""))
		fmt.Println(longestsubstring.LengthOfLongestSubstring("dvdf"))
		fmt.Println(longestsubstring.LengthOfLongestSubstring("abba"))
		fmt.Println(longestsubstring.LengthOfLongestSubstring(" "))
		fmt.Println(longestsubstring.LengthOfLongestSubstring("au"))
		fmt.Println(longestsubstring.LengthOfLongestSubstring("anviaj"))
	case "8":
		fmt.Println(minmumwindowsubstring.MinWindow("ADOBECODEBANC", "ABC"))
		fmt.Println(minmumwindowsubstring.MinWindow("geeksforgeeks", "ork"))
		fmt.Println(minmumwindowsubstring.MinWindow("a", "a"))
		fmt.Println(minmumwindowsubstring.MinWindow("a", "aa"))
		fmt.Println(minmumwindowsubstring.MinWindow("ADOBECODEBANCABC", "ABC"))
		fmt.Println(minmumwindowsubstring.MinWindow("a", "b"))
	case "9":
		fmt.Println(numberofislands.NumIslands([][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '1', '0', '0', '1'},
			[]byte{'1', '0', '0', '1', '1'},
			[]byte{'0', '0', '0', '0', '0'},
			[]byte{'1', '0', '1', '0', '1'},
		}))
		fmt.Println(numberofislands.NumIslands([][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1'},
		}))

		fmt.Println(numberofislands.NumIslands([][]byte{
			[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '0'},
		}))
		fmt.Println(numberofislands.NumIslands([][]byte{
			[]byte{'0', '1', '0'},
			[]byte{'1', '0', '1'},
			[]byte{'0', '1', '0'},
		}))
		fmt.Println(numberofislands.NumIslands([][]byte{
			[]byte{'1', '1', '1'},
			[]byte{'0', '1', '0'},
			[]byte{'1', '1', '1'},
		}))
	case "10":
		fmt.Println(countsubstrings.CountSubstrings("abaab"))
		fmt.Println(countsubstrings.CountSubstrings("abbaeae"))
		fmt.Println(countsubstrings.CountSubstrings("abc"))
		fmt.Println(countsubstrings.CountSubstrings("aaa"))
		fmt.Println(countsubstrings.CountSubstrings("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	case "11":
		fmt.Println(pacificatlantic.PacificAtlantic([][]int{
			[]int{1, 2, 2, 3, 5},
			[]int{3, 2, 3, 4, 4},
			[]int{2, 4, 5, 3, 1},
			[]int{6, 7, 1, 4, 5},
			[]int{5, 1, 1, 2, 4},
		}))
	case "13":
		fmt.Println(eraseoverlapintervals.EraseOverlapIntervals([][]int{
			[]int{1, 100},
			[]int{11, 22},
			[]int{1, 11},
			[]int{2, 12},
		}))
		fmt.Println(eraseoverlapintervals.EraseOverlapIntervals([][]int{
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 4},
			[]int{1, 3},
		}))
		fmt.Println(eraseoverlapintervals.EraseOverlapIntervals([][]int{
			[]int{1, 2},
			[]int{1, 2},
			[]int{1, 2},
		}))
		fmt.Println(eraseoverlapintervals.EraseOverlapIntervals([][]int{
			[]int{0, 2},
			[]int{1, 3},
			[]int{2, 4},
			[]int{3, 5},
			[]int{4, 6},
		}))
	case "16":
		// fmt.Println(divbuy3.Solution("23"))
		// fmt.Println(divbuy3.Solution("0081"))
		// fmt.Println(divbuy3.Solution("022"))
		// fmt.Println(divbuy3.Solution("0"))
		fmt.Println(divbuy3.Solution("100000"))
	case "17":
		// fmt.Println(canfinishcourseschedule.CanFinish(2, [][]int{
		// 	[]int{0, 1},
		// 	[]int{0, 2},
		// 	[]int{1, 3},
		// 	[]int{1, 4},
		// 	[]int{3, 4},
		// }))
		// fmt.Println(canfinishcourseschedule.CanFinish(3, [][]int{
		// 	[]int{0, 1},
		// 	[]int{1, 2},
		// 	[]int{2, 0},
		// }))
		// fmt.Println(canfinishcourseschedule.CanFinish(20, [][]int{
		// 	[]int{0, 10},
		// 	[]int{3, 18},
		// 	[]int{5, 5},
		// 	[]int{6, 11},
		// 	[]int{11, 14},
		// 	[]int{13, 1},
		// 	[]int{15, 1},
		// 	[]int{17, 4},
		// }))
		// fmt.Println(canfinishcourseschedule.CanFinish(2, [][]int{
		// 	[]int{0, 1},
		// }))
		// fmt.Println(canfinishcourseschedule.CanFinish(2, [][]int{
		// 	[]int{1, 0},
		// 	[]int{0, 1},
		// }))
		// fmt.Println(canfinishcourseschedule.CanFinish(2, [][]int{
		// 	[]int{1, 0},
		// }))
		fmt.Println(canfinishcourseschedule.CanFinish(2, [][]int{
			[]int{0, 10},
			[]int{3, 18},
			[]int{5, 5},
			[]int{6, 11},
			[]int{11, 14},
			[]int{13, 1},
			[]int{15, 1},
			[]int{17, 4},
		}))
	case "19":
		wr := findword.Constructor()
		wr.AddWord("bad")
		wr.AddWord("dad")
		wr.AddWord("mad")
		fmt.Println(wr.Search("pad"))
		fmt.Println(wr.Search("bad"))
		fmt.Println(wr.Search(".ad"))
		fmt.Println(wr.Search("b.."))
	case "22":
		// fmt.Println(insertintervals.Insert([][]int{
		// 	[]int{1, 3},
		// 	[]int{6, 9}}, []int{2, 5}))
		// fmt.Println(insertintervals.Insert([][]int{
		// 	[]int{1, 2},
		// 	[]int{3, 5},
		// 	[]int{6, 7},
		// 	[]int{8, 10},
		// 	[]int{12, 16},
		// }, []int{4, 8}))
		// fmt.Println(insertintervals.Insert([][]int{}, []int{5, 7}))
		// fmt.Println(insertintervals.Insert([][]int{
		// 	[]int{1, 5}}, []int{2, 3}))
		// fmt.Println(insertintervals.Insert([][]int{
		// 	[]int{1, 5}}, []int{2, 7}))
		fmt.Println(insertintervals.Insert([][]int{
			[]int{2, 5},
			[]int{6, 7},
			[]int{8, 9}}, []int{0, 1}))
	case "23":
		fmt.Println(longestconsecutive.LongestConsecutive([]int{
			100, 4, 200, 1, 3, 2,
		}))
		fmt.Println(longestconsecutive.LongestConsecutive([]int{
			0, 3, 7, 2, 5, 8, 4, 6, 0, 1,
		}))
		fmt.Println(longestconsecutive.LongestConsecutive([]int{
			0, 3, 7, 8, 9, 10, 11, 12,
		}))
		fmt.Println(longestconsecutive.LongestConsecutive([]int{
			0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999,
		}))
	case "24":
		fmt.Println(findwords2.FindWords([][]byte{
			[]byte{'o', 'a', 'a', 'n'},
			[]byte{'e', 't', 'a', 'e'},
			[]byte{'i', 'h', 'k', 'r'},
			[]byte{'i', 'f', 'i', 'v'},
		}, []string{
			"oath", "pea", "eat", "rain",
		}))

	case "25":
		fmt.Println(lengthoflis.LengthOfLIS([]int{
			10, 9, 2, 5, 3, 7, 101, 18,
		}))
		fmt.Println(lengthoflis.LengthOfLIS([]int{
			0, 1, 0, 3, 2, 3,
		}))
		fmt.Println(lengthoflis.LengthOfLIS([]int{
			7, 7, 7, 7, 7, 7, 7,
		}))
		fmt.Println(lengthoflis.LengthOfLIS([]int{
			10, 9, 2, 0, 3, 7, 1, 2,
		}))
		fmt.Println(lengthoflis.LengthOfLIS([]int{
			0, 3, 1, 6, 2, 2, 7,
		}))
	case "26":
		fmt.Println(combinationsum4.CombinationSum4([]int{
			1, 2, 3,
		}, 4))
	case "27":
		fmt.Println(icpatest.SolveTest("love", []string{"velo", "low", "vole"}))
	case "28":
		fmt.Println(numdecodings.NumDecodings("11106"))
		fmt.Println(numdecodings.NumDecodings("226"))
		fmt.Println(numdecodings.NumDecodings("12"))
	case "29":
		fmt.Println(uniquepaths.UniquePaths(2, 2))
		fmt.Println(uniquepaths.UniquePaths(3, 2))
		fmt.Println(uniquepaths.UniquePaths(3, 7))
		fmt.Println(uniquepaths.UniquePaths(23, 12))
	case "30":
		fmt.Println(canjump.CanJump([]int{2, 3, 1, 1, 4}))
		fmt.Println(canjump.CanJump([]int{3, 2, 1, 0, 4}))
		fmt.Println(canjump.CanJump([]int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9}))
		fmt.Println(canjump.CanJump([]int{3, 0, 8, 2, 0, 0, 1}))
		fmt.Println(canjump.CanJump([]int{1, 0, 1, 0}))
	case "31":
		// fmt.Println(wordbreak.WordBreak("leetcode", []string{"leet", "code"}))
		// fmt.Println(wordbreak.WordBreak("applepenapple", []string{"apple", "pen"}))
		// fmt.Println(wordbreak.WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
		// fmt.Println(wordbreak.WordBreak("cars", []string{"car", "ca", "rs"}))
		//fmt.Println(wordbreak.WordBreak("abcd", []string{"a", "abc", "b", "cd"}))
		fmt.Println(wordbreak.WordBreak("aaaaaaa", []string{"aaaa", "aaa"}))
		fmt.Println(wordbreak.WordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
		fmt.Println(wordbreak.WordBreak("aaaaaaaa", []string{"aaaa", "aaa", "aa"}))
		fmt.Println(wordbreak.WordBreak("aaaaaa", []string{"aa", "a"}))
	case "32":
		//fmt.Println(rainwater.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
		fmt.Println(rainwater.Trap([]int{4, 2, 0, 3, 2, 5}))
	}

}

//[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
