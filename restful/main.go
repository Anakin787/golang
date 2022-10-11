package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type fooHandler struct{}

type User struct {
    First string    `json:"first_name"` //Annotation 표기 : json에서는 이렇게 쓴다(Go에서는 _ 사용 지양)
    Last  string    `json:"last_name"`
    Email string    `json:"email"`
    Creat time.Time `json:"creat"`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //메소드 생성(참조)
    user := new(User)
    err := json.NewDecoder(r.Body).Decode(user)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprint(w, err)
        return
    }
    user.Creat = time.Now()
    data, _ := json.Marshal(user)
    w.Header().Add("content-type", "application/json") //서버에서도 json파일이라고 알려주는것
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, string(data))
}

func bar(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "world"
    }
    fmt.Fprintf(w, "Hello %s", name)
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // HandleFunc 함수를 이용하면 내가 원하는 경로를 함수와 연결(정적) , function으로 직접등록
        fmt.Fprint(w, "hello")
    })
    mux.HandleFunc("/bar", bar)

    mux.Handle("/foo", &fooHandler{}) //참조메서드 호출 , 인스턴스형태

    http.ListenAndServe(":3000", mux)
}
