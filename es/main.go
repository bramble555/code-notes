package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

type DemoModel struct {
	Id         uint   `json:"id"`
	UserName   string `json:"user_name"`
	CreateTime string `json:"create_time"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

func (u *DemoModel) Index() string {
	return "demo_index"
}
func (u *DemoModel) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "title": { 
        "type": "text" // 查询的时候是分词匹配
      },
      "key": { 
        "type": "keyword" // 完整匹配
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}
func main() {
	es = Init()
	fmt.Println(es)
	d := DemoModel{}
	d.CreateIndex()
	// fmt.Println(err)
	// err = d.DeleteIndex()
	// fmt.Println(err)
	// ds := DemoModel{
	// 	Id:         1,
	// 	UserName:   "lc",
	// 	CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	// 	Title:      "今天天气不好",
	// 	Content:    "今天天气不好没我很难过",
	// }
	// CreateDoc(ds)
	FindDoc()
}
func CreateDoc(d DemoModel) {
	ir, err := es.Index().Index(d.Index()).BodyJson(d).Do(context.Background())
	if err != nil {
		log.Panic("添加 doc 出错了")
	}
	fmt.Printf("doc:%#v\n", ir)
}
func FindDoc() {
	d := DemoModel{}
	query := elastic.NewBoolQuery()
	res, err := es.
		Search(d.Index()).       // 指定索引名
		Query(query).            // 查询条件
		From(0).                 // 分页操作
		Size(10).                // 每页大小
		Do(context.Background()) // 执行查询

	if err != nil {
		log.Println("Error executing search:", err)
		return
	}
	count := res.Hits.TotalHits.Value
	fmt.Println("Total Hits:", count)

	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}

func (d *DemoModel) CreateIndex() error {
	exist := d.IndexExists()
	// 创建索引
	if exist {
		return nil
	}
	res, err := es.CreateIndex(d.Index()).BodyString(d.Mapping()).Do(context.Background())
	if !res.Acknowledged {
		log.Printf("创建失败\n")
		return errors.New("创建索引失败")
	}
	return err
}

func (d *DemoModel) DeleteIndex() error {
	exist := d.IndexExists()
	// 删除索引
	if !exist {
		return nil
	}
	_, err := es.DeleteIndex(d.Index()).Do(context.Background())
	if err != nil {
		log.Fatalf("es.DeleteIndex(d.Index()) err:%s\n", err.Error())
	}
	return nil
}
func (d *DemoModel) IndexExists() bool {
	index := d.Index()
	exist, err := es.IndexExists(index).Do(context.Background())
	if err != nil {
		log.Printf("err:%s\n", err.Error())
		return true
	}
	return exist
}

var es *elastic.Client

func Init() *elastic.Client {

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}
