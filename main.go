///webcrawler предназначенный для сбора информации(id, имя, фамилия, дата рождения)
/// в социальной сети VK, на вход требует начальный id пользователя
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type vkUser struct {
	firstName, lastName, bdate, id string
}

var queue []string
var userFriends []string
var newUser vkUser
var idx int = 0

func main() {
	flag.Parse()
	startId := flag.Args()
	queue = append(queue, startId...)
	GetData(queue[0])
}

func GetData(ID string) []vkUser { // собираем данные о пользователе
	url := "https://api.vk.com/method/friends.get?user_id=" + ID +
		"&fields=bdate&access_token=54a3bfcef3a2b7dc5299273727a63ffb1316a16a5f3059637745cb779d1dc78493381d04c6af81040e979&v=5.92"
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	removeStr1 := strings.ReplaceAll(string(body), ":", "")
	removeStr2 := strings.ReplaceAll(removeStr1, ",", "")
	removeStr3 := strings.ReplaceAll(removeStr2, " ", "")

	words := strings.Split(removeStr3, "\"")

	for i := 0; i < len(words); {
		switch words[i] {
		case "first_name":
			i += 2
			newUser.firstName = words[i]
		case "last_name":
			i += 2
			newUser.lastName = words[i]
		case "id":
			i += 1
			userFriends = append(userFriends, words[i])
			newUser.id = words[i]
		case "bdate":
			i += 2
			newUser.bdate = words[i]
			fmt.Println(newUser.id, "\t\t", newUser.firstName, newUser.lastName, "\t\t", newUser.bdate)
			newUser.id = "-"
			newUser.firstName = "-"
			newUser.lastName = "-"
			newUser.id = "-"
		}
		i += 1
	}
	idx += 1
	LoopsRemove()
	GetData(queue[idx])
	return nil
}

func LoopsRemove() {
	a := len(queue)
	control := 0
	for i := 0; i < len(userFriends); i++ {
		for j := 0; j < a; j++ {
			if queue[j] == userFriends[i] {
				control = 1
				break
			}
		}
		if control == 0 {
			queue = append(queue, userFriends[i])
		}
		control = 0
	}
	userFriends = nil
	return
}
