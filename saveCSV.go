package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

// Extracts tweet data from a tweet result
func extractTweetData(tweet twitterscraper.TweetResult) []string {
	// Convert time to string
	tweetTime := tweet.TimeParsed.Format("2006-01-02_15-04-05")
	userID := tweet.UserID
	id := tweet.ID
	// Convert likes to string
	likes := fmt.Sprintf("%d", tweet.Likes)
	// Convert retweets to string
	retweets := fmt.Sprintf("%d", tweet.Retweets)
	// Convert hashtags to string
	hashtag := fmt.Sprintf("%s", tweet.Hashtags)

	// Remove empty lines from the tweet text
	text := strings.TrimSpace(tweet.Text)
	lines := strings.Split(text, "\n")
	nonEmptyLines := []string{}
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	text = strings.Join(nonEmptyLines, " ")

	// Return the tweet information as a slice of strings
	return []string{tweetTime, id, userID, likes, retweets, hashtag, text}
}

// Searches for tweets using a query and returns a DataFrame containing the results
func SearchTweets(query string, numTweets int) (*dataframe.DataFrame, error) {
	scraper := twitterscraper.New()
	tweets := [][]string{{"tweet-time", "tweet-id", "tweet-user", "tweet-likes", "tweet-retweets", "tweet-hashtag", "tweet-text"}}

	for tweet := range scraper.SearchTweets(context.Background(), query, numTweets) {
		if tweet.Error != nil {
			return nil, fmt.Errorf("error searching tweets: %v", tweet.Error)
		}
		tweetData := extractTweetData(*tweet)
		if len(tweetData) > 0 {
			tweets = append(tweets, tweetData)
		}
	}

	if len(tweets) == 1 {
		return nil, errors.New("no tweets found")
	}

	df := dataframe.LoadRecords(tweets)
	return &df, nil
}

// Define constants
const (
	query     = "secim"
	numTweets = 100
)

// main function
func main() {
	// Search for tweets
	df, err := SearchTweets(query, numTweets)
	if err != nil {
		// Log the error and exit gracefully
		log.Fatalf("Error: %v", err)
	}

	// Get the current date and time to use as part of the filename
	t := time.Now()
	filename := t.Format("2006-01-02_15-04-05")

	// Add query to filename
	filename = fmt.Sprintf("%s_%s", filename, query)

	// Open the file to save the dataframe
	f, err := os.Create(fmt.Sprintf("%s.csv", filename))
	if err != nil {
		// Log the error and exit gracefully
		log.Fatalf("Error creating file: %v", err)
	}
	defer f.Close()

	// Save the dataframe as a CSV file
	if err := df.WriteCSV(f); err != nil {
		// Log the error and exit gracefully
		log.Fatalf("Error saving dataframe: %v", err)
	}

	// Print success message
	fmt.Printf("Dataframe saved as %s.csv\n", filename)
}
