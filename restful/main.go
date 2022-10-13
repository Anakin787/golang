package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Account struct {
    Seq         int    `json:"seq" uri:"id"`
    ID          string `json:"id"`
    Pw          string `json:"pw"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    Hp          string `json:"hp"`
    Role        int    `json:"role"`
    State       int    `json:"state"`
    Description string `json:"des"`
}

func GetList(c *gin.Context) {
    data, err := os.ReadFile("./account.json")
    if err != nil {
        fmt.Printf("Failed to read file: %v\n", err) // 디버그할땐 출력해보는게 좋다
    }
    // fmt.Printf("xxxxxxxxxxxxxxxxxx json: %v", data)

    var u []Account
    // err = json.Unmarshal(b, &u) //NewDecoder문법과 같다
    err = json.NewDecoder(bytes.NewBuffer(data)).Decode(&u)
    if err != nil {
        fmt.Printf("Failed to unmarshal: %v\n", err)
    }

    c.JSON(http.StatusOK, u)
}

func FindMembers(c *gin.Context) {
    data, err := os.ReadFile("./account.json") // os와 util의차이점은 취약점관련 버전에따른 이름차이(os를 주로사용)
    if err != nil {                            // 에러처리
        fmt.Printf("Failed to read file: %v\n", err)
    }

    var u []Account
    err = json.Unmarshal(data, &u) // unmarshal이 byte형태의 json을 구조체형태로 변환해줌
    if err != nil {                // 에러처리
        fmt.Printf("Failed to unmarshal: %v\n", err)
    }

    var num = c.Param("id")
    seq, _ := strconv.Atoi(num)

    c.JSON(http.StatusOK, u[seq-1])
}

func UpdateMember(c *gin.Context) {
    data, err := os.ReadFile("./account.json")
    if err != nil {
        fmt.Printf("Failed to read file: %v\n", err)
    }
    var u []Account // 구조체안에는 하나의객체가 아니라 배열로 이루어져있기때문에 '[]구조채명' 으로 선언한다.
    err = json.Unmarshal(data, &u)

    // 입력받은 seq의 특정항목을 수정한후 저장

    id, err := strconv.Atoi(c.Params.ByName("id"))
    if err != nil {
        return
    }
    c.JSON(http.StatusOK, gin.H{"result": id})
    // if err := c.BindJSON(&u); err != nil {
    //     utils.Response(c, utils.BAD_REQUEST, "Bad Params.")
    //     panic(err)
    // }

    // db := utils.DbConnect()

    // if res := db.Model(&Account{}).Where("seq = ?", seq).Updates(u); res.Error != nil {
    //     utils.Response(c, utils.INTERNAL_SERVER_ERROR, "Query Error.")
    //     panic(res.Error)
    // }

    // Response(c, utils.CREATED, "OK")
}

func main() {
    r := gin.Default()
    r.GET("/account", GetList)
    r.GET("/account/:id", FindMembers)
    r.PUT("/account/:id", UpdateMember)
    // r.POST("/account", CreateMember)
    // r.GET("/account/:id", DeleteMembers)
    // r.POST("/account/login", LoginCheck)

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
