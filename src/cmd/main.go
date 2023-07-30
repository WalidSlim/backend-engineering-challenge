package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"unbabel-challenge/internal/models"
	"unbabel-challenge/internal/movingaverage"
	"unbabel-challenge/internal/translation"
)

func main() {
	inputFile := flag.String("input_file", "events.json", "name of the file to calculate average durations on")
	windowSize := flag.Int("window_size", 10, "windowsize of the moving average")
	flag.Parse()

	//read and unmarshal events
	translation_file, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var translationList []models.Translation
	err = json.Unmarshal(translation_file, &translationList)
	if err != nil {
		log.Fatal(err)
	}

	//Start cleaning events by truncating time and keeping only the duration
	var truncatedTimeDurationList []models.TruncatedTimeAndDuration
	for i := range translationList {
		truncatedTimeDuration := translationList[i].FormatToTimeAndDuration()
		truncatedTimeDurationList = append(truncatedTimeDurationList, truncatedTimeDuration)
	}

	// Fill the list of truncated times with the times in between
	filledList := translation.FillTimeDurationList(truncatedTimeDurationList)

	//Instantiate moving average, loop over list of durations and calculate the moving average
	//using a sliding window that only takes account non zero values
	ma := movingaverage.NewMovingAverage(*windowSize)
	var resultList []models.MovingAverageResult
	for _, timeAndValue := range filledList {
		maValue := ma.AddValue(float64(timeAndValue.Duration))
		result := models.MovingAverageResult{Timestamp: timeAndValue.Timestamp, AverageDeliveryTime: maValue}
		resultList = append(resultList, result)
	}
	a, err := json.MarshalIndent(resultList, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("outputfile.json", a, 0644)
}
