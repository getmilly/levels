package level

//TibiaCalculator applies Tibia level formula
type TibiaCalculator struct {
	difficulty int
}

//NewTibiaCalculator creates a new PrtTo TibiaCalculator
func NewTibiaCalculator(difficulty int) Calculator {
	if difficulty <= 0 {
		panic("difficulty must be gt zero")
	}
	return &TibiaCalculator{difficulty}
}

//Calculate player level.
func (calc *TibiaCalculator) Calculate(currentLevel, totalXP, earnedXP int) (*CalcResult, error) {
	result := new(CalcResult)
	nextLevelExperience := calc.calculateExperienceByLevel(currentLevel + 1)

	result.TotalExperience = totalXP + earnedXP
	result.HasUpgraded = result.TotalExperience >= nextLevelExperience

	if result.HasUpgraded {
		result.NextLevelExperience = calc.calculateExperienceByLevel(currentLevel+2) - result.TotalExperience
	} else {
		result.NextLevelExperience = nextLevelExperience - result.TotalExperience
	}

	return result, nil
}

func (calc *TibiaCalculator) calculateExperienceByLevel(level int) int {
	return 50 / calc.difficulty * (level ^ 3 - 6*level ^ 2 + 17*level - 12)
}
