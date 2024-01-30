***Instructions***

1) In the terminal, execute the command "go run Web_Scraper.go.  This will run the Go web scrapign program and generate 10 files that correspond to the URLs that are requested in the assignment. The performance time of the web scraping script will be recorded in the terminal output once the script completes.  

2) In the terminal, execute the command "python run-articles-spider.py".  This will execute the python web scrapping script that was provided with the assignment.  

Based on local runs, the Go script typically executed in ~600 milliseconds.  (Note one run took as long as 1.65 seconds.) This was rather expedient and satisfactory performance against a relatively small amount of URLs and scraped data.  By comparison, the python we scrappin script took about 10 times longer with execution times in the 13 or 14 second range. 

Admittedly, the python script is doing more in that it writes out both the html and JSON files.  That said, the performance advantages of Go are apparent in this assignment.
