# 8월 28일 발표

### 기본

```jsx
package main
 
import "fmt"
 
func main() {
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
}
```

---

여기 보면 두 번째줄 출력은 e가 아니라 101이 출력된다.

고랭은 프린트할 때 바이트로 출력한다고 한다. 

파이썬에서는 e출력하는데 차이가 있어서 신기했다.

또한 C랑 유사하다고 생각한 부분이 text[10] = “hello”라고 선언되어 있다면,

printf(”%c” or “%d”)에 따라서 101을 출력하거나 e를 출력한다.

### const

```jsx
package main
 
import "fmt"
 
func main() {
    const x string = "Hello World"
    fmt.Println(x)
}
```

```python
const x string = "Hello World"
x = "Some other string"
```

이렇게 const 상수로 x가 정의되면 값을 바꿀 수 없다.

아래 코드에서 x를 “some other string”으로 바꾸려고 하면 컴파일 오류가 발생한다

### var

```python
package main
 
import "fmt"
 
func main() {
    var x string = "Hello World"   #문자열 타입 변수 선언  go는 타입을 뒤로하네?
    fmt.Println(x)
}
```

```python
package main
 
import "fmt"
 
func main() {
    var x string
    x = "first"
    fmt.Println(x)
    x = "second"         # var는 따로 고정 값이 아니라 값을 바꿀 수 있음.
    fmt.Println(x)
}
```

### 변수 추측 선언

```python
x := "Hello World"
```

내가 맨 처음 C언어 공부를 했을 때 타입별로 선언해주는게 진짜 짜증났는데 

go는 알아서 해주는거 정말 좋다. 근데 따로 이렇게 하면 리소스가 더 크나?

이거 때문에 파이썬에서 효율성이 떨어지는걸로 아는데, 똑같이 := 이거 쓰면 효울성 떨어지려나?

### 여러 개 변수 선언

```python
var (
    a = 5
    b = 10
    c = 15
)
```

이거는 C언어 구조체랑 비슷하다고 생각했음. 구조체는 값은 넣어주지는 않지만 변수 선언을 저렇게 하고 “.”을 이용해 사용하긴 하는데 go는 값까지 넣어주네.

  

### 반복문 for

```python
func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```

C언어랑 상당히 유사하다.

### 조건문 if

```python
if i % 2 == 0 {
    // 짝수
} else {
    // 홀수
}

===============================================================

if i % 2 == 0 {
    // 2로 나눌 수 있음
} else if i % 3 == 0 {
    // 3으로 나눌 수 있음
} else if i % 4 == 0 {
    // 4로 나눌 수 있음
}
```

괄호를 따로 사용 안 하는게 파이썬 같네.

### switch

```python
switch i {
case 0: fmt.Println("영")
case 1: fmt.Println("일")
case 2: fmt.Println("이")
case 3: fmt.Println("삼")
case 4: fmt.Println("사")
case 5,6: fmt.Println("오")
default: fmt.Println("알 수 없는 숫자")
}
```

저번에 다들 break가 없어도 된다고 했었음. 듣기로 case 1,2가 같이 사용될 수 있다 했는데 맞나??