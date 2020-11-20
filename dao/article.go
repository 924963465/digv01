package dao

import (
	"fmt"
	"github.com/liuhongdi/digv01/global"
	"github.com/liuhongdi/digv01/model"
)

func SelectOneArticle(articleId int64) (*model.Article, error) {
	var articleOne *model.Article
	err := global.DBLink.Where("articleId=?",articleId).First(&articleOne).Error
	if (err != nil) {
		return nil,err
	} else {
		return articleOne,nil
	}
}

func SelectcountAll() (int, error) {
	var count int
	err := global.DBLink.Where("isPublish=?",1).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func SelectAllArticle(pageOffset int,pageSize int) ([]*model.Article, error) {
	fields := []string{"articleId", "subject", "url"}
	rows,err := global.DBLink.Select(fields).Table(model.Article{}.TableName()).Where("isPublish=?",1).Offset(pageOffset).Limit(pageSize).Rows()

	if err != nil {
		fmt.Println("sql is error:")
		fmt.Println(err)
		return nil, err
	}

	//fmt.Println(rows.)
	defer rows.Close()
	var articles []*model.Article
	for rows.Next() {
		fmt.Println("rows.next:")
		r := &model.Article{}
		if err := rows.Scan(&r.ArticleId, &r.Subject, &r.Url); err != nil {
			fmt.Println("rows.next:")
			fmt.Println(err)
			return nil, err
		}
		articles = append(articles, r)
	}

	return articles, nil
}