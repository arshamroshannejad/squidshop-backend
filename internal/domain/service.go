package domain

type Service interface {
	User() UserService
	OTP() OTPService
	Sms() SmsService
	Category() CategoryService
	Product() ProductService
	ProductRating() ProductRatingService
}
