package main

import (
    "bufio"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "golang.org/x/text/encoding/charmap"
    "golang.org/x/text/transform"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

var print = fmt.Println
var timenow = time.Now()

func response(www string) string {
    req, err := http.NewRequest("GET", www, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    //print("Response Status:", resp.Status, "\n")
    pool, err := ioutil.ReadAll(resp.Body)
    text := string(pool[:])
    return text

}

func connect(url string) int {
    res, err1 := http.Get(url)
    if err1 != nil {

        log.Fatal(err1)
    }
    StatusCode := res.StatusCode

    return StatusCode
}

func ParseMCD() {
    url := generateUrl()
    print(url)
    // Request the HTML page.
    //res, err := http.Get("http://10.68.1.222/reports/reports/20181210/report-145001/")
    res, err1 := http.Get(url)
    if err1 != nil {

        log.Fatal(err1)
    }
    defer res.Body.Close()
    if res.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }

    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    band := ""
    // Find the review items
    doc.Find("body").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title

        band = s.Find("div").Text()
    })

    text := strings.Split(band, "\n\n")

    errors := strings.Split(text[19], ":")

    texterrors := strings.Replace(errors[0], "\n  ", "", -1)
    interrors, _ := strconv.Atoi(strings.Replace(errors[1], " ", "", -1))

    if interrors > 8 {
        print(texterrors, " - ", interrors)
    } else {
        os.Exit(0)
    }

    return
}

func generateUrl() string {
    date := timenow.Format("20060102")

    hours := strings.Split(timenow.Format("20060102 1500"), " ")[1] + "02"

    url := "http://10.68.1.222/reports/reports/" + date + "/report-" + hours + "/"

    err := connect(url)

    if err != 200 {
        hours := strings.Split(timenow.Format("20060102 1500"), " ")[1] + "01"
        url = "http://10.68.1.222/reports/reports/" + date + "/report-" + hours + "/"
        err = connect(url)
        if err != 200 {
            print(err)
        }
    }
    return url
}

func scanFile(urls []string) {
    for _, url := range urls {
        filewhrite, _ := os.Create("Test\\" + url + "-WithErrors.csv")

        lines := []string{}
        file, _ := os.Open(url)

        defer file.Close()

        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            text := scanner.Text()
            items := strings.Split(text, ";")

            if len(items[0]) == 0 {

                lines = append(lines, scanner.Text())
                fmt.Fprintf(filewhrite, "%v\n", text)

            }

        }
    }
    return
}
func GenereteFiles() {
    names := []string{}
    urls := []string{"http://10.68.1.57/mgt/file?name=ROUTES&type=TEXT", "http://10.68.1.57/mgt/file?name=TRIP_SHAPES&type=TEXT", "http://10.68.1.57/mgt/file?name=INTERVAL&type=TEXT", "http://10.68.1.57/mgt/file?name=TRIPS&type=TEXT", "http://10.68.1.57/mgt/file?name=STOP_TIMES&type=TEXT", "http://10.68.1.57/mgt/file?name=CHANGE_STOP_TIMES&type=TEXT", "http://10.68.1.57/mgt/file?name=CALENDAR&type=TEXT", "http://10.68.1.57/mgt/file?name=VEHICLES&type=TEXT", "http://10.68.1.57/mgt/file?name=TRIPS_STOPS&type=TEXT"}
    for _, url := range urls {
        names = append(names, url[32:len(url)-10])
        text := response(url)
        fileHandle, _ := os.Create(fmt.Sprintf("CASH\\" + url[32:len(url)-10] + ".csv"))
        writer_encoding := transform.NewWriter(fileHandle, charmap.Windows1251.NewEncoder())
        defer fileHandle.Close()
        defer writer_encoding.Close()
        fmt.Fprint(writer_encoding, fmt.Sprintf("%v\n", text))
    }
    scanFile(names)
}

func Load() {
    ParseMCD()
    GenereteFiles()
}
func main() {
    //"http://10.68.1.222/reports/reports/20181210/report-145001/"

    Load()

}
