package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"translator"
)

type AppConfig struct {
	TwitterAuth TwitterAuthConfig `yaml:"twitter_auth"`
}

type TwitterAuthConfig struct {
	ConsumerKey       string `yaml:"consumer_key"`
	ConsumerSecret    string `yaml:"consumer_secret"`
	AccessToken       string `yaml:"access_token"`
	AccessTokenSecret string `yaml:"access_token_secret"`
}

func loadAppConfig(filepath string) (*AppConfig, error) {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	const configFilepath = "./etc/config.yml"
	appConfig, err := loadAppConfig(configFilepath)
	if err != nil {
		panic(err)
	}

	oauthConfig := oauth1.NewConfig(appConfig.TwitterAuth.ConsumerKey, appConfig.TwitterAuth.ConsumerSecret)
	token := oauth1.NewToken(appConfig.TwitterAuth.AccessToken, appConfig.TwitterAuth.AccessTokenSecret)

	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		tweetString, err := translator.TranslateTweet(&tweet)
		if err != nil {
			panic(err)
		}
		fmt.Println(tweetString)
	}
}
