package simple
type Ref struct {}

func in_slice(sliceElement interface{}, intSlice interface{}) (int, bool) {
	find := false
	findIndex := 0
	for index, element := range intSlice.([]interface{}) {
		if element == sliceElement {
			find = true
			findIndex = index
			break
		}
	}
	return findIndex, find
}