package utils

const (
	VERSION = "chinarun 1.0"
)

//if param is requested
const (
	PARAM_REQUESTED = true
	PARAM_OPTIONAL  = false
)

//party_state
const (
	PARTY_STATE_NORMAL       = 0
	PARTY_STATE_NOPARTY      = -1
	PARTY_STATE_PARTY_CLOSED = -2
)

const (
	TIME_FORMAT_LAYOUT = "2006-01-02 15:04:05"
	DATE_FORMAT_LAYOUT = "2006-01-02"
)

// Http Paramneters
const (
	HP_PARTY_NAME               = "party_name"
	HP_COUNTRY                  = "country"
	HP_PROVINCE                 = "province"
	HP_CITY                     = "city"
	HP_ADDR                     = "addr"
	HP_LOC_LONG                 = "loc_long"
	HP_LOC_LAT                  = "loc_lat"
	HP_LIMITATION               = "limitation"
	HP_LIMITATION_TYPE          = "limitation_type"
	HP_EXTRA_INFO_JSON          = "extra_info_json"
	HP_REG_START_TIME           = "reg_start_time"
	HP_REG_END_TIME             = "reg_end_time"
	HP_START_TIME               = "start_time"
	HP_END_TIME                 = "end_time"
	HP_PARTY_ID                 = "party_id"
	HP_GAME_ID                  = "game_id"
	HP_NAME                     = "name"
	HP_RMB_PRICE                = "rmb_price"
	HP_USD_PRICE                = "usd_price"
	HP_WP_USER_ID               = "wp_user_id"
	HP_GENDER_REQ               = "gender_req"
	HP_MIN_AGE_REQ              = "min_age_req"
	HP_MAX_AGE_REQ              = "max_age_req"
	HP_PAGE_NO                  = "page_no"
	HP_PAGE_SIZE                = "page_size"
	HP_ORDER_BY                 = "order_by"
	HP_INCLUDE_CLOSE            = "include_close"
	HP_PLAYER_ID                = "player_id"
	HP_PLAYER_NAME              = "name"
	HP_CERTIFICATE_TYPE         = "certificate_type"
	HP_CERTIFICATE_NO           = "certificate_no"
	HP_MOBILE                   = "mobile"
	HP_EMAIL                    = "email"
	HP_ADDR1                    = "address1"
	HP_ADDR2                    = "address2"
	HP_ZIPCODE                  = "zip"
	HP_GENDER                   = "gender"
	HP_BIRTH_DATE               = "birth_date"
	HP_BLOOD_TYPE               = "blood_type"
	HP_PLAYER_HEIGHT            = "height"
	HP_PLAYER_WEIGHT            = "weight"
	HP_EMERGENCY_CONTACT_NAME   = "emergency_contact_name"
	HP_EMERGENCY_CONTACT_MOBILE = "emergency_contact_mobile"
	HP_T_SHIRT_SIZE             = "t_shirt_size"
	HP_INDUSTRY                 = "industry"
	HP_JOB_LEVEL                = "job_level"
	HP_INCOME                   = "income"
	HP_FAMILY_STATUS            = "family_status"
	HP_TOKEN                    = "token"
	HP_QUERY_INTERVAL           = "query_interval"
	HP_COUNT                    = "count"
	HP_QUEUE_STATE              = "queue_state"
	HP_PLAYERS                  = "players"
	HP_ORDER_NO                 = "order_no"
	HP_BALANCE                  = "balance"
	HP_NOTICE_TYPE              = "notice_type"
	HP_NOTICE_CONTENT           = "content"
	HP_NOTICE_RECIPIENT         = "recipient"
	HP_PAY_STATUS               = "pay_status"

	HP_NOTICE_SUBJECT = "subject"

	// HP_NOTICE_TEMPLATEID = "templateId"
	// HP_NOTICE_PHONE = "phone"
	HP_NOTICE_VARS = "vars"
)

// 定义数据库字段名称
const (
	//party
	DB_PARTY_FN_NAME           = "name"
	DB_PARTY_FN_START_TIME     = "start_time"
	DB_PARTY_FN_REG_START_TIME = "reg_start_time"
	DB_PARTY_FN_CLOSE_TIME     = "close_time"
	DB_PARTY_FN_CITY           = "city"
	DB_PARTY_FN_PROVINCE       = "province"
	DB_PARTY_FN_COUNTRY        = "country"

	//game
	DB_GAME_FN_NAME     = "name"
	DB_GAME_FN_PARTY_ID = "party_id"

	//player
	DB_PLAYER_CERTIFICATE_TYPE = "certificate_type"
	DB_PLAYER_CERTIFICATE_NO   = "certificate_no"

	//registration
	DB_REG_PLAYER_ID = "player_id"
	DB_REG_GAME_ID   = "game_id"

	//order
	DB_ORDER_ORDER_NO    = "order_no"
	DB_ORDER_SUBMIT_TIME = "submit_time"

	// player_score
	DB_PS_PLAYER_ID = "player_id"
	DB_PS_GAME_ID   = "game_id"
)

