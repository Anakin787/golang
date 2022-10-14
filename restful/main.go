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
    Seq         int    `json:"seq"`
    ID          string `json:"id"`
    PW          string `json:"pw"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    Hp          string `json:"hp"`
    Role        int    `json:"role"`
    State       int    `json:"state"`
    Description string `json:"des"`
}

type Login struct {
    ID string `json="id"`
    PW string `json="pw"`
}

type Input struct {
    ID    string `json:"id"`
    PW    string `json:"pw"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Hp    string `json:"hp"`
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
    data, err := os.ReadFile("./account.json") // 파일가져오기
    if err != nil {
        fmt.Printf("Failed to read file: %v\n", err)
    }
    // u:=[]Account{}
    // var u =new(Account)
    var input Input
    var num, _ = strconv.Atoi(c.Params.ByName("id"))
    err = json.Unmarshal(data, &u) // json을 구조체형으로 변환
    // if u.ID!=""{
    //     u
    // }
    u[num-1].ID = input.ID
    u[num-1].PW = input.PW
    u[num-1].Name = input.Name
    u[num-1].Hp = input.Hp
    u[num-1].Email = input.Email

    c.JSON(http.StatusOK, gin.H{"result": u[num-1]})
}

func LoginCheck(c *gin.Context) {
    data, _ := os.ReadFile("./account.json")
    var u []Account
    var login Login

    _ = json.Unmarshal(data, &u)

    if err := c.ShouldBind(&login); err != nil {
        c.JSON(401, gin.H{"Error": err})
    }
    for index := range u {
        if u[index].ID == login.ID && u[index].PW == login.PW {
            c.JSON(http.StatusOK, gin.H{"Name is": u[index].Name})
            // num++
            break
        } else {
            if index == 19 { //마지막 까지 확인했을때 일치하는게 없으면 출력
                c.JSON(401, "아이디 또는 비밀번호를 다시 확인하세요.")
            }
        }
    }
}

func main() {
    r := gin.Default()
    r.GET("/account", GetList)
    r.GET("/account/:id", FindMembers)
    r.POST("/account/login", LoginCheck)
    r.PUT("/account/:id", UpdateMember) //파일받아와서 특정회원정보 불러와서 입력받아 수정해가지고 다시 저장
    // r.POST("/account", CreateMember)
    // r.GET("/account/:id", DeleteMembers)

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
