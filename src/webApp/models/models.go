package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	// 设置数据库路径
	_DB_NAME = "data/webApp.db"
	//设置数据库名称
	_SQLITE3_DRIVER = "sqlite3"
)

//首页简单音乐推荐
type Music struct {
	Id       int64
	Title    string
	Singer   string
	Src      string
	Img      string
	Time     string
	Star     string
	PlayTime int64
}

//分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Lables          string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	//检查数据库文件
	_, err := os.Stat(_DB_NAME)
	if err != nil {
		fmt.Println("数据库不存在，创建")
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	// 注册模型
	orm.RegisterModel(new(Music), new(Category), new(Topic))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

/**
 * 获取所有的音乐列表
 */
func GetAllMusic() (musicList []*Music, err error) {
	o := orm.NewOrm()

	musicList = make([]*Music, 0)

	qs := o.QueryTable("music")

	_, err = qs.All(&musicList)

	return musicList, err
}

/**
 * 获取所有的文章列表
 * @param category 种类
 * @param lable 标签
 * @param isDesc 是否筛选
 */
func GetAllTopics(category, lable string, isDesc bool) (topics []*Topic, err error) {
	o := orm.NewOrm()

	topics = make([]*Topic, 0)

	qs := o.QueryTable("topic")
	if isDesc {
		if len(category) > 0 {
			qs = qs.Filter("category", category)
		}
		if len(lable) > 0 {
			qs = qs.Filter("lables__contains", "$"+lable+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)

	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

/**
 * 添加文章
 * @param title 标题
 * @param category 种类
 * @param label 标签
 * @param content 内容
 * @param attachment 附件
 */
func AddTopic(title, category, lable, content, attachment string) error {
	// 处理标签
	lable = "$" + strings.Join(strings.Split(lable, " "), "#$") + "#"

	o := orm.NewOrm()

	topic := &Topic{
		Title:      title,
		Category:   category,
		Lables:     lable,
		Content:    content,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(), //TODO这里有问题
	}
	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	// 更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		// 如果不存在我们就直接忽略，只当分类存在时进行更新
		cate.TopicCount++
		_, err = o.Update(cate)
	}

	return err
}

/**
 * 获取一个文章
 * @param tid 文章Id
 */
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	topic.Lables = strings.Replace(strings.Replace(
		topic.Lables, "#", " ", -1), "$", "", -1)
	return topic, nil
}
