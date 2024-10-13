package tribute

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions,omitempty"`
	NextFrom     string        `json:"nextFrom,omitempty"`
}

type Discount struct {
	Enabled bool    `json:"enabled,omitempty"`
	Percent float64 `json:"percent,omitempty"`
}

type Subscription struct {
	ID                   int64    `json:"id,omitempty"`
	ChannelID            int64    `json:"channelId,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	Period               string   `json:"period,omitempty"`
	Currency             string   `json:"currency,omitempty"`
	Amount               float64  `json:"amount,omitempty"`
	InviteLink           string   `json:"inviteLink,omitempty"`
	AppInviteLink        string   `json:"appInviteLink,omitempty"`
	IsDonate             bool     `json:"isDonate,omitempty"`
	IsDeleted            bool     `json:"isDeleted,omitempty"`
	AcceptCards          bool     `json:"acceptCards,omitempty"`
	AcceptWalletPay      bool     `json:"acceptWalletPay,omitempty"`
	UserID               int64    `json:"userID,omitempty"`
	CommentAccessEnabled bool     `json:"commentAccessEnabled,omitempty"`
	FirstPeriodDiscount  Discount `json:"firstPeriodDiscount,omitempty"`
	CancellationDiscount Discount `json:"cancellationDiscount,omitempty"`
}

type SubscriptionMember struct {
	ID                          int64   `json:"id,omitempty"`
	UserID                      int64   `json:"userId,omitempty"`
	SubscriptionID              int64   `json:"subscriptionId,omitempty"`
	Price                       float64 `json:"price,omitempty"`
	Period                      string  `json:"period,omitempty"`
	Currency                    string  `json:"currency,omitempty"`
	ActivatedAt                 int64   `json:"activatedAt,omitempty"`
	ExpiresAt                   int64   `json:"expiresAt,omitempty"`
	Status                      string  `json:"status,omitempty"`
	PaymentMethod               string  `json:"paymentMethod,omitempty"`
	IsDonate                    bool    `json:"isDonate,omitempty"`
	IsDeleted                   bool    `json:"isDeleted,omitempty"`
	IsExpired                   bool    `json:"isExpired,omitempty"`
	Migrated                    bool    `json:"migrated,omitempty"`
	FirstPeriodDiscountPercent  float64 `json:"firstPeriodDiscountPercent,omitempty"`
	CancellationDiscountPercent float64 `json:"cancellationDiscountPercent,omitempty"`
}
type Channel struct {
	ID        int64  `json:"id,omitempty"`
	ChatID    int64  `json:"chatId,omitempty"`
	Name      string `json:"name,omitempty"`
	UserName  string `json:"userName,omitempty"`
	PhotoURL  string `json:"photoUrl,omitempty"`
	IsDeleted bool   `json:"isDeleted,omitempty"`
	IsGroup   bool   `json:"isGroup,omitempty"`
}
type User struct {
	ID         int64  `json:"id,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	UserName   string `json:"userName,omitempty"`
	Sex        int64  `json:"sex,omitempty"`
	PhotoURL   string `json:"photoUrl,omitempty"`
	TelegramID int64  `json:"telegramId,omitempty"`
}
type Image struct {
	ID        int64  `json:"id,omitempty"`
	Path      string `json:"path,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
}
type DonationRequest struct {
	ID                   int64   `json:"id,omitempty"`
	ChannelID            int64   `json:"channelId,omitempty"`
	Amount               float64 `json:"amount,omitempty"`
	MinAmount            float64 `json:"minAmount,omitempty"`
	Currency             string  `json:"currency,omitempty"`
	Title                string  `json:"title,omitempty"`
	Description          string  `json:"description,omitempty"`
	ButtonText           string  `json:"buttonText,omitempty"`
	ImageID              int64   `json:"imageId,omitempty"`
	Image                Image   `json:"image,omitempty"`
	CustomCover          bool    `json:"customCover,omitempty"`
	Link                 string  `json:"link,omitempty"`
	AcceptCards          bool    `json:"acceptCards,omitempty"`
	AcceptWalletPay      bool    `json:"acceptWalletPay,omitempty"`
	AcceptStars          bool    `json:"acceptStars,omitempty"`
	IsDeleted            bool    `json:"isDeleted,omitempty"`
	IsNewDonationRequest bool    `json:"isNewDonationRequest,omitempty"`
}

type Card struct {
	ID    int64  `json:"id,omitempty"`
	Last4 string `json:"last4,omitempty"`
	Brand string `json:"brand,omitempty"`
}

type WithdrawalAccount struct {
	ID             int64       `json:"id,omitempty"`
	CreatedAt      int64       `json:"createdAt,omitempty"`
	Country        string      `json:"country,omitempty"`
	CountryIso2    string      `json:"countryIso2,omitempty"`
	MethodType     string      `json:"methodType,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	ObjectID       int64       `json:"objectId,omitempty"`
	UserID         int64       `json:"userId,omitempty"`
	IsFiatToCrypto bool        `json:"isFiatToCrypto,omitempty"`
	Params         interface{} `json:"params,omitempty"`
	Card           Card        `json:"card,omitempty"`
}

type PayoutInfo struct {
	PayAmount    float64 `json:"payAmount,omitempty"`
	PayCurrency  string  `json:"payCurrency,omitempty"`
	ExchangeRate float64 `json:"exchangeRate,omitempty"`
}

type Donation struct {
	ID                int64   `json:"id,omitempty"`
	UserID            int64   `json:"userId,omitempty"`
	DonationRequestID int64   `json:"donationRequestId,omitempty"`
	ChannelID         int64   `json:"channelId,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	Period            string  `json:"period,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	Active            bool    `json:"active,omitempty"`
	CreatedAt         int64   `json:"createdAt,omitempty"`
	Message           string  `json:"message,omitempty"`
	Anonymously       *bool   `json:"anonymously,omitempty"`
}

type Transaction struct {
	ID                 int64              `json:"id,omitempty"`
	OtherUserID        int64              `json:"otherUserId,omitempty"`
	ChannelID          int64              `json:"channelId,omitempty"`
	DonationRequestID  int64              `json:"donationRequestId,omitempty"`
	Type               string             `json:"type,omitempty"`
	ObjectID           int64              `json:"objectId,omitempty"`
	Amount             float64            `json:"amount,omitempty"`
	Currency           string             `json:"currency,omitempty"`
	CreatedAt          int64              `json:"createdAt,omitempty"`
	ServiceFee         float64            `json:"serviceFee,omitempty"`
	PaymentMethod      string             `json:"paymentMethod,omitempty"`
	Total              float64            `json:"total,omitempty"`
	Subscription       Subscription       `json:"subscription,omitempty"`
	SubscriptionMember SubscriptionMember `json:"subscriptionMember,omitempty"`
	Channel            Channel            `json:"channel,omitempty"`
	User               User               `json:"user,omitempty"`
	Donation           Donation           `json:"donation,omitempty"`
	DonationRequest    DonationRequest    `json:"donationRequest,omitempty"`
	PaymentCurrency    string             `json:"paymentCurrency,omitempty"`
	PaymentAmount      float64            `json:"paymentAmount,omitempty"`
	WithdrawalAccount  WithdrawalAccount  `json:"withdrawalAccount,omitempty"`
	Card               Card               `json:"card,omitempty"`
	PayoutInfo         PayoutInfo         `json:"payoutInfo,omitempty"`
}
