package rubychina

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// https://ruby-china.org/api/users/no13bus/topics/favorite.json
//得到的是一个topics类型
const ruby_chinaAPI = "https://ruby-china.org/api/"

type Users []User
type Topics []Topic
type Replies []Reply
type Nodes []Node
type Sites []Site
type AllSites []ASites

type User struct {
	Id            uint32 `json:id`
	Name          string `json:name`
	Login         string `json:login`
	Location      string `json:location`
	Company       string `json:company`
	Twitter       string `json:twitter`
	Website       string `json:website`
	Bio           string `json:bio`
	Tagline       string `json:tagline`
	Github_url    string `json:github_url`
	Email         string `json:email`
	Gravatar_hash string `json:gravatar_hash`
	Avatar_url    string `json:avatar_url`
	Topics        Topics `json:topics`
}

type Topic struct {
	Id                    uint32  `json:id`
	Title                 string  `json:title`
	Created_at            string  `json:created_at`
	Updated_at            string  `json:updated_at`
	Replied_at            string  `json:replied_at`
	Replies_count         uint32  `json:replies_count`
	Node_name             string  `json:node_name`
	Node_id               uint32  `json:node_id`
	Last_reply_user_id    uint32  `json:last_reply_user_id`
	Last_reply_user_login string  `json:last_reply_user_login`
	User                  User    `json:user`
	Body                  string  `json:body`
	Body_html             string  `json:body_html`
	Hits                  uint32  `json:hits`
	Replies               Replies `json:replies`
}

type Reply struct {
	Id         uint32 `json:id`
	Body       string `json:body`
	Body_html  string `json:body_html`
	Created_at string `json:created_at`
	Updated_at string `json:updated_at`
	Deleted_at string `json:deleted_at`
	Topic_id   uint32 `json:topic_id`
	User       User   `json:user`
}

type Site struct {
	Url        string `json:url`
	Name       string `json:name`
	Favicon    string `json:favicon`
	Desc       string `json:desc`
	Created_at string `json:created_at`
}

// _Id 如何解析? Sites_count如果没有的话 解析出来的值为0
type ASites struct {
	Name        string `json:name`
	Sites       Sites  `json:sites`
	_Id         uint32 `json:_id`
	Sites_count uint32 `json:sites_count`
}

type Node struct {
	Id           uint32 `json:id`
	Name         string `json:id`
	Topics_count uint32 `json:topics_count`
	Summary      string `json:summary`
	Section_id   uint32 `json:section_id`
	Section_name string `json:section_name`
}

func GetJsonData(url string, v interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal(data, v)
	return

}

// 取得所有node信息
func GetNodes() (nodes Nodes, err error) {
	url := ruby_chinaAPI + "nodes.json"
	err = GetJsonData(url, &nodes)
	return
}

// 取得最近的topics信息
func GetTopics() (topics Topics, err error) {
	url := ruby_chinaAPI + "topics.json"
	err = GetJsonData(url, &topics)
	return

}

// 取得用户们的信息
func GetUsers() (users Users, err error) {
	url := ruby_chinaAPI + "users.json"
	err = GetJsonData(url, &users)
	return
}

// 按照用户的名字取得用户信息
func GetUserByName(name string) (user User, err error) {
	url := ruby_chinaAPI + "users/" + name + ".json"
	err = GetJsonData(url, &user)
	return
}

// 按照帖子的id取得帖子的相关信息
func GetTopicByID(id uint32) (topic Topic, err error) {
	url := ruby_chinaAPI + "topics/" + strconv.Itoa(int(id)) + ".json"
	err = GetJsonData(url, &topic)
	return
}

// 取得用户收藏的帖子
// https: //ruby-china.org/api/users/no13bus/topics/favorite.json
func GetFavoriteTopic(username string) (topics Topics, err error) {
	url := ruby_chinaAPI + "users/" + username + "/topics/favorite.json"
	err = GetJsonData(url, &topics)
	return
}

// 取得sites的信息
func GetSites() (sites AllSites, err error) {
	url := ruby_chinaAPI + "sites.json"
	err = GetJsonData(url, &sites)
	return

}
