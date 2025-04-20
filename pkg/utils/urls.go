package utils

// URLConfig represents a set of Domino's API endpoints for a specific country
type URLConfig struct {
	SourceURI string
	Location  struct {
		Find string
	}
	Store struct {
		Find string
		Info string
		Menu string
	}
	Order struct {
		Validate string
		Price    string
		Place    string
	}
	Images     string
	TrackRoot  string
	Track      string
	Token      string
	Upsell     string
	StepUpsell string
}

// USA Domino's Pizza API URLs
var USA = URLConfig{
	SourceURI: "order.dominos.com",
	Location: struct {
		Find string
	}{
		Find: "https://api.dominos.com/store-locator-international-service/findAddress?latitude=${lat}&longitude=${lon}",
	},
	Store: struct {
		Find string
		Info string
		Menu string
	}{
		Find: "https://order.dominos.com/power/store-locator?s=${line1}&c=${line2}&type=${pickUpType}",
		Info: "https://order.dominos.com/power/store/${storeID}/profile",
		Menu: "https://order.dominos.com/power/store/${storeID}/menu?lang=${lang}&structured=true",
	},
	Order: struct {
		Validate string
		Price    string
		Place    string
	}{
		Validate: "https://order.dominos.com/power/validate-order",
		Price:    "https://order.dominos.com/power/price-order",
		Place:    "https://order.dominos.com/power/place-order",
	},
	Images:     "https://cache.dominos.com/olo/6_47_2/assets/build/market/US/_en/images/img/products/larges/${productCode}.jpg",
	TrackRoot:  "https://tracker.dominos.com/tracker-presentation-service/",
	Track:      "v2/orders",
	Token:      "https://order.dominos.com/power/paymentGatewayService/braintree/token",
	Upsell:     "https://api.dominos.com/upsell-service/stores/upsellForOrder/",
	StepUpsell: "https://api.dominos.com/upsell-service/stores/stepUpsellForOrder",
}

// Canada Domino's Pizza API URLs
var Canada = URLConfig{
	SourceURI: "order.dominos.ca",
	Location: struct {
		Find string
	}{
		Find: "https://api.dominos.com/store-locator-international-service/findAddress?latitude=${lat}&longitude=${lon}",
	},
	Store: struct {
		Find string
		Info string
		Menu string
	}{
		Find: "https://order.dominos.ca/power/store-locator?s=${line1}&c=${line2}&type=${type}",
		Info: "https://order.dominos.ca/power/store/${storeID}/profile",
		Menu: "https://order.dominos.ca/power/store/${storeID}/menu?lang=${lang}&structured=true",
	},
	Order: struct {
		Validate string
		Price    string
		Place    string
	}{
		Validate: "https://order.dominos.ca/power/validate-order",
		Price:    "https://order.dominos.ca/power/price-order",
		Place:    "https://order.dominos.ca/power/place-order",
	},
	Images:     "https://cache.dominos.com/nolo/ca/en/6_44_3/assets/build/market/CA/_en/images/img/products/larges/${itemCode}.jpg",
	Track:      "https://order.dominos.ca/orderstorage/GetTrackerData?",
	Token:      "https://order.dominos.com/power/paymentGatewayService/braintree/token",
	Upsell:     "https://api.dominos.com/upsell-service/stores/upsellForOrder/",
	StepUpsell: "https://api.dominos.com/upsell-service/stores/stepUpsellForOrder",
}

// Default URLs configuration (starts with USA)
var URLs = USA

// UseInternational switches the URLs to use international endpoints
func UseInternational(config URLConfig) {
	URLs = config
}
