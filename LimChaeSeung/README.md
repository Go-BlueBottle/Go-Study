# 1장 시작하기
Golang은 고급기능과 함께 명확한 문법을 가진 언어이다. 이로 인해 소프트웨어에 관한 능력은 매우 직관적인데, 다음과 같은 기능에 매우 효과적이다.
1. 요구사항 수집
2. 해결책 모색
3. 소스코드를 작성해 해결채글 구현
4. 소스코드를 실행파일로 컴파일
5. 동작 여부를 확인하기 위한 프로그램 실행과 테스트
__기본적인 내용은 생략하기로 했다__

### 1.4 Go 도구
Go는 컴파일 도구이다. 따라서 Go 프로그래밍을 진해하기 위해서는 컴파일러가 필요하다.
설치는 다음과 같은 [링크](http://www.golang.org)에서 설치 할 수 있다.
go version으로 설치가 되어있는지 확인 할 수 있다. 

<img width="311" alt="go-version" src="https://user-images.githubusercontent.com/45393030/187080471-b4b3036b-d081-4599-a81f-c5836bc5c28e.png">

# 2장 첫 번째 프로그램
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
다음과 같이 언어의 입문 단계인 Hello Wrold이다.

## 2.1 Go 프로그램을 읽는 법
1. `package main` main이라는 패키지로 묶겠다는 의미이다.
2. `import "fmt"` fmt라는 패키지를 가지고 오겠다는 의미이다.
3. `func main()` Entry point이다.
4. `fmt.Println("Hello World")` Hello World를 출력한다.

프로그래밍을 하다보면 document가 필요할 때가 있는데 이는 
godoc fmt Println같이 godoc으로 사용법을 쉽게 알 수 있다.

# 3장 타입
Go는 TypeSystem을 가지고 있는 언어이다.

## 3.1 숫자
Go에서는 숫자를 표현하는 여러 다양 한 타입을 제공한다.
정수: `uint8` = `byte`, `uint16` = `rune`, `uint32`, `uint64`, `int8`, `int16`, `int32`, `int64`
부동 소수점: `float32`, `float64`, `complex64`, `complex128`

## 3.2 문자열
Go의 문자열은 개별 바이트로 구성되어 있다. 문자열은  다음과 같이 "Hello World"나, \`Hello World\`처럼 쓸 수 있다. \`로 감싼 문자열은 \\t, \\n처럼 이스케이프 문자열을 포함하여 나타낼 수 있다.

``` go
package main

import "fmt"

func main() {
	fmt.Println(len("Hello World")) -- 1
	fmt.Println("Hello World"[1]) -- 2
	fmt.Println("Hello " + "World") -- 3
}
```
1. Hello World의 총 수는 11자이다. (공백도 포함된다)
2. [1]이므로 선택되는 것은 e이며, 출력되는 것은 101이다. 
3. 문자열끼리 더하여 표시 할 수 있다.
 
## 3.2 불린
true와 false를 나타내는 1비트 정수 타입이다.
&&, ||, !의 3가지 연산자가 있다.

# 4장 변수
```go
package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)
}
```
위와 같이 `var 이름 타입`으로 변수를 선언할 수 있다.
만약 타입이 추론 가능하다면 `이름 := 값`으로 생략할 수 있다.
## 4.3 상수
Go에서는 상수를 지원하느데 `var` 키워드 대신 `const`를 사용한다.
```go
package main

import "fmt"

func main() {
	const x string = "Hello World"
	fmt.Println(x)
}
```
만약 상수를 다음과 같이 변경하려고 하면 오류가 난다.
```go
package main

import "fmt"

