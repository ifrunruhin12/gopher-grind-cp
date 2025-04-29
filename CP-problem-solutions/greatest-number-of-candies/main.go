//https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/
//Array

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	max := candies[0]
	for _, c := range candies {
		if c > max {
			max = c
		}
	}
    
	for i := range candies {
		if candies[i]+extraCandies >= max {
			res[i] = true
		}
	}

	return res
}