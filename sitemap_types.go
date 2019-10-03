package sitemap

import "time"

type sitemapEntry struct {
	Location        string    `xml:"loc"`
	LastModified    string    `xml:"lastmod,omitempy"`
	ChangeFrequency Frequency `xml:"changefreq,omitempty"`
	Priority        float32   `xml:"priority,omitempty"`
}

func newSitemapEntry() *sitemapEntry {
	return &sitemapEntry{ChangeFrequency: Always, Priority: 0.5}
}

func (e *sitemapEntry) GetLocation() string {
	return e.Location
}

func (e *sitemapEntry) GetLastModified() *time.Time {
	return parseDateTime(e.LastModified)
}

func (e *sitemapEntry) GetChangeFrequency() Frequency {
	return e.ChangeFrequency
}

func (e *sitemapEntry) GetPriority() float32 {
	return e.Priority
}

type sitemapIndexEntry struct {
	Location     string `xml:"loc"`
	LastModified string `xml:"lastmod,omitempty"`
}

func newSitemapIndexEntry() *sitemapIndexEntry {
	return &sitemapIndexEntry{}
}

func (e *sitemapIndexEntry) GetLocation() string {
	return e.Location
}

func (e *sitemapIndexEntry) GetLastModified() *time.Time {
	return parseDateTime(e.LastModified)
}

func parseDateTime(value string) *time.Time {
	if value == "" {
		return nil
	}

	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		// second chance
		// try parse as short format
		t, err = time.Parse("2006-01-02", value)
		if err != nil {
			return nil
		}
	}

	return &t
}
