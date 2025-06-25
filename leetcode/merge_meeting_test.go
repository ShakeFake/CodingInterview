package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/*
问题：给定一个未排序，独立的会议列表，返回一个合并过的会议列表。
例子:    输入: meeting{{1, 2}, {2, 3}, {4, 5}}

	输出: meeting{{1, 3}, {4, 5}}
	输入: meeting{{1, 5}, {2, 3}}
	输出: meeting{{1, 5}}

方法：使用一个数组，用来接收返回值。

	将会议按照首部进行排序，然后比较尾部。
	遍历会议列表，比较尾部大小即可。
*/
func TestMergeMeeting(t *testing.T) {
	type Data struct {
		INPUT  []Meeting
		OUTPUT []Meeting
	}

	testResults := []Data{
		{[]Meeting{}, []Meeting{}},
		{[]Meeting{{1, 2}}, []Meeting{{1, 2}}},
		{[]Meeting{{1, 2}, {2, 3}}, []Meeting{{1, 3}}},
		{[]Meeting{{1, 2}, {3, 4}}, []Meeting{{1, 2}, {3, 4}}},
		{[]Meeting{{1, 2}, {2, 3}, {3, 4}}, []Meeting{{1, 3}}},
		{[]Meeting{{1, 2}, {2, 3}, {4, 5}}, []Meeting{{1, 3}, {4, 5}}},
		{[]Meeting{{1, 5}, {2, 3}, {4, 5}}, []Meeting{{1, 5}}},
		{[]Meeting{{2, 3}, {4, 5}, {1, 6}}, []Meeting{{1, 6}}},
	}

	for _, testResult := range testResults {
		mergeResult := MergeMeeting(testResult.INPUT)
		// todo : 写一个公共比较函数
		fmt.Println(mergeResult)
	}

}

func MergeMeeting(dates []Meeting) []Meeting {
	// 先进行数据的排序，然后再处理。使用sort进行排序。该排序的复杂的为 n^2
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Begin < dates[j].Begin
	})

	// 此时接收到的数据为无序的
	item := []Meeting{}

	for k := range dates {
		if k == 0 {
			item = append(item, dates[k])
			continue
		}

		// 如果是尾部大于下一个开始，即代表有交集。
		if item[len(item)-1].End >= dates[k].Begin {
			// 此处需要处理一下，修改成用函数进行比较。
			if item[len(item)-1].End < dates[k].End {
				item[len(item)-1].End = dates[k].End
			} else {
				item[len(item)-1].End = item[len(item)-1].End
			}
		} else {
			item = append(item, dates[k])
		}
	}
	return item
}

type Meeting struct {
	Begin int
	End   int
}
