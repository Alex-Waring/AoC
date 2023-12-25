package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func cardCompare(card1 string, card2 string) bool {
	conversion_map := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}
	return conversion_map[card1] < conversion_map[card2]
}

func cardComparePart2(card1 string, card2 string) bool {
	conversion_map := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 1,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}
	return conversion_map[card1] < conversion_map[card2]
}

func handType(hand string) int {
	set := make(map[rune]int)
	for _, letter := range hand {
		val, ok := set[letter]
		if ok {
			set[letter] = val + 1
		} else {
			set[letter] = 1
		}
	}

	counts := []int{}
	for _, value := range set {
		counts = append(counts, value)
	}

	slices.Sort(counts)
	if reflect.DeepEqual(counts, []int{5}) {
		return 7
	} else if reflect.DeepEqual(counts, []int{1, 4}) {
		return 6
	} else if reflect.DeepEqual(counts, []int{2, 3}) {
		return 5
	} else if reflect.DeepEqual(counts, []int{1, 1, 3}) {
		return 4
	} else if reflect.DeepEqual(counts, []int{1, 2, 2}) {
		return 3
	} else if reflect.DeepEqual(counts, []int{1, 1, 1, 2}) {
		return 2
	} else {
		return 1
	}
}

func handTypePart2(hand string) int {
	set := make(map[rune]int)
	for _, letter := range hand {
		val, ok := set[letter]
		if ok {
			set[letter] = val + 1
		} else {
			set[letter] = 1
		}
	}

	counts := []int{}
	for _, value := range set {
		counts = append(counts, value)
	}
	slices.Sort(counts)

	count, joker := set['J']
	if joker {
		if count == 1 {
			if reflect.DeepEqual(counts, []int{1, 4}) {
				return 7
			} else if reflect.DeepEqual(counts, []int{1, 1, 3}) {
				return 6
			} else if reflect.DeepEqual(counts, []int{1, 2, 2}) {
				return 5
			} else if reflect.DeepEqual(counts, []int{1, 1, 1, 2}) {
				return 4
			} else {
				return 2
			}
		} else if count == 2 {
			if reflect.DeepEqual(counts, []int{2, 3}) {
				return 7
			} else if reflect.DeepEqual(counts, []int{1, 2, 2}) {
				return 6
			} else if reflect.DeepEqual(counts, []int{1, 1, 1, 2}) {
				return 4
			} else {
				return 1
			}
		} else if count == 3 {
			if reflect.DeepEqual(counts, []int{2, 3}) {
				return 7
			} else if reflect.DeepEqual(counts, []int{1, 1, 3}) {
				return 6
			} else {
				return 1
			}
		} else {
			return 7
		}
	} else {
		if reflect.DeepEqual(counts, []int{5}) {
			return 7
		} else if reflect.DeepEqual(counts, []int{1, 4}) {
			return 6
		} else if reflect.DeepEqual(counts, []int{2, 3}) {
			return 5
		} else if reflect.DeepEqual(counts, []int{1, 1, 3}) {
			return 4
		} else if reflect.DeepEqual(counts, []int{1, 2, 2}) {
			return 3
		} else if reflect.DeepEqual(counts, []int{1, 1, 1, 2}) {
			return 2
		} else {
			return 1
		}
	}
}

type Round struct {
	hand_type int
	cards     string
	bid       int
}

func main() {
	defer utils.Timer("main")
	lines := utils.ReadInput("input.txt")
	hands := []Round{}

	for _, line := range lines {
		hand := strings.Fields(line)[0]
		bid := strings.Fields(line)[1]
		hands = append(hands, Round{
			hand_type: handType(hand),
			cards:     hand,
			bid:       utils.IntegerOf(bid),
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].hand_type != hands[j].hand_type {
			return hands[i].hand_type < hands[j].hand_type
		}
		for k := 0; k < 5; k++ {
			if hands[i].cards[k] != hands[j].cards[k] {
				return cardCompare(string(hands[i].cards[k]), string(hands[j].cards[k]))
			}
		}
		return true
	})

	power := 0
	for i := 0; i < len(hands); i++ {
		power += hands[i].bid * (i + 1)
	}
	fmt.Println(power)

	hands_2 := []Round{}

	for _, line := range lines {
		hand := strings.Fields(line)[0]
		bid := strings.Fields(line)[1]
		hands_2 = append(hands_2, Round{
			hand_type: handTypePart2(hand),
			cards:     hand,
			bid:       utils.IntegerOf(bid),
		})
	}

	sort.Slice(hands_2, func(i, j int) bool {
		if hands_2[i].hand_type != hands_2[j].hand_type {
			return hands_2[i].hand_type < hands_2[j].hand_type
		}
		for k := 0; k < 5; k++ {
			if hands_2[i].cards[k] != hands_2[j].cards[k] {
				return cardComparePart2(string(hands_2[i].cards[k]), string(hands_2[j].cards[k]))
			}
		}
		return true
	})
	power = 0
	for i := 0; i < len(hands_2); i++ {
		power += hands_2[i].bid * (i + 1)
	}
	fmt.Println(power)
}
