package service

import (
	"errors"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/dao"
	"san616qi/app/model"
)

var Article = new(articleService)

type articleService struct{}

//新增文章
func (as *articleService) AddArticle(r *model.ArticleAddApiReq) error {

	var article = &model.Article{}

	//数据转化处理
	article.UserId = r.UserId
	article.Content = r.Content
	article.ArticleTitle = r.ArticleTitle
	article.CreateAt = gtime.Now()
	article.UpdateAt = gtime.Now()
	article.ForumId = r.ForumId

	//处理tags和urls
	len1 := len(r.Tags)
	for i, v := range r.Tags {
		if i == len1-1 {
			article.Tags = article.Tags + v
		} else {
			article.Tags = article.Tags + v + ","
		}
	}
	len2 := len(r.Urls)
	for i, v := range r.Urls {
		if i == len2-1 {
			article.Urls = article.Urls + v
		} else {
			article.Urls = article.Urls + v + ","
		}
	}

	if _, err := dao.Article.Save(article); err != nil {
		return errors.New("新增文章错误")
	} else {
		return nil
	}

}

//更新文章
func (as *articleService) UpdateArticle(r *model.ArticleUpdateApiReq) error {

	//准备更新数据
	var article = &model.Article{}

	article.UpdateAt = gtime.Now()
	article.ArticleId = r.ArticleId
	article.UserId = r.UserId
	article.Content = r.Content
	article.ArticleTitle = r.ArticleTitle
	article.UpdateAt = gtime.Now()
	article.ForumId = r.ForumId

	//处理tags和urls
	len1 := len(r.Tags)
	for i, v := range r.Tags {
		if i == len1-1 {
			article.Tags = article.Tags + v
		} else {
			article.Tags = article.Tags + v + ","
		}
	}
	len2 := len(r.Urls)
	for i, v := range r.Urls {
		if i == len2-1 {
			article.Urls = article.Urls + v
		} else {
			article.Urls = article.Urls + v + ","
		}
	}

	//更新文章
	if _, err := dao.Article.Where("article_id=?", article.ArticleId).Update(article); err != nil {
		return errors.New("更新文章失败")
	} else {
		return nil
	}

}

//删除文章，只能删除自己的，会做一次校验
func (as *articleService) DelArticle(r *model.ArticleDelApiReq) error {

	//获取校验的参数
	userid := r.UserId
	articleid := r.ArticleId

	uid, err := dao.Article.DB.GetAll("select user_id from article where article_id=?",articleid);
	if err != nil {
		return errors.New("数据库查询错误")
	} else {

		if gconv.Int64(uid) == userid {
			return nil
		} else {
			return errors.New("仅能操作属于自己的文章")
		}
	}

}

//获取文章详情
func (as *articleService) GetArticle(articleid int) (error, *model.ArticleProfileEntity) {

	//var entity *model.ArticleProfileEntity
	return nil,nil
}