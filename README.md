# goWordPressXML2HTML 

# About

This code is for a converter that converts WordPress XML into HTML.  
The HTML that results can then be fed into Pandoc to convert to other formats such as docx or pdf. 

# How to build 

go build -o goWordPressXML2HTML 

# How to run 

goWordPressXML2HTML -input <wordpress.xml> -output <outputfilename.html>

# Notes

The code currently relies on the latest version of [go-wordpressxml library ](https://github.com/grokify/go-wordpressxml) 

The system is similar to the Python program - [wordpressXML2HTML](https://github.com/Petess/wordpressXML2HTML)

