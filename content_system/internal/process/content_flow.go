package process

import (
	"ContentSystem/internal/dao"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"

	flow "github.com/s8sg/goflow/flow/v1"
	goflow "github.com/s8sg/goflow/v1"
)

type ContentFlow struct {
	contentDao dao.ContentDao
}

func ExceContentFlow(db *gorm.DB) {
	contentFlow := &ContentFlow{
		contentDao: *dao.NewContentDao(db),
	}
	fs := goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		WorkerConcurrency: 5,
	}
	//注册工作流 localhost:8080/flow/add-flow
	err := fs.Register("content-flow", contentFlow.flowHandle)
	if err != nil {
		return
	}
	fs.Start()
}
func (c *ContentFlow) flowHandle(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	dag.Node("", c.input)
	dag.Node("verify", c.verify)
	dag.Node("finish", c.finish)
	branches := dag.ConditionalBranch("branches", []string{"category",
		"thumbnail", "format", "pass", "fail"},
		func(bytes []byte) []string {
			var data map[string]interface{}
			if err := json.Unmarshal(bytes, &data); err != nil {
				return nil
			}
			if int(data["approval_status"].(float64)) == 2 {
				return []string{"category", "thumbnail", "pass", "format"}
			}
			return []string{"fail"}
		}, flow.Aggregator(func(m map[string][]byte) ([]byte, error) {
			return []byte("ok"), nil
		}))
	branches["categroy"].Node("category", c.category)
	branches["thumbnail"].Node("thumbnail", c.thumbnail)
	branches["format"].Node("format", c.format)
	branches["pass"].Node("pass", c.pass)
	branches["fail"].Node("fail", c.fail)

	dag.Edge("input", "verify")
	dag.Edge("verify", "branches")
	dag.Edge("branches", "finish")

	return nil
}

func (c *ContentFlow) input(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec inpuot")
	var input map[string]int
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	id := input["content_id"]
	detail, err := c.contentDao.First(id)
	if err != nil {
		return nil, err
	}
	result, err := json.Marshal(map[string]interface{}{
		"title":      detail.Title,
		"content_id": detail.ID,
		"video_url":  detail.VideoURL,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 验证
func (c *ContentFlow) verify(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec verify")
	var detail map[string]interface{}
	if err := json.Unmarshal(data, &detail); err != nil {
		return nil, err
	}
	var (
		title    = detail["title"]
		videoURL = detail["video_url"]
		id       = detail["content_id"]
	)
	//机审 人审
	if int(id.(float64))%2 == 0 {
		detail["approval_status"] = 3
	} else {
		detail["approval_status"] = 2
	}
	fmt.Println(id, title, videoURL)
	return json.Marshal(detail)
}

// 分类
func (c *ContentFlow) category(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec category")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "category", "category")
	if err != nil {
		return nil, err
	}
	return []byte("category"), nil
}

// 裁剪封面图
func (c *ContentFlow) thumbnail(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec category")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "thumbnail", "thumbnail")
	if err != nil {
		return nil, err
	}
	return []byte("thumbnail"), nil
}

// 格式化信息
func (c *ContentFlow) format(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec format")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "format", "format")
	if err != nil {
		return nil, err
	}
	return []byte("format"), nil
}

// 通过
func (c *ContentFlow) pass(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec pass")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "approval_status", 2)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 失败
func (c *ContentFlow) fail(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec fail")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "approval)status", 3)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *ContentFlow) finish(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec finish")
	// var input map[string]interface{}
	// if err := json.Unmarshal(data, &input);err != nil{
	// 	return nil, err
	// }
	// contentID := int(input["content_id"].(float64))
	// err := c.contentDao.UpdateByID(contentID,"approval)status",3)
	// if err != nil{
	// 	return nil, err
	// }
	fmt.Println(string(data))
	return data, nil
}
