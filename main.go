package main

import (
	"go.uber.org/zap"
	"logCustomized/logCustomized"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var sugarLogger *zap.SugaredLogger

func simpleHttpGet(url string) {
	defer wg.Done()

	sugarLogger.Debugf("尝试进行 GET 请求：%s", url)

	resp, err := http.Get(url)

	if err != nil {
		sugarLogger.Errorf("获取 URL 时出错：%s，错误信息：%s", url, err)
	} else {
		sugarLogger.Infof("成功！状态码：%s，URL：%s", resp.Status, url)
		err := resp.Body.Close()
		if err != nil {
			return
		}

	}
}

func main() {
	sugarLogger = logCustomized.InitLogger(logCustomized.GetEncoder(), logCustomized.GetLogConsole())

	defer func(sugarLogger *zap.SugaredLogger) {
		err := sugarLogger.Sync()
		if err != nil {
		}
	}(sugarLogger)

	URL := []string{
		"https://www.google.com", "https://www.YouTube.com", "https://www.Facebook.com", "https://www.twitter.com", "https://www.instagram.com",
		"https://www.baidu.com", "https://www.wikipedia.org", "https://www.yandex.ru", "https://www.yahoo.com", "https://www.xvideos.com",
		"https://www.whatsapp.com", "https://www.xnxx.com", "https://www.yahoo.co.jp", "https://www.amazon.com", "https://www.live.com",
		"https://www.netflix.com", "https://www.pornhub.com", "https://www.office.com", "https://www.tiktok.com", "https://www.reddit.com",
		"https://www.zoom.us", "https://www.linkedin.com", "https://www.vk.com", "https://www.xhamster.com", "https://www.discord.com",
		"https://www.bing.com", "https://www.Naver.com", "https://www.twitch.tv", "https://www.mail.ru", "https://www.microsoftonline.com",
		"https://www.duckduckgo.com", "https://www.roblox.com", "https://www.bilibili.com", "https://www.qq.com", "https://www.pinterest.com",
		"https://www.Microsoft.com", "https://www.msn.com", "https://www.docomo.ne.jp", "https://news.yahoo.co.jp", "https://www.globo.com",
		"https://www.samsung.com", "https://www.google.com.br", "https://www.t.me", "https://www.eBay.com", "https://www.turbopages.org",
		"https://www.accuweather.com", "https://www.ok.ru", "https://www.bbc.co.uk", "https://www.fandom.com", "https://www.weather.com"}

	for _, url := range URL {
		wg.Add(1)
		go simpleHttpGet(url)
	}

	wg.Wait()
}
