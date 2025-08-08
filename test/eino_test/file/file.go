package main

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/loader/file"
	"github.com/cloudwego/eino/components/document"
)

func main() {
	ctx := context.Background()

	// 初始化加载器
	loader, err := file.NewFileLoader(ctx, &file.FileLoaderConfig{
		UseNameAsID: true,
	})
	if err != nil {
		panic(err)
	}

	// 加载文档
	docs, err := loader.Load(ctx, document.Source{
		URI: "../data/test.txt",
	})
	if err != nil {
		panic(err)
	}

	// 使用文档内容
	for _, doc := range docs {
		println(doc.Content)
		// 访问元数据
		fileName := doc.MetaData[file.MetaKeyFileName]
		extension := doc.MetaData[file.MetaKeyExtension]
		source := doc.MetaData[file.MetaKeySource]
		println(fileName, extension, source)
	}
}
