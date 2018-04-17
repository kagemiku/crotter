package translator

import (
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"strings"
)

const margin = "  "

func giveMargin(text string) string {
	marginNL := fmt.Sprintf("\n%s", margin)
	return strings.Replace(text, "\n", marginNL, -1)
}

func decolateTweetText(text string) string {
	return giveMargin(text)
}

func TranslateTweet(tweet *twitter.Tweet) (string, error) {
	if tweet == nil {
		return "", errors.New("arg 'tweet' is nil")
	}

	tweetString := ""

	tweetString += fmt.Sprintf("%s%s @%s %s\n", margin, tweet.User.Name, tweet.User.ScreenName, tweet.CreatedAt)
	tweetString += fmt.Sprintf("%s%s\n\n", margin, decolateTweetText(tweet.Text))
	tweetString += fmt.Sprintf("%sRT: %d  Like: %d\n\n", margin, tweet.RetweetCount, tweet.FavoriteCount)

	return tweetString, nil
}
