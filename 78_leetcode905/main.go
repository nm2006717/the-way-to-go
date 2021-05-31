package main

func  main()  {
	nums:=[]int{0,1,2,3,5,7}
	sortArrayByParity(nums)
}

func sortArrayByParity(A []int) []int {
	head:=0
	n:=len(A)
	tail:=n-1
	for {
		if A[head] %2 !=0 && A[tail]%2==0{
			A[head],A[tail] = A[tail],A[head]
		}
		if A[head] % 2==0 {
			head = head + 1
		}
		if A[tail] %2 !=0{
			tail = tail -1
		}
		if   (head >= n/2 && tail <= n/2) || head == n-1 || tail == 0{
			break
		}
	}
	return A
}