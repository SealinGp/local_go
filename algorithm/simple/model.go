package simple

import "reflect"

type Ref struct {}

func inSlice(sliceElement interface{}, Slices interface{}) (int,bool) {
	index  := -1
	exists := false
	if reflect.TypeOf(Slices).Kind() != reflect.Slice {
		return index, false
	}

	value := reflect.ValueOf(Slices)
	for i := 0; i < value.Len(); i++ {
		if reflect.DeepEqual(sliceElement,value.Index(i).Interface()) {
			index  = i
			exists = true
			break
		}
	}
	return index,exists
}