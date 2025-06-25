package leetcode

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	// 测试用例
	// 第一题：尽量均分数组。
	fmt.Println(OneProblems([]int{1, 2, 3, 4, 5, 6}, 6))
	fmt.Println(OneProblems([]int{1, 2, 3, 4, 5, 6}, 7))
	fmt.Println(OneProblems([]int{1, 2, 3, 4, 5, 6}, 1))
	fmt.Println(OneProblems([]int{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println(OneProblems([]int{1, 2, 3, 4, 5, 6, 7}, 3))

	// 第二题：
	// 先查询项目表和绩效表作为ps。然后使用员工表和其关联查询做统计即可。
	// select employeeId,employeeName,count(projectId),avg(score)
	// from employee left join
	// (select projectId,score from project left join score on projectId = score_project_id where project.score > 4 as ps)
	// on employee.projectid = ps.id group by employeeId
}

func OneProblems(dates []int, parts int) [][]int {
	result := [][]int{}
	// parts 为 1，返回二维数组
	if parts == 1 {
		return append(result, dates)
	}

	// 获取每份长度
	length := 0
	if len(dates) < parts {
		length = 1
	} else {
		length = len(dates) / parts
	}
	// 获取最终余数
	multiPart := len(dates) % length

	// 此处循环即可，最后一个一定无法满足。最后添加即可。
	for i := 0; i <= len(dates)-1-multiPart; i += length {
		result = append(result, dates[i:i+length])
	}
	// 如果余数为零，刚好分配完。
	if multiPart == 0 {
		return result
	}
	// 补充最后可能遗留的数
	result[len(result)-1] = append(result[len(result)-1], dates[len(dates)-multiPart:]...)
	return result
}
