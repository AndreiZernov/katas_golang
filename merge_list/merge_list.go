package merge_list

func MergeList(list1, list2 []int64) []int64 {
	result := make([]int64, 0, len(list1)+len(list2))

	for {
		if len(list1) == 0 {
			return append(result, list2...)
		}
		if len(list2) == 0 {
			return append(result, list1...)
		}

		if len(list1) == 0 || list1[0] > list2[0] {
			result, list2 = PopFromList(result, list2)
		} else if len(list2) == 0 || list1[0] < list2[0] {
			result, list1 = PopFromList(result, list1)
		}
	}
}

func PopFromList(result, flist []int64) ([]int64, []int64) {
	result = append(result, flist[0])
	if len(flist) == 1 {
		flist = []int64{}
	} else {
		flist = flist[1:]
	}
	return result, flist
}
