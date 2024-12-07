package flow

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	flow "github.com/s8sg/goflow/flow/v1"
	goflow "github.com/s8sg/goflow/v1"
)

// DAG加工流
// - 调度
// - 工作流程
// - 组织依赖
func Input(data []byte, option map[string][]string) ([]byte, error) {
	var input map[string]int
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	outputInt := input["input"]
	return []byte(strconv.Itoa(outputInt)), nil
}

func AddOne(data []byte, option map[string][]string) ([]byte, error) {
	num, _ := strconv.Atoi(string(data))
	outputInt := num + rand.Intn(10) + 1
	return []byte(strconv.Itoa(outputInt)), nil
}

func AddTwo(data []byte, option map[string][]string) ([]byte, error) {
	num, _ := strconv.Atoi(string(data))
	outputInt := num + rand.Intn(100) + 1
	return []byte(strconv.Itoa(outputInt)), nil
}

// 聚合节点
func Aggregator(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("aggretator = ", string(data))
	return data, nil
}

func Expand10(data []byte, option map[string][]string) ([]byte, error) {
	num, _ := strconv.Atoi((string(data)))
	outputInt := num * 10
	fmt.Println("expand10 = ", outputInt)
	return []byte(strconv.Itoa(outputInt)), nil
}

func Expand100(data []byte, option map[string][]string) ([]byte, error) {
	num, _ := strconv.Atoi((string(data)))
	outputInt := num * 100
	fmt.Println("expand 100 = ", outputInt)
	return []byte(strconv.Itoa(outputInt)), nil
}

func Output(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("data = ", string(data))
	return []byte("ok"), nil
}

func MyFlow(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	//构建节点
	dag.Node("input", Input)
	dag.Node("add-one", AddOne)
	dag.Node("add-two", AddTwo)
	//聚合
	dag.Node("aggregator", Aggregator, flow.Aggregator(func(m map[string][]byte) ([]byte, error) {
		num1, _ := strconv.Atoi(string(m["add-one"]))
		num2, _ := strconv.Atoi(string(m["add-two"]))
		sum := num1 + num2
		return []byte(strconv.Itoa(sum)), nil
	}))
	//分支
	branches := dag.ConditionalBranch("judge", []string{"moreThan", "lessThan"}, func(bytes []byte) []string {
		num, _ := strconv.Atoi(string(bytes))
		fmt.Println("conditionBranch = ", num)
		if num > 10 {
			return []string{"moreThan"}
		}
		return []string{"lessThan"}
	}, flow.Aggregator(func(m map[string][]byte)([]byte, error){
		if v, ok := m["moreThan"];ok{
			return v,nil
		}
		if v, ok := m["lessThan"];ok{
			return v, nil
		}
		return nil, nil
	}))
	branches["moreThan"].Node("expand-10", Expand10)
	branches["lessThan"].Node("expand-100", Expand100)
	//输出
	dag.Node("output", Output)
	//构建依赖关系
	dag.Edge("input", "add-one")
	dag.Edge("input", "add-two")
	dag.Edge("add-one", "aggregator")
	dag.Edge("add-two", "aggregator")
	dag.Edge("aggregator","judge")

	return nil
}

func main() {
	fs := goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		WorkerConcurrency: 5,
	}
	//注册工作流 localhost:8080/flow/add-flow
	err := fs.Register("add-flow", MyFlow)
	if err != nil {
		return
	}
	err = fs.Start()
	if err != nil {
		panic(err)
	}

}
