package main

func Search(nums []int, target int) int {
    left :=0 //smallest index target could be
    right := len(nums)-1 //largest index target could be

    for left+1<right {
        mid := left+(right-left)/2 //find middle of possible range
        if nums[mid] == target { //if middle of list is target, return mid
            return mid
        } else if nums[mid] > target { //if target is lower than mid, cut off larger half
            right=mid
        } else { //if target is higher than mid, cut off lower half
            left=mid
        }
    }
    //once the number is found, return the index of the number
    if nums[left] == target {return left}
    if nums[right] ==target {return right}
    return -1
}

