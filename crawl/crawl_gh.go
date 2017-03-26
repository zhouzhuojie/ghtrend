package crawl

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	ghTrendingURLBase = "https://github.com/trending?l=%s"
	htmlTemplateFile  = "mail/email_template.html"
	htmlTemplateName  = "email_template.html"
)

var (
	languages = []string{"", "go", "python", "javascript"}
)

// GithubTrendingProject is an item that represents the trending item
type GithubTrendingProject struct {
	// Title is the tile of the project
	Title string

	// Description is the description of the project
	Description string

	// URL is the url of the project
	URL string
}

// CrawlGithubTrendingPageWithLanguage will crawl the page given the programming language
var CrawlGithubTrendingPageWithLanguage = func(language string) ([]*GithubTrendingProject, error) {

	doc, err := goquery.NewDocument(fmt.Sprintf(ghTrendingURLBase, language))

	if err != nil {
		return nil, err
	}

	var projects []*GithubTrendingProject

	doc.Find("ol.repo-list li").Each(func(i int, s *goquery.Selection) {

		url, _ := s.Find("h3 a").Attr("href")

		description := s.Find("p.col-9").Text()
		description = strings.Replace(description, "\n", "", -1)
		description = strings.Replace(description, "  ", "", -1)

		p := &GithubTrendingProject{
			Title:       strings.TrimSpace(s.Find("h3 a").Text()),
			Description: description,
			URL:         "https://github.com" + url,
		}

		projects = append(projects, p)
	})

	return projects, nil
}

// FormHTML forms a HTML
var FormHTML = func(m map[string][]*GithubTrendingProject) []byte {

	var b bytes.Buffer

	t := template.Must(template.New(htmlTemplateName).ParseFiles(htmlTemplateFile))
	err := t.Execute(&b, m)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return b.Bytes()
}

// CrawlGithubTrendingPages crawls the pages using the languages defined
var CrawlGithubTrendingPages = func() []byte {
	m := make(map[string][]*GithubTrendingProject)
	var wg sync.WaitGroup
	for _, l := range languages {
		ll := l
		wg.Add(1)
		go func() {
			projects, _ := CrawlGithubTrendingPageWithLanguage(ll)
			m[ll] = projects
			wg.Done()
		}()
	}
	wg.Wait()
	return FormHTML(m)
}
