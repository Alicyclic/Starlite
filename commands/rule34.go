package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/starlite/modules"
	. "github.com/starlite/utils"
	"github.com/valyala/fasthttp"
)

func Hi() {
	fmt.Print("I'm working... on toast!")
}

// testing phase?

const (
	URL = "https://api.rule34.xxx/index.php?page=dapi&json=1&s=post&q=index&tags="
)

var (
	rule34Cache = make(map[string]Rule34Posts, 0)
)

type Rule34Post struct {
	Preview  string `json:"preview_url"`
	ImageURL string `json:"file_url"`
}

type Rule34Posts []*Rule34Post

func NewHTTPRequest(url string) Rule34Posts {
	posts := make(Rule34Posts, 0)
	_, body, err := fasthttp.Get(nil, URL+url)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &posts)
	return posts
}

func FilterPostsImagesOnly(r Rule34Posts) Rule34Posts {
	for _, post := range r {
		go func(post *Rule34Post) {
			if !IsImage(post.ImageURL) {
				post.ImageURL = post.Preview
			}
		}(post)
	}
	return r
}

func GenerateNewPosts(tags string) Rule34Posts {
	posts := NewHTTPRequest(tags)
	rule34Cache[tags] = posts
	return posts
}

func GetPosts(tags string) Rule34Posts {
	if get, ok := rule34Cache[tags]; ok {
		return get
	}
	return GenerateNewPosts(tags)
}

func (r *Rule34Post) CreateEmbed(tags string) *Embed {
	return NewEmbed().
		SetImage(r.ImageURL).
		SetURL(r.ImageURL).
		SetTitle(strings.Title(tags))
}

func RandomResponse(tags string) *Embed {
	post := GetPosts(tags)
	rand.Seed(time.Now().Unix())
	random := rand.Intn(len(post))
	tags = strings.Replace(tags, "_", " ", -1)
	return post[random].CreateEmbed(tags)
}

func init() {
	modules.NewCommand("rule34", "Search on rule34.xxx").
		AddOption("tags", "tags to search for", 3, true).
		SetHandler(func(c *modules.Command) {
			arg, _ := c.GetOption("tags")
			tags := strings.Replace(arg.StringValue(), " ", "_", -1)
			c.SendNSFWMessage(RandomResponse(tags))
		})

}
