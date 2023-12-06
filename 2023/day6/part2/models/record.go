package models

type Record struct {
	Time     int
	Distance int
}

func (r Record) WaysToBeatUsing(boat Boat) int {
	holdingTimeToDistance := make(map[int]int)

	for i := 0; i < r.Time; i++ {
		holdingTime := i
		timeLeft := r.Time - holdingTime
		speedAfterHolding := holdingTime * boat.HoldingIncrease

		if distance := timeLeft * speedAfterHolding; distance > r.Distance {
			holdingTimeToDistance[holdingTime] = distance
		}
	}

	return len(holdingTimeToDistance)
}
