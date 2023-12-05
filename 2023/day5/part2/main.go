package part2

import (
	"2023/day5/part2/models"
	"2023/day5/part2/parsers"
	"2023/day5/part2/utils"
	"fmt"
	"sync"
)

func Run(input []string) int {
	var results []int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	seeds, categoryMappings := parsers.Parse(input)

	for _, seed := range seeds {
		wg.Add(1)
		go func(s models.Seeds) {
			defer wg.Done()

			result := s.FindLocations(categoryMappings)

			mutex.Lock()
			results = append(results, result...)
			fmt.Print(".")
			mutex.Unlock()
		}(seed)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return utils.FindMinInSlice(results)

}
