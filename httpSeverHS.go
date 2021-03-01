//	Code for YouTube Go Tutorial course Titled:
//	Advanced Enterprise Platform Coding with Go in Depth Programming
//	https://youtu.be/whL0Toym5u4

package main
import "net/http"
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":80", nil)
}
