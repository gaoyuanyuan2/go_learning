package try_test
import "testing"

func TestFibList(t *testing.T){
    var a int =1
    var b int =1
    t.Log(a)
    for i := 0;  i< 5; i++ {
      t.Log(" ",b)
      tmp := a
      a = b
      b = tmp + a
    }
    t.Log();
}

func TestExchange(t *testing.T){
    a :=1
    b :=2
    // temp := b
    // b = a
    // a = temp
    a, b = b, a
    t.Log(a,b);
}
