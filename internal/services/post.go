package services

import (
	"post/internal/enums"
	"post/internal/models"
	"post/pkg/app/util"
)

func GetPosts(query models.QueryPost) (*[]models.Post, error) {
	posts, err := models.FindPosts(query)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPost(postId string) (*models.Post, error) {
	post, err := models.FindPost(postId)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func CreatePost(createPost models.Post) (*models.Post, error) {
	result, err := models.CreatePost(models.Post{
		Title:       createPost.Title,
		Description: createPost.Description,
		Status:      createPost.Status,
		UserId:      createPost.UserId,
		CreatedAt:   util.GetLocalTime("Bangkok"),
		UpdatedAt:   util.GetLocalTime("Bangkok"),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ArchivePost(postId string) (*models.Post, error) {
	post, err := models.UpdatePost(postId, models.Post{
		IsArchive: true,
	})
	if err != nil {
		return nil, err
	}
	return post, nil
}

func UpdatePostStatus(postId string, status enums.PostStatus) (*models.Post, error) {
	post, err := models.UpdatePost(postId, models.Post{
		Status: status,
	})
	if err != nil {
		return nil, err
	}
	return post, nil
}

// func UpdatePost(postId string, createPost models.Post) (*models.Post, error) {
// 	post, err := models.UpdatePost(postId, createPost)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return post, nil
// }

// func DeletePost(postId string) error {
// 	post, err := models.DeletePost(postId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
