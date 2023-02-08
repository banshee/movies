Case - API Project Structure - Estimation 3 hours.
Please write a small Golang server to search movies from http://www.omdbapi.com/ and
get a movie detail.
The Backend should:
Have 2 endpoint named "/search" with GET method and "/detail/:id" with GET
method (single movie detail)
Contain access credential as env to call the API such as:
Key : "****"
URL : http://www.omdbapi.com/
Example url call to search is --> GET
http://www.omdbapi.com/?apikey=*****&s=Batman&page=2
Be written in Golang (can use HTTP framework like Gin, Gorilla Mux, etc)
Important aspects:
Readability of code
Good display on the knowledge of "Separation of Concerns for Codes"
Write unit tests on some of the important files. Bigger plus points for complete unit
test cases
Plus points:
Implementation of Clean Architecture is a BIG plus
Complete Unit tests
Submission:
Github repository link