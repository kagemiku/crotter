package translator

import (
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

func TranslateTweet(tweet *twitter.Tweet) (string, error) {
	if tweet == nil {
		return "", errors.New("arg 'tweet' is nil")
	}

	tweetString := ""

	tweetString += fmt.Sprintf("  %s @%s %s\n", tweet.User.Name, tweet.User.ScreenName, tweet.CreatedAt)
	tweetString += fmt.Sprintf("  %s\n\n", tweet.Text)
	tweetString += fmt.Sprintf("  RT: %d  Like: %d\n\n", tweet.RetweetCount, tweet.FavoriteCount)

	return tweetString, nil
}
