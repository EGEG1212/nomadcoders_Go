package main

import (
	"fmt"
	"log"
	"net/http"
)

type extractedJob struct {
	id string
	title string
	location string
	salary string
	summary string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	var jobs []extractedJob
	c := make(chan []extractedJob)//여러개니까 []
	totalPages := getPages()
	// fmt.Println(totalPages)
	for i := 0; i < totalPages; i++ {
		//extractedJob := getPage(i)
		go getPage(i, c)
		//jobs = append(jobs, extractedJob...) //2개의 배열을 합치려면 ... 쩜3개👌 만약 안찍으면 배열안에 배열을 넣는 모양이 됨..[[],[],[]] ;; chan만들면서 삭제
	}
for i:=0; i<totalPages; i++ {
	extractedJob := <- c
	jobs = append(jobs, extractedJob...)
}

	writeJobs(jobs) //csv파일을만들어줌
	fmt.Println("Done, extracted", Length(jobs) )
}

func getPage(page int,mainC chan<- []extractedJob){ //pane num 몇페이지인지 알아야하니까
	var jobs = []extractedJob
	c := make(chan extractedJob) //👩‍🚀채널로 전송할 값은 extractedJob이다.
	//pageURL := baseURL + "&start=" + page*50 //string과 int의 조합은 아래와 같이 쓴다.
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() //닫아줘야 메모리가 새어나가는걸 막아줄수있다✋

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection){
		//job := extractedJob(card, c) //이제 아래와같이 goroutine 되었다.
		//jobs = append(jobs, job)
		go extractedJob(card, c)
	})
	for i:=0; i<searchCards.Length(); i++ { //카드당한번씩반복되기때문에 한페이지에 50카드씩 계속반복
		job := <-c//채널에서 받은 메세지를 job변수에 저장
		jobs = append(jobs, job)
	}

	//return jobs
	mainC <- jobs

}

func getPages(page int) []extractedJob {
	var jobs = []extractedJob
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() //닫아줘야 메모리가 새어나가는걸 막아줄수있다✋

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	// fmt.Println(doc)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) { //
		// fmt.Println(s.Html()) //HTML확인
		// fmt.Println(s.Find("a").Length()) //하단페이지 갯수가져오는지 확인
		pages = s.Find("a").Length()
	})

	return pages
}


func extractedJob(card *goquery.Selection, c chan<- extractedJob)  {
	id, _ := card.Attr("href")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	//fmt.Println(id, title, location, salary, summary) //😭윽..보기힘들어
	c <-  extractedJob{id:id,   //👩‍🚀return할 필요없고, 대신 채널에 값을 전송하기
						title:title, 
						location:location, 
						salary:salary, 
						summary:summary}
	
}


func cleanString(str string) []string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")//Fields는 string으로 된 배열을 반환한다.
}


func writeJobs(jobs []extractedJob){
	file. err := os.Create("jobs.csv") //파일생성, 에러체크
	checkErr(err)

	w := csv.NewWriter(file) //쓰고
	defer w.Flush()			//종료시점에서 Flush를 실행하도록 

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	
	wErr := w.Wirete(headers)
	checkErr(wErr)

	for _, job := range jobs {  //#4.6 https://github.com/tsak/concurrent-csv-writer 라이브러리로 구현가능
		jobSlice := []string{"https://kr.indeed.com/?vjk=" + job.id, job.title, job.location, jon.salary, job.summary}
		jwErr := w.Wirete(jobSlice) //파일생성
		checkErr(jwErr)
	}
}


func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) { //👆위에 Get부분을 확인하면 *포인터가있다
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

