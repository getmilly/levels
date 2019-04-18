package level

//TibiaCalculator applies Tibia level formula
type TibiaCalculator struct {
	difficulty int
}

//NewTibiaCalculator creates a new PrtTo TibiaCalculator
func NewTibiaCalculator(difficulty int) *TibiaCalculator {
	if difficulty <= 0 {
		panic("difficulty must be gt zero")
	}
	return &TibiaCalculator{difficulty}
}

//Calculate player level.
func (calc *TibiaCalculator) Calculate(currentLevel, totalXP, earnedXP int) (*CalcResult, error) {
	result := new(CalcResult)

	result.Level = currentLevel
	result.TotalExperience = totalXP + earnedXP
	nextLevelExperience := calc.CalculateExperienceByLevel(result.Level + 1)
	result.HasUpgraded = result.TotalExperience >= nextLevelExperience

	if result.HasUpgraded {
		result.Level = result.Level + 1
		nextLevelExperience = calc.CalculateExperienceByLevel(result.Level + 1)

		if result.TotalExperience >= nextLevelExperience {
			return calc.Calculate(result.Level, result.TotalExperience, 0)
		}

		result.ExperienceToUpgrade = nextLevelExperience - result.TotalExperience

		return result, nil
	}

	result.ExperienceToUpgrade = nextLevelExperience - result.TotalExperience

	return result, nil
}

//CalculateExperienceByLevel ...
func (calc *TibiaCalculator) CalculateExperienceByLevel(level int) int {
	return 50 / calc.difficulty * (level ^ 3 - 6*level ^ 2 + 17*level - 12)
}
