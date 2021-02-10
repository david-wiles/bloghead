package internal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type xmlLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type xmlEntry struct {
	Title   string  `xml:"title"`
	Link    xmlLink `xml:"link"`
	Updated string  `xml:"updated"`
	ID      string  `xml:"id"`
	Content struct {
		Type string `xml:"type,attr"`
		Text string `xml:",cdata"`
	} `xml:"content"`
}

type feedXML struct {
	XMLName  xml.Name  `xml:"feed"`
	Title    string    `xml:"title"`
	Subtitle string    `xml:"subtitle"`
	Links    []xmlLink `xml:"link"`
	Updated  string    `xml:"updated"`
	ID       string    `xml:"id"`
	Author   struct {
		Name  string `xml:"name"`
		Email string `xml:"email"`
	} `xml:"author"`
	Entries []xmlEntry `xml:"entry"`
}

// Write an RSS feed.xml based on the pages in the config's Articles field
// The site's domain and author fields must be configured for this to work
func (bh *BlogHead) writeFeed() error {
	feed := feedXML{
		Title:    bh.config.Title,
		Subtitle: bh.config.SubTitle,
		Links: []xmlLink{
			{
				Href: "https://" + path.Join(bh.config.Domain, "feed.xml"),
				Rel:  "self",
				Type: "application/atom+xml",
			},
			{
				Href: "https://" + bh.config.Domain,
				Rel:  "alternate",
				Type: "text/html",
			},
		},
		Updated: time.Now().Format(time.RFC3339),
		ID:      bh.config.Domain,
		Author: struct {
			Name  string `xml:"name"`
			Email string `xml:"email"`
		}{
			Name:  bh.config.Author,
			Email: bh.config.Email,
		},
		Entries: []xmlEntry{},
	}

	for _, page := range bh.config.Articles {
		text, title, updated, err := bh.getArticleData(page)
		if err != nil {
			return err
		}

		feed.Entries = append(feed.Entries, xmlEntry{
			Title: title,
			Link: xmlLink{
				Href: "https://" + path.Join(bh.config.Domain, trimPath(bh.Root, page)),
			},
			Updated: updated,
			ID:      path.Join(bh.config.Domain, trimPath(bh.Root, page)),
			Content: struct {
				Type string `xml:"type,attr"`
				Text string `xml:",cdata"`
			}{"html", text},
		})
	}

	f, err := createFile(path.Join(bh.Output, "feed.xml"))
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := xml.NewEncoder(f)
	if err := encoder.Encode(feed); err != nil {
		return err
	}

	return nil
}

func (bh *BlogHead) getArticleData(page string) (text, title, updated string, err error) {

	// Get article metadata
	b, err := ioutil.ReadFile(page[:len(page)-5] + "_meta.json")
	if err != nil {
		return text, title, updated, err
	}

	// Make a temporary file to compile templates with
	tmpF, err := os.Create(path.Join(os.TempDir(), "_tmp.html"))
	if err != nil {
		return "", "", "", err
	}
	defer tmpF.Close()

	if _, err := tmpF.WriteString(
		fmt.Sprintf("{{template \"%v\" .}}", path.Join(".data", path.Base(page), "content.html"))); err != nil {
		return "", "", "", err
	}

	// Copy article metadata to temporary file
	tmplData, err := os.Create(path.Join(os.TempDir(), "_tmp_meta.json"))
	if err != nil {
		return "", "", "", err
	}
	defer tmplData.Close()

	if _, err = tmplData.Write(b); err != nil {
		return "", "", "", err
	}

	meta := make(map[string]string)
	if err := json.Unmarshal(b, &meta); err != nil {
		return "", "", "", err
	}

	var textBytes []byte
	textBytes, err = bh.compile(tmpF.Name())
	if err != nil {
		return "", "", "", err
	}

	text = string(textBytes)

	if m, ok := meta["title"]; ok {
		title = m
	}

	if m, ok := meta["updated"]; ok {
		updated = m
	}

	return text, title, updated, nil
}
