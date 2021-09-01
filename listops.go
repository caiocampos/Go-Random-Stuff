package listops

// binFunc type defines a function with two integers as input parameters and returns an integer
type binFunc (func(int, int) int)

// unaryFunc type defines a function with one integer as input parameter and returns an integer
type unaryFunc func(int) int

// binFunc type defines a function with one integer as input parameter and returns a boolean
type predFunc func(int) bool

// IntList type defines a List of integers
type IntList []int

// Append method appends a IntList to another
func (list IntList) Append(l IntList) IntList {
	return append(list, l...)
}

// Concat method concatenates multiple IntList
func (list IntList) Concat(l []IntList) IntList {
	res := list.Append(IntList{})
	for _, el := range l {
		res = res.Append(el)
	}
	return res
}

// Filter method filters the IntList
func (list IntList) Filter(fun predFunc) IntList {
	res := IntList{}
	for _, el := range list {
		if fun(el) {
			res = append(res, el)
		}
	}
	return res
}

// Foldl method folds the IntList (Left)
func (list IntList) Foldl(fun binFunc, initial int) int {
	length := list.Length()
	res := initial
	for i := 0; i < length; i++ {
		res = fun(res, list[i])
	}
	return res
}

// Foldr method folds the IntList (Right)
func (list IntList) Foldr(fun binFunc, initial int) int {
	length := list.Length()
	res := initial
	for i := length - 1; i >= 0; i-- {
		res = fun(list[i], res)
	}
	return res
}

// Length method returns the length of the IntList
func (list IntList) Length() int {
	return len(list)
}

// Map method applies a function to all elements of the IntList
func (list IntList) Map(fun unaryFunc) IntList {
	res := IntList{}
	for _, el := range list {
		res = append(res, fun(el))
	}
	return res
}

// Reverse method generates the reverse IntList
func (list IntList) Reverse() IntList {
	length := list.Length()
	res := IntList{}
	for i := length - 1; i >= 0; i-- {
		res = append(res, list[i])
	}
	return res
}
