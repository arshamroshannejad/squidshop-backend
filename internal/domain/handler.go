package domain

type Handler interface {
	Auth() AuthHandler
	User() UserHandler
	Category() CategoryHandler
	Product() ProductHandler
	ProductRating() ProductRatingHandler
	ProductImage() ProductImageHandler
	ProductComment() ProductCommentHandler
}
