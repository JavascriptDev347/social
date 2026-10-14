package main

import (
	"net/http"

	"github.com/JavascriptDev347/social.git/internal/store"
)

// GetUserFeed godoc
//
//	@Summary		Fetches the user feed
//	@Description	Fetches posts from the user and the users they follow, with pagination, search and filtering
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int		false	"Max number of posts to return"	default(20)
//	@Param			offset	query		int		false	"Number of posts to skip"		default(0)
//	@Param			sort	query		string	false	"Sort order by creation date"	Enums(asc, desc)	default(desc)
//	@Param			tags	query		string	false	"Comma separated list of tags to filter by (max 5)"
//	@Param			search	query		string	false	"Search term matched against post title and content"
//	@Param			since	query		string	false	"Only posts created after this date (2006-01-02 15:04:05)"
//	@Param			until	query		string	false	"Only posts created before this date (2006-01-02 15:04:05)"
//	@Success		200		{array}		store.PostWithMetaData
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	ctx := r.Context()
	feed, err := app.store.Post.GetUserFeed(ctx, int64(1), fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
