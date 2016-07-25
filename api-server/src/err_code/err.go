package err_code

const (
	OK                       = 0   //操作正确
	InvalidSessionId         = -1  //无效session id
	UserExist                = -2  //User已注册
	IrregularUsername        = -3  //注册错误，用户名不符合要求
	IrregularPassword        = -4  //注册错误，密码不符合要求
	ErrPassword              = -5  //登陆错误，认证失败
	NoUser                   = -6  //登陆错误， 用户名不存在
	PartyExist               = -7  //赛事创建失败，赛事名已存在
	PartyInvalidData         = -8  //赛事创建失败，赛事信息不全
	InvalidPartyId           = -9  //无效partyid，赛事不存在
	InvalidGameId            = -10 //无效比赛id，比赛不存在
	NoPermission             = -11 //用户无此权限
	InvalidToken             = -12 //无效token
	InvalidData              = -13 //无效数据
	RefundTimeOut            = -14 //退款超时
	ErrCode                  = -15 //验证码失效或者为空
	InvalidrRegId            = -16 //无效reg_id
	InvalidOrderId           = -17 //无效order_id
	PhotoUnexist             = -18 //头像不存在
	NoUserId                 = -19 //查询个人信息错误，user_id不存在
	UserInfoEncryption       = -20 //查询个人信息错误，用户信息加密
	ErrUserCode              = -21 //选手郑洁号码有误
	TeamExist                = -22 //队伍已存在
	InvalidPlayerId          = -23 //无效player_id
	InvalidTeamId            = -24 //无效team_id
	InvalidTeamName          = -25 //无效team_name
	PartyNameExist           = -26 //赛事名称已存在
	IrregularOrUsernameExist = -27 //用户名不符合要求或者已被注册

	SetSessionFailed = -28 //设置sessionid失败
	DelSessionFailed = -29 //删除sessionid失败

	UserNoLogin = -30

	NoOrder           = -31
	NoPhotoChose      = -32
	WechatReturnError = -33
	InvalidOrderNo    = -34 //无效order_no
	NotValidIP        = -35 //无效order_no
	PartyClosed       = -36
	SendEmail_ERR     = -37
	PriceIsNegative   = -38
	GameNameExists    = -39
	PlayerRegNotExist = -40
	NoData            = -41
	NoDetail          = -42
	InvalidPartyState = -43

	ServerErr = -1000 //服务器内部错误
	OtherErr  = -1001 //其他
	ErrJson   = -1002 //json格式不正确
	NoCmd     = -1003 //没有设置cmd参数
	ErrCmd    = -1004 //无效cmd
)

func Get_err_msg(err_code int) (err_code_msg string) {

	switch err_code {
	case InvalidSessionId:
		err_code_msg = "无效session id"
	case UserExist:
		err_code_msg = "User已注册"
	case IrregularUsername:
		err_code_msg = "注册错误，用户名不符合要求"
	case IrregularPassword:
		err_code_msg = "注册错误，密码不符合要求"
	case ErrPassword:
		err_code_msg = "登陆错误，认证失败"
	case NoUser:
		err_code_msg = "登陆错误， 用户名不存在"
	case PartyExist:
		err_code_msg = "赛事创建失败，赛事名已存在"
	case PartyInvalidData:
		err_code_msg = "赛事创建失败，赛事信息不全"
	case InvalidPartyId:
		err_code_msg = "无效partyid，赛事不存在"
	case InvalidGameId:
		err_code_msg = "无效比赛id，比赛不存在"
	case NoPermission:
		err_code_msg = "用户无此权限"
	case InvalidToken:
		err_code_msg = "无效token"
	case InvalidData:
		err_code_msg = "无效数据"

	case RefundTimeOut:
		err_code_msg = "退款超时"
	case ErrCode:
		err_code_msg = "验证码失效或者为空"
	case InvalidrRegId:
		err_code_msg = "无效reg_id"
	case InvalidOrderId:
		err_code_msg = "无效order_id"
	case PhotoUnexist:
		err_code_msg = "头像不存在"
	case NoUserId:
		err_code_msg = "查询个人信息错误，user_id不存在"
	case UserInfoEncryption:
		err_code_msg = "查询个人信息错误，用户信息加密"
	case ErrUserCode:
		err_code_msg = "选手证件号码有误"
	case TeamExist:
		err_code_msg = "队伍已存在"
	case InvalidPlayerId:
		err_code_msg = "无效player_id"
	case InvalidTeamId:
		err_code_msg = "无效team_id"
	case InvalidTeamName:
		err_code_msg = "无效team_name"

	case PartyNameExist:
		err_code_msg = "赛事名称已存在"
	case IrregularOrUsernameExist:
		err_code_msg = "用户名不符合要求或者已被注册"
	case SetSessionFailed:
		err_code_msg = "设置sessionid失败"
	case DelSessionFailed:
		err_code_msg = "删除sessionid失败"
	case UserNoLogin:
		err_code_msg = "用户未登录"
	case NoOrder:
		err_code_msg = "订单不存在"
	case NoPhotoChose:
		err_code_msg = "未选择上传的图片"
	case WechatReturnError:
		err_code_msg = "微信服务器返回码错误"
	case PartyClosed:
		err_code_msg = "赛事已关闭"
	case SendEmail_ERR:
		err_code_msg = "发送邮件失败"
	case GameNameExists:
		err_code_msg = "比赛名已存在"

	case ServerErr:
		err_code_msg = "服务器内部错误"
	case OtherErr:
		err_code_msg = "其他错误"
	case ErrJson:
		err_code_msg = "json格式不正确"
	case NoCmd:
		err_code_msg = "没有设置cmd参数"
	case ErrCmd:
		err_code_msg = "无效cmd"
	case InvalidOrderNo:
		err_code_msg = "无效订单号"
	case NotValidIP:
		err_code_msg = "无效请求ip"
	case PriceIsNegative:
		err_code_msg = "价格为负数"
	case PlayerRegNotExist:
		err_code_msg = "没有选手的报名信息"
	case NoData:
		err_code_msg = "没有数据"
	case InvalidPartyState:
		err_code_msg = "无效赛事状态"
	case OK:
		err_code_msg = ""
	default:
		err_code_msg = "无效错误码"

	}

	return err_code_msg
}