//wechat params
const (
	WECHAT_PAY_APP_ID           = "wxb5636d699fc1b88b"
	WECHAT_PAY_MCH_ID           = "1227429802"
	WECHAT_PAY_NOTIFY_URL       = "http://139.196.20.209:8080/api/wechat_pay"
	WECHAT_PAY_API_KEY          = "F5483873309762D60EBE8F6E2F987AD9"
	WECHAT_PAY_APP_SECRET       = "1fea0e0ce8ecada2c9f4e319cdbd8d14"
	WECHAT_PAY_TRADE_TYPE       = "NATIVE"
	WECHAT_PAY_WECHAT_ORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
)

// 定义证件类型
const (
	CERT_TYPE_NONE                   = -1
	CERT_TYPE_ID_CARD                = 0 //身份证
	CERT_TYPE_PASSPORT               = 1 //护照
	CERT_TYPE_MILITARY_ID_CARD       = 2 //军官证
	CERT_TYPE_MTP                    = 3 //台胞证
	CERT_TYPE_EEP_HK_MACAU           = 4 //港澳通行证
	CERT_TYPE_HK_MACAU_RETURN_PERMIT = 5 //港澳回乡证
)

// 定义性别类型
const (
	GENDER_NONE   = 0
	GENDER_MALE   = 1
	GENDER_FEMALE = 2
)

//报名排队状态, 0：需要继续等待,1：排队成功,请根据给定reg_token报名, 2：报名名额已满
const (
	REG_QUEUE_STATE_QUEUING        = 0
	REG_QUEUE_STATE_PROCCESSED     = 1
	REG_QUEUE_STATE_FULL           = 2
	REG_QUEUE_STATE_ERROR          = 3 //处理过程遇到错误
	REG_QUEUE_STATE_ALREADY_SIGNUP = 4 //已经报名
)

const (
	CURRENCY_TYPE_NONE = 0
	CURRENCY_TYPE_RMB  = 1
	CURRENCY_TYPE_USD  = 2
)

const (
	PAY_METHOD_NONE   = 0
	PAY_METHOD_ALIPAY = 1
	PAY_METHOD_WECHAT = 2
	PAY_METHOD_PAYPAL = 3
	PAY_METHOD_OTHER  = 4
)

const (
	DEF_PARTY_PAGE_SIZE = 100
	DEF_PARTY_PAGE_NO   = 0

	DEF_ORDER_QUERY_PAGE_SIZE = 20
	DEF_ORDER_QUERY_PAGE_NO   = 0

	STR_TRUE = "true"
	STR_ONE  = "1"
)

const (
	ORDER_PAY_STATE_NONE = 100 //无效状态， 因为是uint8， 所以不用-1

	ORDER_PAY_STATE_FIRST    = 0
	ORDER_PAY_STATE_WAIT_PAY = 0 //等待支付
	ORDER_PAY_STATE_PAYED    = 1 //已支付
	ORDER_PAY_STATE_CANCELED = 2 //已取消
	ORDER_PAY_STATE_REFUNDED = 3 //已退款
	ORDER_PAY_STATE_LAST     = 3
)

//定义衣服尺寸
const (
	T_SHIRT_SIZE_NONE = 0
	T_SHIRT_SIZE_XXS  = 1
	T_SHIRT_SIZE_XS   = 2
	T_SHIRT_SIZE_S    = 3
	T_SHIRT_SIZE_M    = 4
	T_SHIRT_SIZE_L    = 5
	T_SHIRT_SIZE_XL   = 6
	T_SHIRT_SIZE_XXL  = 7
	T_SHIRT_SIZE_XXXL = 8
)

const (
	USER_ID_TYPE_EMAIL   = 1
	USER_ID_TYPE_MOBILE  = 2
	USER_ID_TYPE_UNKNOWN = 3
)

