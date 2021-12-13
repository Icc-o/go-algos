package sort

//快排
func Qsort(arr []int) {
	if len(arr) < 2 {
		return
	}
	sortArr(arr, 0, len(arr)-1)
}

//1、快排首先选定一个用于比较的值
//2、把比这个值大的放到它的右边，比它小的放到左边
//3、继续排这个值的左边和右边的数组

//其中第二布使用双指针的方式，右指针先找比这个值小的位置，放到这个值的坑位（begin）
//左指针找比这个值大的位置，放到右指针的坑位
//如此循环，直到begin==end
//把这个用于比较的值放到begin位置，完成排序

func sortArr(arr []int, begin, end int) {
	i := begin
	j := end

	if begin >= end {
		return
	}

	var tmp = arr[begin]
	for begin < end {
		for ; begin < end && arr[end] > tmp; end-- {
		}
		arr[begin] = arr[end]

		for ; begin < end && arr[begin] < tmp; begin++ {
		}
		arr[end] = arr[begin]
	}

	arr[begin] = tmp

	// fmt.Println(arr)

	sortArr(arr, i, begin-1)
	sortArr(arr, begin+1, j)
}
