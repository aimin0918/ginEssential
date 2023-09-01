package e

type subscriber struct {
	Privacy         string
	CachePath       string
	CustomerPrivacy string
	CustomerBenefit string

	Region string
}

var Subscriber = subscriber{
	Privacy:         "privacys",
	CachePath:       `tmp/json_caches/privacys/`,
	CustomerPrivacy: "CustomerPrivacy",
	CustomerBenefit: "CustomerBenefit",

	Region: "region",
}
