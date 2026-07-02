package domain

// User is the shared business representation of an authenticated user.
// It intentionally has no transport, database, or framework annotations.
type User struct {
	ID       string
	Name     string
	Company  string
	Email    string
	Password string
	Token    string
	Role     string
	Status   string
}

// Address is the shared business representation of a customer address.
type Address struct {
	Line1      string
	Line2      string
	City       string
	State      string
	Country    string
	PostalCode string
}

// Customer is the shared business representation of a customer assigned to a user.
type Customer struct {
	ID              string
	FirstName       string
	LastName        string
	Email           string
	Phone           string
	Address         *Address
	ShippingAddress *Address
	BillingAddress  *Address
	UserID          string
	User            *User
	CreatedAt       string
	UpdatedAt       string
}

// Service is the shared business representation of a customer-orderable service.
type Service struct {
	ID              string
	Name            string
	Category        string
	Price           float64
	Type            string
	ScheduleDate    string
	StartDate       string
	AgentName       string
	Description     string
	CreatedAt       string
	UpdatedAt       string
	SKU             string
	Active          bool
	Images          []*CatalogImage
	EffectivePrice  float64
	DiscountPercent float64
	DiscountLabel   string
	Clearance       bool
}

// Product is a service-catalog product exposed through the catalog service boundary.
type Product struct {
	ID              string
	Name            string
	Description     string
	SKU             string
	Price           float64
	Active          bool
	CreatedAt       string
	UpdatedAt       string
	Model           string
	Currency        string
	InventoryCount  int32
	Variants        []*ProductVariant
	Categories      []*ProductCategory
	Images          []*CatalogImage
	EffectivePrice  float64
	DiscountPercent float64
	DiscountLabel   string
	Clearance       bool
	Closeout        bool
}

type ProductVariant struct {
	ID              string
	ProductID       string
	Name            string
	Model           string
	Size            string
	Color           string
	SKU             string
	Price           float64
	Active          bool
	InventoryCount  int32
	CreatedAt       string
	UpdatedAt       string
	EffectivePrice  float64
	DiscountPercent float64
	DiscountLabel   string
	Clearance       bool
	Closeout        bool
}

type CatalogImage struct {
	ID          string
	OwnerType   string
	OwnerID     string
	FileName    string
	ContentType string
	SizeBytes   int64
	SortOrder   int32
	Primary     bool
	CreatedAt   string
}

type CatalogImageUpload struct {
	FileName    string
	ContentType string
	Data        []byte
}

type CatalogDiscount struct {
	ID         string
	Name       string
	Scope      string
	TargetID   string
	PercentOff float64
	Active     bool
	StartsAt   string
	EndsAt     string
	Closeout   bool
	CreatedAt  string
	UpdatedAt  string
}

type ProductCategory struct {
	ID          string
	Name        string
	Slug        string
	ParentID    string
	Path        string
	Level       int32
	Active      bool
	CreatedAt   string
	UpdatedAt   string
	AncestorIDs []string
}

type ProductSearchFilter struct {
	Query        string
	CategoryIDs  []string
	Models       []string
	Sizes        []string
	Colors       []string
	PriceBucket  string
	Availability string
	Sort         string
	ActiveOnly   bool
	PageSize     int
	PageToken    string
}

type ProductPriceBucket struct {
	Value        string
	Label        string
	Min          float64
	Max          float64
	IncludesMax  bool
	UnboundedMax bool
}

var ProductPriceBuckets = []ProductPriceBucket{
	{Value: "price_0_50", Label: "0-50", Min: 0, Max: 50},
	{Value: "price_50_100", Label: "50-100", Min: 50, Max: 100},
	{Value: "price_100_200", Label: "100-200", Min: 100, Max: 200},
	{Value: "price_200_500", Label: "200-500", Min: 200, Max: 500},
	{Value: "price_500_1000", Label: "500-1,000", Min: 500, Max: 1000},
	{Value: "price_1000_2000", Label: "1,000-2,000", Min: 1000, Max: 2000, IncludesMax: true},
	{Value: "price_over_2000", Label: "Over 2,000", Min: 2000, UnboundedMax: true},
}

func ProductPriceBucketFor(price float64) ProductPriceBucket {
	for _, bucket := range ProductPriceBuckets {
		if price < bucket.Min {
			continue
		}
		if bucket.UnboundedMax || price < bucket.Max || (bucket.IncludesMax && price <= bucket.Max) {
			return bucket
		}
	}
	return ProductPriceBuckets[0]
}

