# url-shortener
url-shortener

The objective is to provide a REST API URL Shortener in Go.

It provides the following functionalities :
- Shorten URL
- Redirect to original URL when request short URL
- Provide the count of domains of which URLs were shortened

### Running the Application
## Running with docker

0. In Config file, add the required configurations
1. docker compose up    
2. Access the below links, mentioned in Usage

## Running locally
0. In Config file, change the Storage to memory
1. Run "go run main.go"


### Usage

a. To visit Index page, with instructions
    http://localhost:8081

b. To shorten url:
    http://localhost:8081/shortenurl?url=google.com
    http://localhost:8081/shortenurl?url=chat.openai.com

c. To get domain count of which URLs were shortened:
    http://localhost:8081/domaincount

d. Redirect to :
    http://localhost:8081/



TODO:
Basic WebServer Setup: Completed
Parse URL to get Link: Completed
URL Shortner Logic: Partially Completed. Algo tuning Pending
Config Manager Setup: Completed
In Memory Storage Setup: Completed
Validations: Completed
Database Setup: Redis Database Completed
Interfaces for supporting multiple Databases: Completed
Logging: Always In Progress
Testing: Always In Progress
