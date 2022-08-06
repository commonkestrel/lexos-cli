# Lexos

This tool is used for gathering the Lexile level, Atos(AR) level, and AR Points of books via their ISBN.

It uses the [isbn](https://github.com/moraes/isbn) package to validate the ISBN, along with the [playwright-go](https://github.com/playwright-community/playwright-go) package to find the results in a headless browser. 
This does take quite a while depending on your internet connection, but unfortunatly since Lexile's Find A Book requires Javascript, and ARBookFinder is a collection of ASPX pages, there is currently no workaround.

If you have Go installed on your system, run ```go install github.com/Jibble330/lexos@latest``` to install.
If you don't, download the files and add the folder to your PATH.

Usage: ```lexos <ISBN> [--raw, --ln, --install]``` <br/>
```--raw```: Print the raw numbers to the output, without labels (Prints in order: Lexile Level, Atos Level, AR Points, as well as printing -1 if the result cannot be found) <br/>
```--ln```: Seperates the outputs with a new line <br/>
```--install```: Installs the necessary driver and browser to run. This argument is required if it has not already been run, otherwise the program will throw an error.
