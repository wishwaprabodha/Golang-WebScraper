# Web Page Analyzer

This is a simple web application built in Go that allows users to analyze a web page by providing its URL. The application uses the Mux router for handling HTTP requests and the goquery library for parsing and querying the HTML of the web page. For the efficiency, Concurrency is being used and Clean Architecture Pattern is applied  

## Installation

To run the application, make sure you have Go installed on your machine. Clone the repository and navigate to the root directory of the project. Then, use the following command to install the required dependencies:

```bash
go get -u github.com/PuerkitoBio/goquery
go get -u github.com/gorilla/mux
```

## How to Use

1. Start the server by running the following command:

```bash
go run main.go
```

2. Once the server is up and running, you can access the web page analyzer by making a POST request to the `/api/analyze` endpoint with the `url` parameter set to the URL of the web page you want to analyze.

For example, you can use tools like `curl` or Postman to make a request to the analyzer:

```bash
curl -X GET "http://localhost:8080/api/analyze?url=https://example.com"
```

## Results

The application will provide the following information about the analyzed web page:

- HTMLVersion: The version of the HTML used in the web page.
- PageTitle: The title of the web page.
- HeadingCounts: The number of headings of each level (h1, h2, h3, h4, h5, h6) present in the document.
- InternalLinks: The number of internal links (links starting with "/") on the web page.
- ExternalLinks: The number of external links (links starting with "http") on the web page.
- InaccessibleLinks: The number of inaccessible external links, i.e., links that returned an error or invalid status code during analysis.
- HasLoginForm: Indicates whether the web page contains a login form.
- Error: Indicates if any error occurred.

In case the URL provided is not reachable or the analysis takes too long (30s), appropriate error messages will be displayed in the response.

