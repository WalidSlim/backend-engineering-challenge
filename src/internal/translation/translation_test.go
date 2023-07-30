package translation_test

import (
	"testing"
	"time"
	"unbabel-challenge/internal/models"
	"unbabel-challenge/internal/translation"

	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	t.Run("Reading from a file should fill a Translation object successfully", func(t *testing.T) {
		expectedTimeValue, _ := time.Parse(time.DateTime, "2018-12-26 18:11:08.509654")
		expectedTime := models.Time{Time: expectedTimeValue}
		expected_translaiton := models.Translation{
			Timestamp:      expectedTime,
			TranslationId:  "5aa5b2f39f7254a75aa5",
			SourceLanguage: "en",
			TargetLanguage: "fr",
			ClientName:     "airliberty",
			EventName:      "translation_delivered",
			NrWords:        30,
			Duration:       20,
		}
		var translation_from_file models.Translation
		translation_from_file, err := translation.ReadTranslationJson("input.json")
		assert.NoError(t, err)
		assert.True(t, assert.ObjectsAreEqualValues(translation_from_file, expected_translaiton))
	})
	t.Run("Reading from a file should fill a Translation list successfully", func(t *testing.T) {
		expectedTimeValue, _ := time.Parse(time.DateTime, "2018-12-26 18:11:08.509654")
		expectedTime := models.Time{Time: expectedTimeValue}
		expected_translaiton := models.Translation{
			Timestamp:      expectedTime,
			TranslationId:  "5aa5b2f39f7254a75aa5",
			SourceLanguage: "en",
			TargetLanguage: "fr",
			ClientName:     "airliberty",
			EventName:      "translation_delivered",
			NrWords:        30,
			Duration:       20,
		}
		var translation_from_file []models.Translation
		translation_from_file, err := translation.ReadTranslationList("input_list.json")
		assert.NoError(t, err)
		assert.True(t, assert.ObjectsAreEqualValues(translation_from_file[0], expected_translaiton))
	})
}
