package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func checkTwitter(username string) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:105.0) Gecko/20100101 Firefox/105.0"),
	)
	c.OnHTML("a[class=\"profile-card-username\"]", func(e *colly.HTMLElement) {
		fmt.Println("[+] Twitter: https://twitter.com/" + e.Attr("href"))
	})

	c.OnHTML("div[class=\"profile-website\"]", func(e *colly.HTMLElement) {
		fmt.Println("[+] Twitter website:", e.ChildAttr("a", "href"))
	})

	c.OnError(func(r *colly.Response, err error) {
		if r.StatusCode != 404 {
			log.Println("Twitter:", r.StatusCode, err.Error())
		}

	})
	c.Visit("https://nitter.net/" + username)
}

func checkInstagram(username string) {
	c := colly.NewCollector()

	c.OnHTML("meta[property=\"og:url\"]", func(h *colly.HTMLElement) {
		fmt.Println("[+] Instagram:", h.Attr("content"))
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Instagram:", r.StatusCode, err.Error())
	})

	c.Visit("https://www.instagram.com/" + username + "/")
}

func checkTelegram(username string) {
	c := colly.NewCollector()
	c.OnHTML("div[class=\"tgme_page_title\"]", func(h *colly.HTMLElement) {
		fmt.Println("[+] Telegram: https://t.me/" + username)
		fmt.Println("[+] Telegram name:", h.Text)
	})
	c.OnHTML("div [class=\"tgme_page_description \"]", func(h *colly.HTMLElement) {
		fmt.Println("[+] Telegram description:", h.Text)
	})
	c.Visit("https://t.me/" + username)
}

func checkTryHackMe(username string) {
	c := colly.NewCollector()
	c.OnHTML("span[class~=\"level\"]", func(h *colly.HTMLElement) {
		fmt.Println("[+] TryHackMe: https://tryhackme.com/p/" + username)
	})
	c.OnHTML("div[class=\"text-muted links-blank\"]", func(h *colly.HTMLElement) {
		for _, i := range h.ChildAttrs("a", "href") {
			fmt.Println("[+] TryHackMe link:", i)
		}
	})
	c.Visit("https://tryhackme.com/p/" + username)
}

func main() {
	checkTwitter(os.Args[1])
	checkInstagram(os.Args[1])
	checkTelegram(os.Args[1])
	checkTryHackMe(os.Args[1])
}
