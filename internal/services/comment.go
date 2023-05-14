package services

import (
	"post/internal/models"
	"post/pkg/app/util"
)

func CreateComment(createComment models.Comment) (*models.Comment, error) {
	result, err := models.CreateComment(models.Comment{
		Text:   createComment.Text,
		PostId: createComment.PostId,
		UserId: createComment.UserId,

		CreatedAt: util.GetLocalTime("Bangkok"),
		UpdatedAt: util.GetLocalTime("Bangkok"),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