const (
	JSON_PASSWORD   = "password"
	JSON_USER_NAME  = "user_name"
	JSON_USER_ID    = "user_id"
	JSON_ORDER_ID   = "order_id"
	JSON_ORDER_NO   = "order_no"
	JSON_WP_USER_ID = "wp_user_id"
	JSON_GAME_ID    = "game_id"
	JSON_MESSAGE    = "message"

	RETURN_JSON_RESULT    = "result"
	RETURN_JSON_MSG       = "msg"
	RETURN_JSON_GAME_ID   = "game_id"
	RETURN_JSON_PARTY_ID  = "party_id"
	RETURN_JSON_STATE     = "state"
	RETURN_JSON_RETURNURL = "return_url"
)

const (
	PARSE_IN_LOCATION_LAYOUT = "2006-01-02 15:04:05" //golang时间戳转日期标准模板
)

const (
	DB_ID       = "id"
	MSG_USER_ID = "userid"
)

const (
	SESSION_USER_NAME = "user_name"
	SESSION_USER_ID   = "user_id"

	MIN_PASSWORD_LEN = 6

	UPLOAD_TPL                    = "upload.tpl"
	UPLOAD_FAILER_TPL             = "upload_failed.tpl"
	UPLOAD_SUCCESS_TPL            = "upload_success.tpl"
	UPLOAD_FORM_PHOTO_NAME        = "photo"
	UPLOAD_FORM_HEADPORTRAIT_NAME = "head_portrait"

	STATIC_PATH_MODEL = "photo"
)

const (
	WECHAT_NOTICE_PARAMS_RETURN_CODE  = "return_code"
	WECHAT_NOTICE_PARAMS_RESULT_CODE  = "result_code"
	WECHAT_NOTICE_PARAMS_APPID        = "appid"
	WECHAT_NOTICE_PARAMS_MCH_ID       = "mch_id"
	WECHAT_NOTICE_PARAMS_SIGN         = "sign"
	WECHAT_NOTICE_PARAMS_OUT_TRADE_NO = "out_trade_no"
	WECHAT_NOTICE_PARAMS_OPEN_ID      = "openid"

	WECHAT_NOTICE_RETURN_SUCCESS = "SUCCESS"
	WECHAT_NOTICE_RETURN_FAIL    = "FAIL"

	ALIPAY_NOTICE_SUCCESS = "success"
	ALIPAY_NOTICE_FAIL    = "fail"
)

const (
	SEND_SMS_SUCCESS = "success"
	SEND_SMS_FAIL    = "fail"

	SEND_EMAIL_SUCCESS = "success"
	SEND_EMAIL_FAIL    = "fail"

	//send_cloud短信参数
	SEND_CLOUD_PARAMS_SMSUSER     = "smsUser"
	SEND_CLOUD_PARAMS_TEMPLATEDID = "templateId"
	SEND_CLOUD_PARAMS_PHONE       = "phone"
	SEND_CLOUD_PARAMS_VARS        = "vars"
	SEND_CLOUD_PARAMS_SIGNATURE   = "signature"

	//send_cloud邮件参数
	SEND_CLOUD_PARAMS_SUBJECT   = "subject"
	SEND_CLOUD_PARAMS_TO        = "to"
	SEND_CLOUD_PARAMS_HTML      = "html"
	SEND_CLOUD_PARAMS_X_SMTPAPI = "x_smtpapi"
	SEND_CLOUD_PARAMS_API_USER  = "api_user"
	SEND_CLOUD_PARAMS_API_KEY   = "api_key"
	SEND_CLOUD_PARAMS_FROM      = "from"
	SEND_CLOUD_PARAMS_FROM_NAME = "fromname"
)

// table name
const (
	DB_TABLE_GAME         = "game"
	DB_TABLE_USER         = "user"
	DB_TABLE_PARTY        = "party"
	DB_TABLE_ORDER        = "order"
	DB_TABLE_PLAYER       = "player"
	DB_TABLE_REGISTRATION = "registration"
	DB_TABLE_PLAYER_SCORE = "player_score"
	DB_TABLE_USER_PARTY   = "user_party"
	DB_TABLE_MSG          = "message"
)

//message const
const (
	//message type
	MSG_PARTY_REG    = 1
	MSG_PARTY_START  = 2
	MSG_PARTY_PHOTO  = 3
	MSG_USER_FOLLOW  = 4
	MSG_ORDER_PAID   = 5
	MSG_ORDER_REFUND = 6

	//message state const
	MSG_UNREAD  = 1
	MSG_READ    = 2
	MSG_DELETED = 3
)
