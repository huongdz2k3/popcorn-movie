// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"PopcornMovie/ent"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type AvailableRoomOutput struct {
	IsAvailableRoom bool `json:"isAvailableRoom"`
}

type ChangePasswordInput struct {
	OldPassword        string `json:"oldPassword"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

type CheckOutOutput struct {
	Bin           string `json:"Bin"`
	AccountNumber string `json:"AccountNumber"`
	AccountName   string `json:"AccountName"`
	Amount        int    `json:"Amount"`
	Description   string `json:"Description"`
	OrderCode     int    `json:"OrderCode"`
	Currency      string `json:"Currency"`
	PaymentLinkID string `json:"PaymentLinkId"`
	Status        string `json:"Status"`
	CheckoutURL   string `json:"CheckoutUrl"`
	QRCode        string `json:"QRCode"`
}

type CreateFoodOrderLineInput struct {
	FoodID   string `json:"foodID"`
	Quantity int    `json:"quantity"`
}

type CreateMovieInput struct {
	Title      string         `json:"title"`
	Genre      string         `json:"genre"`
	Status     MovieStatus    `json:"status"`
	Language   string         `json:"language"`
	Director   string         `json:"director"`
	Cast       string         `json:"cast"`
	Rated      string         `json:"rated"`
	Duration   int            `json:"duration"`
	Trailer    string         `json:"trailer"`
	OpeningDay time.Time      `json:"openingDay"`
	Story      string         `json:"story"`
	File       graphql.Upload `json:"file"`
}

type CreateSessionInput struct {
	ID           string    `json:"id"`
	UserID       string    `json:"UserID"`
	RefreshToken string    `json:"RefreshToken"`
	ExpiresAt    time.Time `json:"ExpiresAt"`
}

type CreateShowTimeInput struct {
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
	MovieID string    `json:"movieId"`
	RoomID  string    `json:"roomId"`
}

type CreateTicketInput struct {
	ID    string  `json:"ID"`
	Price float64 `json:"price"`
}

type CreateTransactionInput struct {
	TicketIDs []*CreateTicketInput        `json:"ticketIDs"`
	Foods     []*CreateFoodOrderLineInput `json:"foods"`
}

type CreateUserInput struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        *Role  `json:"role,omitempty"`
}

type DailyRevenueOutput struct {
	Date  time.Time `json:"date"`
	Total float64   `json:"total"`
}

type GenerateTicketInput struct {
	ShowTimeID string  `json:"showTimeID"`
	Price      float64 `json:"price"`
}

type Jwt struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ListAvailableRoomFilter struct {
	StartAt *time.Time `json:"startAt,omitempty"`
	EndAt   *time.Time `json:"endAt,omitempty"`
	RoomID  *string    `json:"roomID,omitempty"`
}

type ListAvailableRoomInput struct {
	Filter *ListAvailableRoomFilter `json:"filter,omitempty"`
}

type ListAvailableRoomOutput struct {
	Data       []*ent.Room       `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListAvailableSeatFilter struct {
	ShowTimeID *string `json:"showTimeID,omitempty"`
}

type ListAvailableSeatInput struct {
	Filter     *ListAvailableSeatFilter `json:"filter"`
	Pagination *PaginationInput         `json:"pagination"`
}

type ListAvailableSeatOutput struct {
	Data       []*ent.Seat       `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListCommentFilter struct {
	MovieID string `json:"movieId"`
}

type ListCommentInput struct {
	Filter     *ListMovieFilter `json:"filter"`
	Pagination *PaginationInput `json:"pagination"`
}

type ListCommentOutput struct {
	Data       []*ent.Comment    `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListFoodInput struct {
	Pagination *PaginationInput `json:"pagination,omitempty"`
}

type ListFoodOutput struct {
	Data       []*ent.Food       `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListMovieFilter struct {
	Status *MovieStatus `json:"status,omitempty"`
}

type ListMovieInput struct {
	Filter     *ListMovieFilter `json:"filter,omitempty"`
	Pagination *PaginationInput `json:"pagination,omitempty"`
}

type ListMovieOutput struct {
	Data       []*ent.Movie      `json:"data,omitempty"`
	Pagination *PaginationOutput `json:"pagination,omitempty"`
}

type ListRoomFilter struct {
	TheaterID  string     `json:"theaterID"`
	ShowTimeID *string    `json:"showTimeID,omitempty"`
	StartAt    *time.Time `json:"startAt,omitempty"`
	EndAt      *time.Time `json:"endAt,omitempty"`
}

type ListRoomInput struct {
	Filter     *ListRoomFilter  `json:"filter,omitempty"`
	Pagination *PaginationInput `json:"pagination,omitempty"`
}

type ListRoomOutput struct {
	Data       []*ent.Room       `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListSeatFilter struct {
	RoomID *string `json:"roomID,omitempty"`
}

type ListSeatInput struct {
	Filter     *ListSeatFilter  `json:"filter"`
	Pagination *PaginationInput `json:"pagination"`
}

type ListSeatOutput struct {
	Data       []*ent.Seat       `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListShowTimeFilter struct {
	MovieID   *string    `json:"movieId,omitempty"`
	RoomID    *string    `json:"roomId,omitempty"`
	TheaterID *string    `json:"theaterId,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
}

type ListShowTimeInput struct {
	Filter     *ListShowTimeFilter `json:"filter"`
	Pagination *PaginationInput    `json:"pagination"`
}

type ListShowTimeOutput struct {
	Data       []*ent.ShowTime   `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListTheaterFilter struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ListTheatersInput struct {
	Pagination *PaginationInput   `json:"pagination,omitempty"`
	Filter     *ListTheaterFilter `json:"filter,omitempty"`
}

type ListTheatersOutput struct {
	Data       []*ent.Theater    `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListTicketFilter struct {
	ShowTimeID string `json:"showTimeID"`
}

type ListTicketInput struct {
	Filter     *ListTicketFilter `json:"filter"`
	Pagination *PaginationInput  `json:"pagination"`
}

type ListTicketOutput struct {
	Data       []*ent.Ticket     `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type ListTransactionFilter struct {
	UserID string `json:"userID"`
}

type ListTransactionInput struct {
	Pagination *PaginationInput       `json:"pagination"`
	Filter     *ListTransactionFilter `json:"filter"`
}

type ListTransactionOutput struct {
	Data       []*ent.Transaction `json:"data"`
	Pagination *PaginationOutput  `json:"pagination"`
}

type ListUserInput struct {
	Pagination *PaginationInput `json:"pagination"`
}

type ListUserOutput struct {
	Data       []*ent.User       `json:"data"`
	Pagination *PaginationOutput `json:"pagination"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MessageCreateOutput struct {
	Output string `json:"output"`
}

type MonthlyRevenueOutput struct {
	Total float64 `json:"total"`
	Month int     `json:"month"`
}

type Mutation struct {
}

type PaginationInput struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PaginationOutput struct {
	Total int `json:"total"`
}

type Query struct {
}

type RegisterInput struct {
	DisplayName     string `json:"displayName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type RenewAccessTokenInput struct {
	RefreshToken string `json:"refreshToken"`
}

type ResetPasswordInput struct {
	Token              string `json:"token"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

type RevenueInput struct {
	Year int `json:"year"`
}

type UpdateMovieInput struct {
	ID         string          `json:"id"`
	Title      *string         `json:"title,omitempty"`
	Genre      *string         `json:"genre,omitempty"`
	Status     *MovieStatus    `json:"status,omitempty"`
	Language   *string         `json:"language,omitempty"`
	Director   *string         `json:"director,omitempty"`
	Cast       *string         `json:"cast,omitempty"`
	Rated      *string         `json:"rated,omitempty"`
	Duration   *int            `json:"duration,omitempty"`
	Trailer    *string         `json:"trailer,omitempty"`
	OpeningDay *time.Time      `json:"openingDay,omitempty"`
	Story      *string         `json:"story,omitempty"`
	File       *graphql.Upload `json:"file,omitempty"`
}

type UpdateShowTimeInput struct {
	ID      string     `json:"id"`
	StartAt *time.Time `json:"startAt,omitempty"`
	EndAt   *time.Time `json:"endAt,omitempty"`
	MovieID *string    `json:"movieId,omitempty"`
	RoomID  *string    `json:"roomId,omitempty"`
}

type UpdateUserInput struct {
	ID          string  `json:"id"`
	DisplayName *string `json:"displayName,omitempty"`
	IsLocked    *bool   `json:"isLocked,omitempty"`
}

type YearlyRevenueOutput struct {
	Total float64                 `json:"total"`
	Arr   []*MonthlyRevenueOutput `json:"arr"`
}

type MovieStatus string

const (
	MovieStatusUpcoming MovieStatus = "UPCOMING"
	MovieStatusOngoing  MovieStatus = "ONGOING"
	MovieStatusOver     MovieStatus = "OVER"
)

var AllMovieStatus = []MovieStatus{
	MovieStatusUpcoming,
	MovieStatusOngoing,
	MovieStatusOver,
}

func (e MovieStatus) IsValid() bool {
	switch e {
	case MovieStatusUpcoming, MovieStatusOngoing, MovieStatusOver:
		return true
	}
	return false
}

func (e MovieStatus) String() string {
	return string(e)
}

func (e *MovieStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MovieStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MovieStatus", str)
	}
	return nil
}

func (e MovieStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RevenueType string

const (
	RevenueTypeDaily   RevenueType = "DAILY"
	RevenueTypeWeekly  RevenueType = "WEEKLY"
	RevenueTypeMonthly RevenueType = "MONTHLY"
	RevenueTypeYearly  RevenueType = "YEARLY"
)

var AllRevenueType = []RevenueType{
	RevenueTypeDaily,
	RevenueTypeWeekly,
	RevenueTypeMonthly,
	RevenueTypeYearly,
}

func (e RevenueType) IsValid() bool {
	switch e {
	case RevenueTypeDaily, RevenueTypeWeekly, RevenueTypeMonthly, RevenueTypeYearly:
		return true
	}
	return false
}

func (e RevenueType) String() string {
	return string(e)
}

func (e *RevenueType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RevenueType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RevenueType", str)
	}
	return nil
}

func (e RevenueType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleCustomer      Role = "CUSTOMER"
	RoleStaff         Role = "STAFF"
	RoleTicketManager Role = "TICKET_MANAGER"
	RoleAdmin         Role = "ADMIN"
)

var AllRole = []Role{
	RoleCustomer,
	RoleStaff,
	RoleTicketManager,
	RoleAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleCustomer, RoleStaff, RoleTicketManager, RoleAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SeatCategory string

const (
	SeatCategoryStandard SeatCategory = "STANDARD"
	SeatCategoryDouble   SeatCategory = "DOUBLE"
)

var AllSeatCategory = []SeatCategory{
	SeatCategoryStandard,
	SeatCategoryDouble,
}

func (e SeatCategory) IsValid() bool {
	switch e {
	case SeatCategoryStandard, SeatCategoryDouble:
		return true
	}
	return false
}

func (e SeatCategory) String() string {
	return string(e)
}

func (e *SeatCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SeatCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SeatCategory", str)
	}
	return nil
}

func (e SeatCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
