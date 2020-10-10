package main

func main() {
	//I leave it in debug mode so it easier to spot when app is runing in console
	//gin.SetMode(gin.ReleaseMode)
	r := NewRouter()
	r.Run(":8080")
}
