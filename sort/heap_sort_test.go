package sort

import (
	"fmt"
	"testing"
)

//func HeapSort(nums []int)  {
//
//	N := len(nums)-1
//	//从底部到顶部构建大顶堆，最后一个非叶子节点开始
//	for i := N/2-1; i >= 0; i--{   //减少1才是最后一个非叶子节点。
//		sink(nums,i,N)
//	}
//
//	//将堆顶值和末尾交换，重新调整堆
//	for i:=N;i>=0;i--{
//		nums[0],nums[i] = nums[i],nums[0]
//		sink(nums,0,i-1)//交换之后，数组最后一位不算在堆内，需要减1操作
//	}
//
//	//不同写法，结果一样
//	/*for N>=0{
//	    wap(nums,0,N)
//	    N--
//	    sink(nums,0,N)
//	}*/
//}

//func sink(nums []int,k,N int)  {
//	for{
//		i:=2*k+1
//		if i>N{
//			break
//		}
//		//找左右子节点最大值
//		if i+1<=N && nums[i+1]>nums[i]{
//			i++  //找到右边节点位置
//		}
//		//已经大于最大值，不需要再交换，此时i的位置是合适的，不要再sink了
//		if nums[k]>=nums[i]{
//			break
//		}
//		nums[k],nums[i] = nums[i],nums[k]
//		k = i //继续向上调整
//	}
//}

func HeapSort(nums []int) {
	N := len(nums) - 1
	for i := N/2 - 1; i >= 0; i-- {
		sink(nums, 0, N)
	}
	for i := N; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		sink(nums, 0, i-1)
	}
}

func sink(nums []int, k, N int) {
	for {
		i := 2*k + 1
		if i > N {
			break
		}

		if i+1 < N && nums[i] < nums[i+1] {
			i++
		}
		if nums[i] <= nums[k] {
			break
		}
		nums[i], nums[k] = nums[k], nums[i]
		k = i
	}
}
func TestHeapSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 8, 0, 6, 5, 6, 7, 8}
	fmt.Println(a)
	HeapSort(a)
	fmt.Println(a)
	//var b = []int{222,333,444,555}

}