func ProductPriceBucketByValue(value string) (ProductPriceBucket, bool) {
	for _, bucket := range ProductPriceBuckets {
		if bucket.Value == value {
			return bucket, true
		}
	}
	return ProductPriceBucket{}, false
}

type ProductSearchFacet struct {
	Field  string
	Label  string
	Values []ProductSearchFacetValue
}

type ProductSearchFacetValue struct {
	Value string
	Label string
	Count int32
}

type ProductSearchResult struct {
	Products      []*Product
	Facets        []ProductSearchFacet
	Total         int
	NextPageToken string
}

// CustomerService links a customer to an ordered service.
type CustomerService struct {
	ID         string
	CustomerID string
	ServiceID  string
	Customer   *Customer
	Service    *Service
	Status     string
	OrderedAt  string
}

// Cart groups service and product items selected by a customer before checkout.
type Cart struct {
	ID         string
	CustomerID string
	Customer   *Customer
	Items      []*CartItem
	Subtotal   float64
	Total      float64
	CreatedAt  string
	UpdatedAt  string
	Version    int64
}

// CartItem stores one selected service or product in a cart.
type CartItem struct {
	ID               string
	CartID           string
	ItemType         string
	ServiceID        string
	ProductID        string
	Service          *Service
	Product          *Product
	Quantity         int32
	UnitPrice        float64
	TotalPrice       float64
	Notes            string
	CreatedAt        string
	UpdatedAt        string
	ProductVariantID string
	ProductVariant   *ProductVariant
	ListPrice        float64
	DiscountPercent  float64
	DiscountLabel    string
	Clearance        bool
	Closeout         bool
}

// Order groups one or more services for a customer.
type Order struct {
	ID               string
	CustomerID       string
	UserID           string
	OrderNumber      string
	Status           string
	ScheduleDate     string
	StartDate        string
	DueDate          string
	Subtotal         float64
	Discount         float64
	Tax              float64
	Shipping         float64
	Total            float64
	DiscountCode     string
	DiscountLabel    string
	ShippingMethodID string
	ShippingLabel    string
	TaxRate          float64
	TaxLabel         string
	ShippingAddress  *Address
	BillingAddress   *Address
	Notes            string
	Customer         *Customer
	Services         []*OrderService
	Products         []*OrderProduct
	PaymentStatus    string
	CreatedAt        string
	UpdatedAt        string
}

type OrderPricingRule struct {
	ID               string
	Name             string
	Kind             string
	Code             string
	Calculation      string
	Value            float64
	MinimumSubtotal  float64
	FreeShippingOver float64
	Country          string
	State            string
	TaxableShipping  bool
	Active           bool
	StartsAt         string
	EndsAt           string
	Priority         int32
	CreatedAt        string
	UpdatedAt        string
}

type OrderPricingSummary struct {
	Subtotal         float64
	Discount         float64
	Shipping         float64
	TaxableAmount    float64
	Tax              float64
	Total            float64
	DiscountCode     string
	DiscountLabel    string
	ShippingMethodID string
	ShippingLabel    string
	TaxRate          float64
	TaxLabel         string
}

type OrderProduct struct {
	ID               string
	OrderID          string
	ProductID        string
	ProductVariantID string
	ProductName      string
	VariantName      string
	SKU              string
	Model            string
	Size             string
	Color            string
	UnitPrice        float64
	Quantity         int32
	TotalPrice       float64
	Product          *Product
	ProductVariant   *ProductVariant
	CreatedAt        string
	UpdatedAt        string
}

type PaymentSession struct {
	ID                string
	OrderID           string
	AmountCurrency    string
	Amount            float64
	Provider          string
	ProviderSessionID string
	Status            string
	CreatedAt         string
	UpdatedAt         string
}

// OrderService stores the service snapshot captured when an order is created.
type OrderService struct {
	ID           string
	OrderID      string
	ServiceID    string
	ServiceName  string
	Category     string
	Type         string
	UnitPrice    float64
	Quantity     int32
	TotalPrice   float64
	AgentName    string
	ScheduleDate string
	StartDate    string
	DueDate      string
	Service      *Service
	CreatedAt    string
	UpdatedAt    string
}

// CreateOrderServiceInput is the domain input for adding a service to an order.
type CreateOrderServiceInput struct {
	ServiceID    string
	Quantity     int32
	AgentName    string
	ScheduleDate string
	StartDate    string
	DueDate      string
}
