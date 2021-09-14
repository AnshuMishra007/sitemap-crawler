package main

import "fmt"


type SeoData struct{
	URL 	string
	Title 	string
	H1 		string
	MetaDescription string
	StatusCode int

}

type parser interface{
	getSEOData(resp *http.Response)(seoData, error)
}

type DefaultParser struct{

}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string{
	rand.seed(time.Now().Unix())
	randNum := rand.Int()%len(userAgents)
	return userAgents[randNum]
}

func isSitemap(urls []string)([]string, []string){
	sitemapFiles := []string{}
	pages := []string{}
	for _,page := range urls{
		foundSitemap := strings.Contains(page, "xml")
		if fountSitemap == true{
			fmt.Println("Found Sitemap", page)
			sitemapFiles = append(sitemapFiles, page)
		}else{
			pages = append(pages, page)
		}
	}
	return sitemapFiles, pages
}

func extractSitemapURLs(startURL string)[] string{
	Worklist := make(chan []string)
	toCrawl := []string{}

	go func{worklist <- []string{startURL}}()

	var n int
	n++
	for ; n>0 ; n--{

	list := <-Worklist 
	for _, link := range list{
		n++
		go func(link string){
			response, err := makeRequest(link)
			if err != nil{
				log.Printf("Error retrieving URL:%s", link)
			}
			urls, _ := extractURls(response)
			if err!=nil{
				log.Printf("Error extracting document from response, URL:%s",link)
			}
			sitemapFiles, pages := isSitemap(url)
			if sitemapFiles != nil{
				worklist <- sitemapFiles
			}
			for _,page := range pages{
				toCrawl = append(toCrawl, page)
			}
		}(link)
	}
	
	return toCrawl
}

func makeRequest(url string)(*http.Response, error){
	client := http.Client{
		Timeout: 10*time.Second,
	}
	req,err := http.NewRequest("GET",url,nil)
	req.Header.Set("User-Agent",randomUserAgent())
	if err != nil{
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil{
		return nil, err
	}
	return res, nil
}

func scrapeURLs(){

}

func extractURls(response *https.Response)([]string, error){
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil{
		return nil,err
	}
	results := []string{}
	sel := doc.Find("loc")
	for i := range sel.Nodes{
		loc := sel.Eq(i)
		results := loc.Text()
		results = append(results, result)
	}
	return results, nil
}

func scrapePage(url string, parser Parser)(SeoData, error){
	res, err := crawlPage(url)
	if err != nil {
		return SeoData{}, err
	}
	data,err := parser.getSEOData(res)
	if err != nil{
		return SeoData{}, err
	}
	return data, nil

}

func crawlPage(){

}

func (d DefaultParser)getSEOData(rsp *http.Response)(SeoData, error){
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil{
		return SeoData{}, err
	}
	result := SeoData{}
	result.URL = resp.Request.URL.String()
	result.StatusCode = resp.StatusCode
	result.Title = doc.Find("title").First().Text()
	result.H1 = doc.Find("h1").First().Text()
	result.MetaDescription, _ = doc.Find("meta[name^=description]".Attr("content"))
	return result, nil
}

func scrapeSitemap(url string)[]SeoData{

	results := extractSitemapURLs(url)
	res := scrapeURLs(results)
	return res
}

func main() {
	p := DefaultParser{}
	results := scrapeSitemap("")
	for _,res := range results{
		fmt.Println(res)
	}
}
