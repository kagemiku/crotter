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

	tweetString := fmt.Sprintf("%s: %s", tweet.IDStr, tweet.Text)

	return tweetString, nil
}
