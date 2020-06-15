package main

import "fmt"

//	选择排序返回数组中的最大值
func selectMax(arr []int) int {
	length := len(arr) //	取得数组长度
	//	如果数组的长度为1 则直接返回数组的0 下标
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]                 //	设定数组的0 下标就是数组的最大值
		for i := 1; i < length; i++ { //	循环获取数组中的最大值将其赋值给 arr[0]
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

//	选择排序
func selectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		//	数组长度小于等于 1 无需进行排序直接返回
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			min := i //	声明一个存储最小值的变量
			for j := i + 1; j < length; j++ {
				if arr[min] > arr[j] {
					min = j //	如果 min 的值比 j 的值大，则把小的值赋给 min
				}
			}
			if i != min {
				//	这里需要注意如果当前的最小值不是 i 则需要将 min 的值和 i 的值互换位置
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
		return arr
	}
}

//	堆排序获取最大值
func heapMax(arr []int, length int) []int {
	// length = len(arr)
	if length <= 1 {
		return arr
	} else {
		//	先确定二叉树的深度[深度就是有多少个二叉树的意思]
		depth := length/2 - 1
		for i := depth; i >= 0; i-- {
			//	倒序最后让最大的值浮起来到键值最小的
			topMax := i           //	假定最大的值是在 i 节点上的
			leftChild := 2*i + 1  //	此时的左子节点的值
			rightChild := 2*i + 2 //	此时右节点的值
			if leftChild <= length-1 && arr[leftChild] > arr[topMax] {
				//	左节点不能越界[最大只能到数组最大值减一]
				//	左节点如果大于最大值将其互换
				topMax = leftChild
			}
			if rightChild <= length-1 && arr[rightChild] > arr[topMax] {
				//	右节点不能越界[最大只能到数组最大值减一]
				//	右节点如果大于最大值将其互换
				topMax = rightChild
			}

			if topMax != i {
				//	如果最大值不在 i 上将其位置互换
				arr[topMax], arr[i] = arr[i], arr[topMax]
			}
		}
		return arr
	}
}

//	堆排序
func heapSort(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			lastLen := length - i
			heapMax(arr, lastLen)
			if i < length {
				arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
			}
		}
		return arr
	}
}

//	冒泡排序返回数组中的最大值
func popMax(arr []int, length int) int {
	for j := 1; j < length; j++ {
		if arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}

	return arr[length-1]
}

//	冒泡数组排序
//		冒泡排序多用于已有序列的数组新插入一个数据后迅速将其进行重新排序
func popSort(arr []int, length int) []int {
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//	每次根据步长对原数组进行重新排序
func shellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start; i < length; i += gap {
		backup := arr[i] //	对插入值进行备份
		j := i - gap     //	上一个循环的位置就是插值位置
		fmt.Println("当前起始位置是", i, "，操作位置是 ", j)
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] //	大值上浮
			j -= gap            //	跳过当前步长
		}
		arr[j+gap] = backup //	插备份值
		// fmt.Println(arr)
	}
}

//	希尔排序
//		最常见最好用的多线程快速排序
//		原理就是将数组的长度截一半做为步长将数组进行子数组拆分对子数组进行排序然后对步长进行递减继续拆分数组
//		直到步长为0
func shellSort(arr []int, length int) []int {
	if length <= 1 {
		return arr
	} else {
		gap := length / 2 //	步长初始值一般是数组除以二的长度
		for gap > 0 {
			//	只要 gap 大于 0 就一直进行除 2 拆分
			for i := 0; i < gap; i++ {
				shellSortStep(arr, i, gap)
			}
			gap /= 2
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0, 31, 12, 13, 98, 21} //	简单的数组最大值排序
	length := len(arr)
	// fmt.Println("选择排序返回数组中的最大值")
	// fmt.Println(selectMax(arr))
	// fmt.Println("选择排序")
	// fmt.Println(selectSort(arr))
	// fmt.Println("堆排序快速获取最大值")
	// fmt.Println(heapMax(arr, length))
	// fmt.Println("堆排序")
	// fmt.Println(heapSort(arr))
	// fmt.Println("冒泡返回数组的最大值")
	// fmt.Println(popMax(arr, length))
	// fmt.Println("冒泡排序")
	// fmt.Println(popSort(arr, length))
	fmt.Println("希尔排序")
	fmt.Println(shellSort(arr, length))
}
