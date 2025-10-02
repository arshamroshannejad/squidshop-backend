package domain

type Service interface {
	User() UserService
	OTP() OTPService
	Sms() SmsService
	Category() CategoryService
	Product() ProductService
	ProductRating() ProductRatingService
	ProductImage() ProductImageService
	ProductComment() ProductCommentService
	ProductCommentLike() ProductCommentLikeService
	S3() S3Service
}
