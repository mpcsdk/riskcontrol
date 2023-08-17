package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/bilibili/gengine/engine"
)

const rulePool = `
rule "rulePool" "rule-des" salience 10
begin
sleep()
print("do ", FunParam.Name)

// FunParam.Name = "newName"
fname = FunParam.Name
// return FunParam
return fname
end `

type FunParam struct {
	Name    string
	Version int
}

func Sleep() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(1000)
	time.Sleep(time.Nanosecond * time.Duration(i))
}

func main() {
	Sleep()
	apis := make(map[string]interface{})
	apis["print"] = fmt.Println
	apis["sleep"] = Sleep
	pool, e1 := engine.NewGenginePool(1, 3, 2, rulePool, apis)
	if e1 != nil {
		panic(e1)
	}
	g1 := int64(0)
	g2 := int64(0)
	g3 := int64(0)
	g4 := int64(0)
	g5 := int64(0)
	cnt := int64(0)

	go func() {
		for {
			param := map[string]interface{}{}
			f := FunParam{Name: "func1", Version: 1}
			// param["FunParam"] = f
			e2, rst := pool.ExecuteSelectedRules(param, []string{"rulePool"})
			// pool.UpdatePooledRules()
			// pool.UpdatePooledRulesIncremental()
			if e2 != nil {
				println(fmt.Sprintf("e2: %+v", e2))
			}
			fmt.Println(rst["rulePool"])
			f.Version = 2
			fmt.Println(rst["rulePool"])
			fmt.Println(param["FunParam"])
			time.Sleep(1 * time.Second)
			atomic.AddInt64(&cnt, 1)
			g1++
		}
	}()

	// go func() {
	// 	for {
	// 		param := &FunParam{Name: "func2"}
	// 		e2 := pool.ExecuteRules("FunParam", param, "", nil)
	// 		if e2 != nil {
	// 			println(fmt.Sprintf("e2: %+v", e2))
	// 		}
	// 		//time.Sleep(1 * time.Second)
	// 		atomic.AddInt64(&cnt, 1)
	// 		g2++
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		param := &FunParam{Name: "func3"}
	// 		e2 := pool.ExecuteRules("FunParam", param, "", nil)
	// 		if e2 != nil {
	// 			println(fmt.Sprintf("e2: %+v", e2))
	// 		}
	// 		//time.Sleep(1 * time.Second)
	// 		atomic.AddInt64(&cnt, 1)
	// 		g3++
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		param := &FunParam{Name: "func4"}
	// 		e2 := pool.ExecuteRules("FunParam", param, "", nil)
	// 		if e2 != nil {
	// 			println(fmt.Sprintf("e2: %+v", e2))
	// 		}
	// 		//time.Sleep(1 * time.Second)
	// 		atomic.AddInt64(&cnt, 1)
	// 		g4++
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		param := &FunParam{Name: "func5"}
	// 		e2 := pool.ExecuteRules("FunParam", param, "", nil)
	// 		if e2 != nil {
	// 			println(fmt.Sprintf("e2: %+v", e2))
	// 		}
	// 		//time.Sleep(1 * time.Second)
	// 		atomic.AddInt64(&cnt, 1)
	// 		g5++
	// 	}
	// }()
	// 主进程运行5秒
	time.Sleep(5 * time.Second)
	// 统计各个子进程分别运行次数
	println(g1, g2, g3, g4, g5)
	// 统计在引擎池下总的各个子进程总的运行测试
	println(g1+g2+g3+g4+g5, cnt)
}
