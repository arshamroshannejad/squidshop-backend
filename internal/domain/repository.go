package domain

type Repository interface {
	User() UserRepository
	Category() CategoryRepository
	Product() ProductRepository
	ProductRating() ProductRatingRepository
	ProductImage() ProductImageRepository
	ProductComment() ProductCommentRepository
}
