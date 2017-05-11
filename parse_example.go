package main
import "fmt"
import "regexp"
import "encoding/json"
import "strings"

type incoming_message struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

func parse_msg( income [] byte ) {

	var msg incoming_message

	err := json.Unmarshal(income, &msg)
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	fmt.Println(msg.Id)
	re, err := regexp.Compile("[^,]+")
	if err != nil {
		fmt.Println("error:",err)
		return
	}

	if strings.Compare(msg.Id, "login") == 0 {
	fmt.Println(re.FindAllString(msg.Data, 2)[0])
	fmt.Println(re.FindAllString(msg.Data, 2)[1])
	}
	
	if strings.Compare(msg.Id, "registration") == 0 {
	fmt.Println(re.FindAllString(msg.Data, 2)[0])
	fmt.Println(re.FindAllString(msg.Data, 2)[1])
	}
}

func main(){
	data := []byte(`
	{
		"id":"login",
		"data":"username,password"
	}
	`)

	parse_msg(data)
}