func main() {
	const x string = "Hello World"
	x = "Hello Go"
	fmt.Println(x)
}
```
## 4.4 여러개의 변수 정의
만약 변수가 여러개 필요한 경우 축약형으로 지원 해준다.
```go
var ( 
	a = 5
	b = 10
	c = 15
)
```

# 5장 제어 구조
## 5.3 For
```go
func main() {
	for i: = 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
```
## 5.4 If
```go
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "짝수")
	} else {
		fmt.Println(i, "홀수")
	}
}
```
## 5.4 Switch
```go
switch {
	case 0: fmt.Println("영")
	case 1: fmt.Println("일")
	case 2: fmt.Println("이")
	case 3: fmt.Println("삼")
	case 4: fmt.Println("사")
	case 5: fmt.Println("오")
	case 6: fmt.Println("육")
	default: fmt.Println("알 수 없는 숫자")
}
```
# 6장 배열, 슬라이스, 맵
## 6.1배열
배열은 다음과 같이 표현한다.
```go
var x [5] int
```
배열을 순환할 때 다음과 같이 순환하면 좋다
```go
var total float 64 = 0
for _, value := range x {
	total += value
}
fmt.Println(total / float64(len(x)))
```
배열을 저장할 때 다음과 같이 한번에 저장 할 수 있다.
```go
x := [5]float64(98, 03, 77, 82, 83)
```
## 6.2 슬라이스
다음과 같이 슬라이스 할 수 있다.
```go
x := make([]float64, 5)
x := make([]float64, 5, 10)
```
이를 줄이는 방법이 있다.
```go
arr := []float64{1, 2, 3, 4, 5}
x := arr[0:5]
```
여기서 배열을 더 활용 할 수 있다.
다음은 append에 관한 내용이다.
```go
func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2) // 12345
}
```
다음은 copy에 관한 내용이다.
```go
func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2) // 123, 12
}
```
## 6.3맵
맵을 표현하는 방법은 다음과 같다.
```go
var x map[string]int
```
맵에 데이터를 넣는 방법은 다음과 같다.
```go
x := make(map[string]int)
x["key"] = 10
fmt.Println(x)
```
맵에서 없을 때 0으로 리턴해 준다. Go에서는 좀 더 좋은 방법으로 나타낸다.
```go
name, ok := element["Ok?"]
fmt.Println(name, ok) // "", false
```
이를 이용하여 Go에서는 다음과 같은 형태로 나타낸다.
```go
if name, ok := element["Ok?"]; ok {
	fmt.Println(name, ok)
}
```
# 7장 함수
## 7.2 다중값 반환
```go
func f() (int, int) {
	return 1, 2
}
func main() {
	x, y := f()
}
```
## 7.3 가변 함수
여러개 받을 때 사용 된다.
```go
func add (args ...int) int {
	total := 0
	for _, v := range args {
		total += v	
	}
}
func main() {
	fmt.Println(add(1, 2, 3))
}
```
배열을 펼쳐서 전달할 때 사용 된다.
```go
func main() {
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}
```
## 7.4 클로저
함수 안에 함수를 만들 수 있다.
```go
func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1)) // 2
}
```
함수를 반환하는 함수를 작성할 수 있다.
```go
func addOne () func () uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 1
		return
	}
}
func main() {
	nextEvent := addOne()
	fmt.Println(nextEvent()) // 0
	fmt.Println(nextEvent()) // 1
	fmt.Println(nextEvent()) // 2
}
```
## 7.5 재귀
```go
func factorical (x uint) uint {
	if x == 0 {
		return 
	}
	return x * factorial(x - 1)
}
```
## 7.6 지연, 패닉, 복구
defer라는 명령어가 있는데, 이는 기본적으로 마지막에 지연으로 끝낼 때 쓰는 문법이다.
```go
package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()
}
```
위와 같이 쓰면 1st 2nd가 순서대로 나온다.
defer은 일반적으로 자원을 해제해야 할 때 쓴다.
```go
db, _ := mysql.Open('DBPath')
defer db.Close()
```
### 패닉과 복구
`panic`은 오류를 일으키는 함수이다. `recover`함수를 이용하면 런타임 패닉을 처리 할 수 있다.
```go
package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	} ()
	panic("PANIC")
}
```
다음과 같이 복구는 defer와 함께 써야 처리가 된다.
# 8장 포인터
## 8.1 \*와 & 연산자
Go에서는 포인터를 \*로  나타낸다. 또한 역참조를 할 수 있는데 &로 나타내기도 한다.
```go
func zero(xPtr *int*) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // 0
}
```
## 8.2 new
```go
func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // 1
}
```
`new`는 인자로 타입을 받아 해당 타입의 값에 맞는 충분한 메모리를 할당한 후 그것에 대한 포인터를 반환한다.

# 9장 구조체와 인터페이스
## 9.1 구조체
구조체란 이름이 지정된 필드가 포함된 타입이다.
``` go
type User struct {
	name string
	email string
	sex string
}
```
이는 다음과 같이 겹쳐서 쓸 수 있다.
``` go
type User struct {
	name, email, sex string
}
```

### 초기화
새 User 타입은 다양한 방법으로 인스턴스를 생성할 수 있다.
```go
var c User
c := new(User)
c := User(name: "임채승", email: "cotmd6203@naver.com", sex: "male")
c := User("임채승", "cotmd6203@naver.com", sex: "male")
```
### 필드
`.`연산사를 이용한다면 필드에 접근할 수 있다.
```go
fmt.Println(c.name, c.email) // 임채승 cotmd6203@naver.com
```
참고로 struct는 항상 값을 복사해서 전달한다. 그래서 앞장에서 나온 포인터와 같이 많이 쓰인다.
```go
func changeName(c *User, n string) string {
	c.name = n
	return n
}
```
## 9.2메서드
재미있게도 nodejs와 매우 비슷한 점이 있는데 메서드를 추가할 수 있다는 것이다.
```go
func (c *User) changeName(n: string) string {
	c.name = n
	return n
}

