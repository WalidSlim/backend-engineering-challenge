package translation

import (
	"encoding/json"
	"io/ioutil"
	"time"
	"unbabel-challenge/internal/models"
)

func ReadTranslationJson(filename string) (models.Translation, error) {
	translation_file, _ := ioutil.ReadFile(filename)
	var translation models.Translation
	err := json.Unmarshal(translation_file, &translation)

	return translation, err
}

func ReadTranslationList(filename string) ([]models.Translation, error) {
	translation_file, _ := ioutil.ReadFile(filename)
	var translation []models.Translation
	err := json.Unmarshal(translation_file, &translation)

	return translation, err
}

//The idea behind this is to get to the right data structure, going from continuous time to discrete
//and only having the duration as everything else is clutter for the scope of the exercise.
//We also pad the missing times by adding zeros, which we will later not take into account
//When Calculating the moving averages
func FillTimeDurationList(timeAndDurationList []models.TruncatedTimeAndDuration) []models.TruncatedTimeAndDuration {
	timeAndDurationListLength := len(timeAndDurationList)
	//Taking into account that the list is sorted
	minTime, maxTime := timeAndDurationList[0].Timestamp.Add(-1*time.Minute), timeAndDurationList[timeAndDurationListLength-1].Timestamp
	duration := maxTime.Sub(minTime)
	var filledTimeDurationList []models.TruncatedTimeAndDuration
	//Making a map to make a search O(1)
	timesMap := make(map[models.Time]bool)
	for _, times := range timeAndDurationList {
		timesMap[times.Timestamp] = false
	}
	timeAndDurationListIndex := 0
	//For each minute, we will take the duration value if it exists in the timesMap
	//Otherwise, we will add a 0 value to that time.
	for i := 0; i < int(duration.Minutes())+1; i++ {

		currentTime := models.Time{Time: minTime.Add(time.Duration(i) * time.Minute)}
		_, existingTime := timesMap[currentTime]

		if existingTime {
			currentTimeDuration := models.TruncatedTimeAndDuration{Timestamp: currentTime, Duration: timeAndDurationList[timeAndDurationListIndex].Duration}
			filledTimeDurationList = append(filledTimeDurationList, currentTimeDuration)
			timeAndDurationListIndex += 1
		} else {
			currentTimeDuration := models.TruncatedTimeAndDuration{Timestamp: currentTime, Duration: 0}
			filledTimeDurationList = append(filledTimeDurationList, currentTimeDuration)
		}
	}
	return filledTimeDurationList
}
