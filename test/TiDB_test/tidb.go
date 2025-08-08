package main

import (
	"context"
	"database/sql/driver" // 实现自定义类型数据库读写接口
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/common/llm/gemini" // Gemini Embedding 客户端
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Gemini Embedding API配置
var (
	GeminiApiKey   = "AIzaSyCKURVV8jEX3CsRu_4pysxmJm3IH4mr8VU"
	GeminiModel    = "gemini-2.0-flash"
	GeminiEmbedder = "gemini-embedding-001"
)

// Vector3072 自定义类型，表示长度为3072的浮点向量
type Vector3072 [3072]float32

// GormDataType 返回给GORM的通用数据类型名称，告诉它这是向量类型
func (v Vector3072) GormDataType() string {
	return "vector"
}

// GormDBDataType 返回数据库具体数据类型，这里声明为 TiDB 的 VECTOR(3072)
func (v Vector3072) GormDBDataType(db *gorm.DB, _ interface{}) string {
	return "VECTOR(3072)"
}

// Value 实现 database/sql/driver.Valuer 接口，用于将 Vector3072 转换为数据库可接受的格式
// 这里将向量转成 "[f1,f2,f3,...]" 形式的字符串，TiDB支持以字符串形式插入向量
func (v Vector3072) Value() (driver.Value, error) {
	str := "["
	for i, f := range v {
		if i > 0 {
			str += ","
		}
		// 格式化浮点数
		str += fmt.Sprintf("%f", f)
	}
	str += "]"
	return str, nil
}

// EmbeddedDocument 定义数据库表结构
type EmbeddedDocument struct {
	ID        int        `gorm:"primaryKey"`        // 主键
	Document  string     `gorm:"type:text"`         // 文本内容，存储文档原文
	Embedding Vector3072 `gorm:"type:VECTOR(3072)"` // 3072维向量字段
}

func main() {
	// TiDB 云端数据库连接串，注意替换为你的真实账号密码和地址
	// format(Totally 3 arguments to be coonfiged): 1.usernmae 2.password 3. ip address
	dsn := "kC45oJcMFwADSBQ.root:TgZ5oJm3IH4mr8VU@tcp(gateway01.eu-central-1.prod.aws.tidbcloud.com:4000)/test?charset=utf8mb4&parseTime=True&loc=Local&tls=true"

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	// 自动迁移，会根据 EmbeddedDocument 结构自动创建或更新表结构
	if err := db.AutoMigrate(&EmbeddedDocument{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// 需要插入的文本
	doc := "test message"
	// 创建 Gemini 客户端
	gemini := gemini.NewGeminiLLM(GeminiApiKey, GeminiModel, GeminiEmbedder)

	// 调用 Gemini API 获取文本向量，返回 []float32
	val, err := gemini.Embeding(context.Background(), doc, "")
	if err != nil {
		log.Fatalf("gemini embedding failed: %v", err)
	}

	// 定义一个 Vector3072 类型变量，将返回的切片复制到定长数组
	var embedding Vector3072
	copy(embedding[:], val)

	// 创建要插入的记录切片
	docs := []EmbeddedDocument{
		{ID: 1, Document: doc, Embedding: embedding},
	}

	// 逐条插入数据库
	for _, d := range docs {
		if err := db.Create(&d).Error; err != nil {
			log.Fatalf("insert failed: %v", err)
		}
	}

	fmt.Println("Insert finished")
}
