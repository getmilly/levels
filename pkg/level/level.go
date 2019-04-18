package level

//Calculator wraps the logic behind calculate the player level.
type Calculator interface {
	Calculate(currentLevel, totalXP, earnedXP int) (*CalcResult, error)
}

//CalcResult has all information about a level calculation.
type CalcResult struct {
	TotalExperience     int
	HasUpgraded         bool
	NextLevelPercentage float64
	NextLevelExperience int
}
