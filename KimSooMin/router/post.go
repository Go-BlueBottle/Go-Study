package router

import(
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/kssumin/golang_board/Repository"

)


func GetPosts(c *gin.Context) {
	posts := Repository.post.GetPostsByDB()
	fmt.Println(posts)
	c.JSON(200, posts)
}

func GetPost(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	post := Repository.post.GetPostByDB(int(id))
	fmt.Println(post)
	c.JSON(200, post)
}



