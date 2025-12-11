package main


import (
	"fmt"
	"strings"
)


func main(){
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Index("test", "a"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Join([]string{"a", "b"}, ""))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
	fmt.Println([]byte("test"))
	fmt.Println(string([]byte("test")))
	fmt.Println(string([]byte{'t', 'e', 's', 't'}))
}

func Copy(dst Writer, src Reader) (written int64, err error){
	var buf bytes.Buffer
	buf.Write([]bytes["test"])

}

