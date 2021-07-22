package utils

// Deduplicate de-duplication.
func Deduplicate(list []uint16) []uint16 {
	set := make(map[uint16]struct{}, len(list))
	j := 0
	for _, v := range list {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		list[j] = v
		j++
	}

	return list[:j]
}
