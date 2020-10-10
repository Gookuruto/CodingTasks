===============Running Options===============
1. Standard way run command "go run main.go" for dev if you want to build then run go build and run generated executable filesizeformat
2. Docker way without compose:
    - from direcotry above "codingtask" run comand "docker build codingtask -t yourtag"
    - then start container with "docker run -it -p "8080:8080" yourtag"
3. Docker with docker compose
    - from direcotry above "codingtask" run comand "docker build codingtask -t ct"
    - then from directory where docer-compose.yml file is run docker-compose up

in every scenarion api should be visible on localhost:8080