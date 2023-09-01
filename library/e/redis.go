package e

type redis struct {
	FirstOrder       string //用户首单
	InviteFirstOrder string //邀请用户首单
	Register         string
	CustomerInfo     string
}

var Redis = redis{
	FirstOrder:       "firstOrder",
	Register:         "register",
	CustomerInfo:     "customerInfo",
	InviteFirstOrder: "inviteFirstOrder",
}
