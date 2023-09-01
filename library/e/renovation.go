package e

type pageName struct {
	HomePage         string // 首页
	UserCenter       string // 会员中心页
	GoodsPage        string // 菜品页
	ConfirmOrderPage string // 菜品页
}

var PageName = pageName{
	HomePage:         "HomePage",
	UserCenter:       "UserCenter",
	GoodsPage:        "GoodsPage",
	ConfirmOrderPage: "ConfirmOrderPage",
}

var PageNameRange = []string{
	PageName.HomePage,
	PageName.UserCenter,
	PageName.GoodsPage,
	PageName.ConfirmOrderPage,
}
