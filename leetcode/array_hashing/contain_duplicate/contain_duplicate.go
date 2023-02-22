package contain_duplicate

// 重複していたら、trueを返す
// [1, 5, 7, 1]
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		// 重複していなかったらtrueにする
		if _, ok := m[v]; !ok {
			m[v] = true
		} else {
			// 既にある = 重複している
			return true
		}
	}
	return false
}
