package translator

import (
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
	"strings"
)

const margin = "  "

var (
	underline = color.New(color.Underline).SprintFunc()
	green     = color.New(color.FgGreen).SprintFunc()
	magenta   = color.New(color.FgMagenta).SprintFunc()
	yellow    = color.New(color.FgYellow).SprintFunc()
)

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

	tweetString += fmt.Sprintf("%s%s @%s %s\n", margin, yellow(tweet.User.Name), tweet.User.ScreenName, underline(tweet.CreatedAt))
	tweetString += fmt.Sprintf("%s%s\n\n", margin, decolateTweetText(tweet.Text))
	tweetString += fmt.Sprintf("%s%s: %d  %s: %d\n\n", margin, green("RT"), tweet.RetweetCount, magenta("Like"), tweet.FavoriteCount)

	return tweetString, nil
}
