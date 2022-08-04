package main

import (
	"fmt"
	"os"
	"strings"

	isbnpkg "github.com/moraes/isbn"
	"github.com/playwright-community/playwright-go"
)

const (
    LEXILE = "https://hub.lexile.com/find-a-book/book-details/"
    LEXILE_SELECTOR = "#content > div > div > div > div.details > div.metadata > div.sc-kexyCK.cawTwh > div.header-info > div > span"
    
    ATOS = "https://www.arbookfind.com/UserType.aspx?RedirectURL=%2fadvanced.aspx"
    RAD = "#radLibrarian"
    SUBMIT = "#btnSubmitUserType"
    ISBN_BOX = "#ctl00_ContentPlaceHolder1_txtISBN"
    SEARCH = "#ctl00_ContentPlaceHolder1_btnDoIt"
    SEARCH_FAIL = "#ctl00_ContentPlaceHolder1_lblSearchResultFailedLabel"
    TITLE = "#book-title"
    ATOS_SELECTOR = "#ctl00_ContentPlaceHolder1_ucBookDetail_lblBookLevel"
)

var (
    pw *playwright.Playwright
    browser playwright.Browser
    page playwright.Page

    Args []string
    Flags map[string]bool
)

func main() {
    ProcessFlags()
    
    if Flag("install", false) {
        run := playwright.RunOptions{Browsers: []string{"chromium"}}
        playwright.Install(&run)
    }

    if len(Args) > 1 {
        fmt.Print("Need 1 argument: <ISBN>")
        return
    } else if len(Args) == 0 {
        if !Flag("install", false) {
            Help()
        }
        return
    }
    isbn := Args[0]
    valid := isbnpkg.Validate(isbn)
    if !valid {
        fmt.Print("Invalid ISBN!")
        return
    }

    var err error
    pw, err = playwright.Run()
    catch(err)
    defer pw.Stop()

    browser, err = pw.Chromium.Launch()
    catch(err)
    defer browser.Close()

    page, err = browser.NewPage()
    catch(err)
    
    ar := Atos(isbn)
    lex := Lexile(isbn)
    
    Print(lex, ar)
}

func Lexile(isbn string) int {
    page.Goto(fmt.Sprint(LEXILE, isbn))
    if page.URL() == "https://hub.lexile.com/find-a-book/book-results" {
        return -1
    }

    str, err := page.TextContent(LEXILE_SELECTOR)
    if err != nil {
        return -1
    }
    var lex int
    if _, err := fmt.Sscan(str, &lex); err != nil {
        return -1
    }
    return lex
}

func Atos(isbn string) float64 {
    page.Goto(ATOS)
    page.Click(RAD) //Select Librarian and submit
    page.Click(SUBMIT)

    page.WaitForSelector(ISBN_BOX)
    page.Type(ISBN_BOX, isbn)
    page.Click(SEARCH)
    
    page.WaitForLoadState("domcontentloaded")
    fail, _ := page.Locator(SEARCH_FAIL)
    count, _ := fail.Count()
    if count > 0 {
        return -1
    }

    page.WaitForSelector(TITLE)
    page.Click(TITLE) //Click on first book

    
    
    str, err := page.TextContent(ATOS_SELECTOR) //Get level from selector
    if err != nil {
        return -1
    }
    var ar float64
    fmt.Sscan(str, &ar) //Convert level to float
    return ar
}

func Print(lex int, ar float64) {
    raw := Flag("raw", false)
    ln := Flag("ln", false)

    var lexile string
    var atos string

    if lex == -1 {
        lexile = "Not found!"
    } else {
        lexile = fmt.Sprint(lex)
    }

    if ar == -1 {
        atos = "Not found!"
    } else {
        atos = fmt.Sprint(ar)
    }

    if raw {
        fmt.Print(lex)
        if ln {
            fmt.Println()
        } else {
            fmt.Print(" ")
        }
        fmt.Print(ar)
    } else {
        fmt.Print("Lexile: ", lexile)
        if ln {
            fmt.Println()
        } else {
            fmt.Print(" | ")
        }
        fmt.Print("AR: ", atos)
    }
}

func Help() {
    fmt.Println(`
Lexos cli:
This tool is used for gathering the Lexile and Atos levels of books via their ISBN

Usage: lexos <ISBN> [--raw, --ln, --install]
--raw: Print the raw numbers to the output, without labels (Prints Lexile followed by the Atos, as well as printing -1 if the result cannot be found)
--ln: Seperates the outputs with a new line
--install: Installs the necessary driver and browser to run. This argument is required if it has not already been run, otherwise the program will throw an error.`)
}

func ProcessFlags() {
    Flags = make(map[string]bool)
    args := os.Args[1:]
    for _, arg := range args {
        if strings.HasPrefix(arg, "--") {
            Flags[strings.TrimPrefix(arg, "--")] = true
        } else if strings.HasPrefix(arg, "-") {
            Flags[strings.TrimPrefix(arg, "-")] = true
        } else {
            Args = append(Args, arg)
        }
    }
}

func Flag(name string, def bool) bool {
    if Flags[name] {
        return !def
    } else {
        return def
    }
}

func catch(err error) {
    if err != nil {
        panic(err)
    }
}

