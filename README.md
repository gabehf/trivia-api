Stateless, Containerized Trivia API
---

This project has a blog post going through the development process. Read the post here: https://mnrva.dev/posts/stateless-containerized-trivia-api-go

## Running the server
Copy the repo
```bash
git clone git@github.com:gabehf/trivia-api.git
```
Then, run the server using Go
```bash
cd trivia-api
go run .
```
Or, build the docker image and run that
```bash
docker build --tag trivia-api . 
```
Then,
```bash
docker run -p 3000:3000 trivia-api
```