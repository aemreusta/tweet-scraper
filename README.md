# Twitter Scraper Go

This is a Go program that searches for tweets using the [Twitter Scraper](github.com/n0madic/twitter-scraper "Twitter Scraper") package and saves the results as a CSV file.

## Requirements

Go 1.16 or later
[github.com/n0madic/twitter-scraper](github.com/n0madic/twitter-scraper "Twitter Scraper") package

## Installation

Install Go from the [official website](https://go.dev/doc/install).
Install the github.com/n0madic/twitter-scraper package by running

```bash
go get -u github.com/n0madic/twitter-scraper
```

Clone this repository or download the ZIP file and extract it.

## Usage

1. Open a terminal or command prompt.
2. Navigate to the directory containing saveCSV.go.
3. Run

    ```bash
    go run saveCSV.go
    ```

    to search for tweets using the default query ("golang") and number of tweets (10).
4. Optionally, run

    ```bash
    go run main.go -q=<query> -n=<numTweets>
    ```

    to search for tweets using a custom query and number of tweets.
5. Check the current directory for a CSV file containing the tweet data.

## Options

The following options can be passed to the program:

- -q=\<query\>: the query to search for tweets (default: "golang").
- -n=<numTweets\>: the number of tweets to search for (default: 10).

## Output

The program saves the tweet data as a CSV file in the current directory. The filename is in the format *YYYY-MM-DD_HH-MM-SS_query.csv*, where YYYY-MM-DD_HH-MM-SS is the current date and time and query is the search query.

The CSV file contains the following columns:

1. tweet-time: the timestamp of the tweet in the format YYYY-MM-DD_HH-MM-SS.
2. tweet-id: the unique ID of the tweet.
3. tweet-user: the username of the tweet author.
4. tweet-likes: the number of likes of the tweet.
5. tweet-retweets: the number of retweets of the tweet.
6. tweet-hashtag: the hashtag(s) used in the tweet.
7. tweet-text: the text of the tweet.

## License

This program is licensed under the MIT License.

## Credits

This program uses the following open-source libraries:

[github.com/n0madic/twitter-scraper](github.com/n0madic/twitter-scraper "Twitter Scraper"): a Go package for scraping tweets from Twitter.

## Contributing

Contributions are welcome! If you find a bug or have an idea for a feature, please open an issue or submit a pull request.

## Author

This program was written by [Ahmet Emre Usta](https://www.linkedin.com/in/a-emreusta/ "LinkedIn").
