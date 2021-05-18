package gobudins

import "time"

const (
	RouteAccessToken   = "/auth/token/access"
	RouteUsers         = "/users"
	RouteAccounts      = "/accounts"
	RouteAuthTokenCode = "/auth/token/code"
)

type ErrorResponse struct {
	Code        string `json:"error"`
	Description string `json:"error_description"`
}

type ConnectCallbackData struct {
	Code         string `json:"code"`
	ConnectionID string `json:"connectionID"`
}

type AskForToken struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

const UserMe = "me"

type User struct {
	ID       int       `json:"id"`
	Signin   time.Time `json:"signin"`
	Platform int       `json:"platform"`
}

// Account as described at https://docs.budget-insight.com/reference/bank-accounts#response-bankaccount-object
type Account struct {
	ID           int              `json:"id"`
	ConnectionID *int             `json:"id_connection"`
	UserID       *int             `json:"id_user"`
	SourceID     *int             `json:"id_source"`
	ParentID     *int             `json:"id_parent"`
	Number       *string          `json:"number"`
	OriginalName string           `json:"original_name"`
	Balance      *float64         `json:"balance"`
	Coming       *float64         `json:"comming"`
	Display      bool             `json:"display"`
	LastUpdate   *time.Time       `json:"last_update"`
	Deleted      *time.Time       `json:"deleted"`
	Disabled     *time.Time       `json:"disabled"`
	IBAN         *string          `json:"iban"`
	Currency     *Currency        `json:"currency"`
	Type         AccountType      `json:"type"`
	TypeID       int              `json:"id_type"`
	Bookmarked   int              `json:"bookmarked"`
	Name         string           `json:"name"`
	Error        *string          `json:"error"`
	Usage        BankAccountUsage `json:"usage"`
	Ownsership   string           `json:"ownership"`
	CompanyName  *string          `json:"company_name"`
	Loan         *Loan            `json:"loan"`
}

// BankAccountUsage as described at https://docs.budget-insight.com/reference/bank-accounts#bankaccountusage-values
type BankAccountUsage string

const (
	BankAccountUsagePriv BankAccountUsage = "PRIV"
	BankAccountUsageOrga BankAccountUsage = "ORGA"
	BankAccountUsageAsso BankAccountUsage = "ASSO"
)

// Loan as described at https://docs.budget-insight.com/reference/bank-accounts#loan-object
type Loan struct {
	TotalAmount       *float64   `json:"total_amount"`
	AvailableAmount   *float64   `json:"available_amount"`
	UsedAcmount       *float64   `json:"used_amount"`
	SubscriptionDate  *time.Time `json:"subscription_date"`
	MaturityDate      *time.Time `json:"maturity_date"`
	NextPaymentAmount *float64   `json:"next_payment_amount"`
	NextPatmentAmount *time.Time `json:"next_payment_date"`
	Rate              *float64   `json:"rate"`
	NbPaymentsLeft    *int       `json:"nb_payments_left"`
	NbPaymentsDone    *int       `json:"nb_payments_done"`
	NbPaymentsTotal   *int       `json:"nb_payments_total"`
	LastPaymentAmount *float64   `json:"last_payment_amount"`
	LastPaymentDate   *time.Time `json:"last_payment_date"`
	AccountLabel      *string    `json:"account_label"`
	InsuranceLabel    *string    `json:"insurance_label"`
	Duration          *int       `json:"duration"`
}

type Currency struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	Precision int    `json:"precision"`
}

type AccountTypeName string

const (
	AccountTypeNameCheckings      AccountTypeName = "checking"
	AccountTypeNameSavings        AccountTypeName = "savings"
	AccountTypeNameDeposit        AccountTypeName = "deposit"
	AccountTypeNameLoan           AccountTypeName = "loan"
	AccountTypeNameMarket         AccountTypeName = "market"
	AccountTypeNameJoint          AccountTypeName = "joint"
	AccountTypeNameCard           AccountTypeName = "card"
	AccountTypeNameLifeInsurance  AccountTypeName = "lifeinsurance"
	AccountTypeNamePEE            AccountTypeName = "pee"
	AccountTypeNamePERCO          AccountTypeName = "perco"
	AccountTypeNameArticle83      AccountTypeName = "article83"
	AccountTypeNameRSP            AccountTypeName = "rsp"
	AccountTypeNamePEA            AccountTypeName = "pea"
	AccountTypeNameCapitalisation AccountTypeName = "capitalisation"
	AccountTypeNamePERP           AccountTypeName = "perp"
	AccountTypeNameMadelin        AccountTypeName = "madelin"
	AccountTypeNameUnknow         AccountTypeName = "unknown"
)

// AccountType as described at https://docs.budget-insight.com/reference/bank-account-types#response-accounttype-object
type AccountType struct {
	ID           int             `json:"id"`
	Name         AccountTypeName `json:"name"`
	ParentID     *int            `json:"id_parent"`
	IsInvest     bool            `json:"is_invest"`
	DisplayName  string          `json:"display_name"`
	DisplayNameP string          `json:"display_name_p"`
}

type FinanceSecurityType string

const (
	FinanceSecurityTypeOPCVM   FinanceSecurityType = "OPCVM"
	FinanceSecurityTypeETF     FinanceSecurityType = "Trackers - ETF"
	FinanceSecurityTypeActions FinanceSecurityType = "Actions"
)

type Investment struct {
	ID                int                  `json:"id"`
	AccountID         int                  `json:"id_account"`
	SecurityID        int                  `json:"id_security"`
	TypeID            *FinanceSecurityType `json:"id_type"`
	Label             string               `json:"label"`
	Code              *string              `json:"code"`
	CodeType          string               `json:"code_type"`
	Source            string               `json:"source"`
	Description       *string              `json:"description"`
	Quantity          float64              `json:"quantity"`
	UnitPrice         float64              `json:"unitprice"`
	UnitValue         float64              `json:"unitvalue"`
	Valuation         float64              `json:"valuation"`
	Diff              float64              `json:"diff"`
	DiffPercent       float64              `json:"diff_percent"`
	PrevDiff          *float64             `json:"prev_diff"`
	PrevDiffPercent   *float64             `json:"prev_diff_percent"`
	VDate             time.Time            `json:"vdate"`
	PrevVDate         *time.Time           `json:"prev_vdate"`
	PortfolioShare    float64              `json:"portfolio_share"`
	Calculated        float64              `json:"calculated"`
	Deleted           *time.Time           `json:"deleted"`
	LastUpdate        *time.Time           `json:"last_update"`
	OriginalCurrency  *Currency            `json:"original_currency"`
	OriginalValuation *float64             `json:"original_valuation"`
	OriginalUnitvalue *float64             `json:"original_unitvalue"`
	OriginalUnitprice *float64             `json:"original_unitprice"`
	OriginalDiff      int                  `json:"original_diff"`
	// Details           int             `json:"details"`
}
