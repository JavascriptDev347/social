package main

import (
	"net/http"

	"github.com/JavascriptDev347/social.git/internal/store"
)

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required,max=1000"`
}

// CreateComment godoc
//
//	@Summary		Creates a comment on a post
//	@Description	Creates a comment on the post with the given ID
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Post ID"
//	@Param			payload	body		CreateCommentPayload	true	"Comment content"
//	@Success		201		{object}	store.Comment
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{id}/comments [post]
func (app *application) createCommentForPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateCommentPayload
	var userID int64 = 1
	post := app.getPostFromCtx(r)

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	comment := &store.Comment{
		UserID:  userID,
		PostID:  post.ID,
		Content: payload.Content,
	}

	ctx := r.Context()
	err := app.store.Comments.Create(ctx, comment)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
	}

}
