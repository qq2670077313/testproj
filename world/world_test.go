package world

import (
	"errors"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Log(64 << 10) // 65536   左移一位 x2
	t.Log(2 << 10)  // 2048  2的10次方
	t.Log(64 * 32 * 32)

	m := make(map[string]string)
	// m["a"] = "1"
	// m["b"] = "1"
	m[""] = "333"
	// m["333"] = ""

	getV := func() map[string]string {
		m1 := make(map[string]string)
		for k, v := range m {
			m1[v] = k
		}
		return m1
	}
	getV()

	setV := func(k, v string) { /// 去重
		// mt := getV()
		// for mt[v] != "" { /// [:123] 判断错误
		// 	delete(m, mt[v])
		// 	mt = getV()
		// }

		for key, val := range m {
			t.Log(key)
			t.Log(val)
			if val == v {
				delete(m, key)
				t.Log(m)
			}
		}
		m[k] = v
	}
	t.Log(m)
	// setV("d", "2")
	setV("", "123")
	// setV("f", "1")
	// delete(m, "c")
	t.Log(m)

	mm := make(map[string]string)
	mm["f"] = "1"

	// t.Log(mm["1"])
	// delete(m, mm["1"]) // 不存在是 空的
	// t.Log(m)

	result := Greet()
	if result != "Hello GitHub Actions" {
		t.Errorf("Greet() = %s; Expected Hello GitHub actions", result)
	}
}

func TestFunc1(t *testing.T) {
	errF := func(i int) (int, error) {
		return i, errors.New("aa")
	}

	i, err := errF(1)
	t.Logf("i=%v, i=%p, err=%p", i, &i, &err)
	// var j int
	// j, err = errF(2) // err同一个地址
	// t.Logf("j=%v, j=%p, err=%p", j, &j, &err)
	j, err := errF(2) // err同一个地址
	t.Logf("j=%v, j=%p, err=%p", j, &j, &err)
	if _, err = errF(2); err != nil {
		t.Logf(" err=%p", &err)
	}
}

func TestErrorExclude(t *testing.T) {
	tests := []int{-1, 0, 1}
	for i := 0; i < len(tests); i++ {
		err := ErrorExclude(tests[i])
		if err != nil {
			t.Log(err)
		}
	}
}

func TestErrorExcludeReturn(t *testing.T) {
	err := ErrorExcludeReturn(0)
	if err != nil {
		t.Log(err)
	}
}
