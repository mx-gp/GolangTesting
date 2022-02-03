package main

import (
	"GolangTesting/helper"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadTxt struct {
	Text string `json:"text"`
}

type TopTen struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func wordCount(str string) []TopTen {

	// convert string to lowercase
	lowerCaseString := strings.ToLower(str)

	// split words by whitespace
	wordList := strings.Fields(lowerCaseString)

	// map which takes Word as a key (string) and occurance of that word as a value (int)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	results := make([]TopTen, len(counts))
	i := 0
	for k, v := range counts {
		results[i] = TopTen{
			Word:  k,
			Count: v,
		}
		i = i + 1
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Count > results[j].Count
	})

	if len(results) > 10 {
		return results[:10]
	}

	return results
}

func WordOccurance(c *gin.Context) {
	uploadTxt := UploadTxt{}
	error := c.BindJSON(&uploadTxt)
	if error != nil {
		helper.HandleError(c, error, error.Error(), 500)
		return
	}

	topten := wordCount(uploadTxt.Text)

	c.JSON(200, gin.H{
		"code":          200,
		"Top Ten Words": topten,
		"success":       true,
	})
}
