package main

func RemoveDuplicates(nums []int) int {
	prev := nums[0]
	i := 1
	for j := 1; j < len(nums); j++ { 
		if nums[j] != prev {
			nums[i] = nums[j]
			i++
		}

		prev = nums[j]
	}
	return i;
}

/*
	var nums = [8]int{
		0,1, 2, 3, 4, 5, 5, 4,
	}
	// var result = removeDuplicates(nums[:])

	// fmt.Println(result)

	prev := nums[0]
	i := 1

	fmt.Println("ANGKA PERTAMA : ", prev, i)

	fmt.Println("Panjang Nums : ", len(nums))
	for j := 1; j < len(nums); j++ {

		fmt.Println("====== Iterasi J", j, "=======")
		if nums[j] != prev { // Jika Nums Index J Nilanya Tidak Sama Dengan Prev Nilainya Maka Prev Isinya == nums[j]

			fmt.Println(" ")
			fmt.Println("Prev Increment", prev)
			fmt.Println(" ")
			fmt.Println(" ")
			nums[i] = nums[j] // Nums Dengan Index

			fmt.Println("Increment i :", i)
			fmt.Println("Variabel I : ", nums[i])
			fmt.Println("Variabel J : ", nums[j])

			i++ // Lalu I Increment
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Akhir Iterasi Ke-", prev)
			fmt.Println(" ")
			fmt.Println(" ")
		}

		prev = nums[j]
	}

	fmt.Println("Hasil :", i)
*/