c := User("임채승", "cotmd6203@naver.com", sex: "male")
c.changeName("loopy")
fmt.Println(c.name) // loopy
```
### 포함 타입
구조체는 보통 `has-a`관계이다. 하지만 `is-a`처럼 쓰일 수도 있다. 다음과 같은 예시를 통해서 보자.
```go
type User struct {
	name string
	email string
	sex string
}

func (c *User) changeName(n: string) string {
	c.name = n
	return n
}

type Producer struct {
	User
	userType string
}
```
위와 같이 정의 되어 있을 때, `Producer`은 2가지 방법으로 `changeName(c: string)`에 접근할 수 있다.
```go
p := new(Producer)
p.User.name
// 또는
p.name
```
즉 go에서는 자동으로 `is-a`처럼 쓸 수 있게 해준다.
## 9.3 인터페이스
```go
type Thing interface {
	getName() string
}
```
위와 같이 물건(혹은 인물)에 대해서 이름을 가지고 있는 받는 인터페이스가 있다고 가정하자.
이는 다음과 같이 표현하여, 다른 `struct`에서 `getName`을 구현한다면, `thing`이라는 인터페이스를 구현한 것이다. 그래서 다음과 같이 여러개의 물체 혹은 인물에 대해서 이름들을 구하려면 다음과 같이 구현할 수 있다는 것이다.
```go
func getAllNames(things ...Thing) string {
	var names string
	for _, n := range things {
		names += n.getName()
	}
	return names
}
```
재미 있는 점은 인터페이스는 필드로도 사용 가능하다는 것이다.
# 10장 동시성
go에서는 고루틴(goroutine)과 채널(channel)을 통해 동시성을 지원한다.
## 10.1 고루틴
```go
package main

import "fmt"

func f(n int) {
	for i:= 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}
```
고루틴이 동시에 실행되는 것이 아니라 순서대로 실행되는 것처럼 보일 수 있다. 하지만 그렇지 않다는 것을 보자.
```go
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
```
위와 같이 실행이 되면 무작위의 시간마다 다르게 동작하는데, 이를 확인할 수 있다.
## 10.2 채널
채널(channel)은 두 고루틴이 서로 통신하고 실행흐름을 동기화하는 수단을 제공한다. 채널은  `chan`으로 나타낸다.
```go
package main

import (
	"fmt"
	"time"
)

func econo(c chan string) {
	for i:= 0; ;i++ {
		c <- "econo"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	
	var input string
	fmt.Scanln(&input)
}
```
### 채널 방향
채널 타입에 방향을 지정할 수 있다. c를 보내기만 하려면 `func pinger(c chan<- string)`같이 쓰고, c를 받기만 하려면 `func pinger(c <-chan string`로 작성해야 한다.
### Select
Go에는 `switch`와 비슷하지만 채널에 대해서만 동작하는 select라는 특별한 구문이 있다.
```go
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
				case msg1 := <- c1:
				fmt.Println(msg1)
				case msg2 := <- c2:
				fmt.Println(msg2)
			}
		}
	}()
	
	var input string
	fmt.Scanln(&input)
}
```
### 버퍼 채널
채널을 만들 때 `make`함수에 두 번째 매개변수를전달하는 것도 가능하다.
```go
c := make(chan int, 1)
```
위와같이 용량이 1인 버퍼 채널이 만들어진다. 보통 채널은 동기적으로 동작한다. 반면 버퍼 채널은 비동기적이다. 즉, 메시지를 보내거나 받을 때 채널이 이미 꽉 차있지 않는 이상 기다리지 않는다.
