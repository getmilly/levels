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

	result.Level = currentLevel
	result.TotalExperience = totalXP + earnedXP
	result.NextLevelExperience = calc.calculateExperienceByLevel(result.Level + 1)

	hasUpgraded := result.TotalExperience >= result.NextLevelExperience

	if hasUpgraded {
		result.Level = result.Level + 1
		result.TotalExperience = result.TotalExperience - result.NextLevelExperience
		result.NextLevelExperience = calc.calculateExperienceByLevel(result.Level + 1)

		if result.TotalExperience >= result.NextLevelExperience {
			return calc.Calculate(result.Level, result.TotalExperience, 0)
		}

		return result, nil
	}

	return result, nil
}

func (calc *TibiaCalculator) calculateExperienceByLevel(level int) int {
	return 50 / calc.difficulty * (level ^ 3 - 6*level ^ 2 + 17*level - 12)
}
