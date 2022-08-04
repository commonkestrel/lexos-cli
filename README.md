# Lexos

This tool is used for gathering the Lexile and Atos(AR) levels of books via their ISBN
It uses the playwright-go package found [here](https://github.com/playwright-community/playwright-go) to find the results in a headless browser. \
This does take quite a while depending on your internet connection, but unfortunatly since Lexile book finder requires Javascript, and ARBookFinder is a collection of ASPX pages, there is currently no workaround.

In order to use in the terminal, you will first have to add the parent folder to your path, in Environment Variables.

Usage: ```lexos <ISBN> [--raw, --ln, --install]``` <br/>
```--raw```: Print the raw numbers to the output, without labels (Prints Lexile followed by the Atos, as well as printing -1 if the result cannot be found) <br/>
```--ln```: Seperates the outputs with a new line <br/>
```--install```: Installs the necessary driver and browser to run. This argument is required if it has not already been run, otherwise the program will throw an error.
