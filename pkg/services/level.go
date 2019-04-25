package services

import (
	"time"

	gnats "github.com/getmilly/grok/nats"
	"github.com/getmilly/levels/pkg/level"
	"github.com/getmilly/levels/pkg/messages"
	"github.com/getmilly/levels/pkg/models"
	"github.com/getmilly/levels/pkg/repositories"
)

//LevelService ...
type LevelService interface {
	HandleReward(event *messages.RewardReached) error
}

type levelService struct {
	experienceRewardKind string
	levelUpgradedSubject string
	levelUpdatedSubject  string
	producer             *gnats.Producer
	calculator           level.Calculator
	experienceRepository repositories.ExperienceHistory
	calculatedRepository repositories.CalculatedLevelRepository
}

//NewLevelService ...
func NewLevelService(
	levelUpgradedSubject string,
	levelUpdatedSubject string,
	producer *gnats.Producer,
	calculator level.Calculator,
	experienceRepository repositories.ExperienceHistory,
	calculatedRepository repositories.CalculatedLevelRepository,
) LevelService {
	return &levelService{
		producer:             producer,
		calculator:           calculator,
		levelUpgradedSubject: levelUpgradedSubject,
		levelUpdatedSubject:  levelUpdatedSubject,
		experienceRepository: experienceRepository,
		calculatedRepository: calculatedRepository,
		experienceRewardKind: "xp",
	}
}

func (service *levelService) HandleReward(event *messages.RewardReached) error {
	if event.Kind != service.experienceRewardKind {
		return nil
	}

	if err := service.experienceRepository.Add(&models.Experience{
		CreatedAt:  time.Now(),
		Experience: event.Data.Experience,
		GameID:     event.GameID,
		PlayerID:   event.PlayerID,
		Reason:     event.Reason,
	}); err != nil {
		return err
	}

	sumExperience, err := service.experienceRepository.Sum(event.PlayerID)

	if err != nil {
		return err
	}

	result, _ := service.calculator.Calculate(sumExperience.TotalExperience)

	calculated, err := service.calculatedRepository.Get(event.PlayerID)

	if err != nil {
		return err
	}

	hasUpgraded := result.Level > calculated.Level

	calculated.Level = result.Level
	calculated.PlayerID = event.PlayerID
	calculated.TotalExperience = result.TotalExperience
	calculated.ExperienceToUpgrade = result.ExperienceToUpgrade

	err = service.calculatedRepository.Set(calculated)

	if err != nil {
		return err
	}

	return service.dispatchEvents(hasUpgraded, calculated)
}

func (service *levelService) dispatchEvents(hasUpgraded bool, calculated *models.CalculatedLevel) error {
	if hasUpgraded {
		return service.publish(service.levelUpgradedSubject, messages.LevelUpgraded{
			CalculatedLevel: calculated,
		})
	}

	return service.publish(service.levelUpdatedSubject, messages.LevelUpdated{
		CalculatedLevel: calculated,
	})
}

func (service *levelService) publish(subject string, m interface{}) error {
	message, err := gnats.NewMessage(m)

	if err != nil {
		return err
	}

	return service.producer.Publish(subject, message)
}
