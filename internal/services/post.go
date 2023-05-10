package services

import "post/internal/models"

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

func ArchivePost(postId string) (*models.Post, error) {
	post, err := models.UpdatePost(postId, models.Post{
		IsArchive: true,
	})
	if err != nil {
		return nil, err
	}
	return post, nil
}

// func CreatePost(createPost models.Post) (*models.Post, error) {
// 	post, err := models.CreatePost(createPost)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return post, nil
// }

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
