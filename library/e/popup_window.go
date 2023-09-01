package e

type popupWindow struct {
	StateEnabled       int    //状态正常
	StateDisabled      int    //状态不可用
	MemberGroupAllUser string //用户组所有用户
	MemberGroupNon     string //用户组非会员
	MemberGroupMember  string //用户组会员
	RateCategoryDay    string //频率类型每日
	RateCategoryEvery  string //频率类型每日
}

var PopupWindow = popupWindow{
	StateEnabled:       1,
	StateDisabled:      2,
	MemberGroupAllUser: "all_user",
	MemberGroupNon:     "non_member",
	MemberGroupMember:  "member",
	RateCategoryDay:    "day",
	RateCategoryEvery:  "every",
}
