package parseParamsSrv

import (
	"github.com/chinarun/utils"
)

func GetEditPersonalInfoMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_USER_NAME:   {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_USER_AVATAR: {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},

		HP_ADDRESS:       {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_GENDER:        {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_CONSTELLATION: {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_PROFESSION:    {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_ABOUT:         {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
	}
}

func GetPartyCreateHttpParametersMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_PARTY_NAME:      {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_LIMITATION:      {PARAM_OPTIONAL, utils.DATA_TYPE_UINT32},
		HP_COUNTRY:         {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_PROVINCE:        {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_CITY:            {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_ADDR:            {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_LIMITATION_TYPE: {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_LOC_LAT:         {PARAM_OPTIONAL, utils.DATA_TYPE_FLOAT32},
		HP_LOC_LONG:        {PARAM_OPTIONAL, utils.DATA_TYPE_FLOAT32},
		HP_REG_START_TIME:  {PARAM_REQUESTED, utils.DATA_TYPE_DATETIME},
		HP_REG_END_TIME:    {PARAM_REQUESTED, utils.DATA_TYPE_DATETIME},
		HP_START_TIME:      {PARAM_REQUESTED, utils.DATA_TYPE_DATETIME},
		HP_END_TIME:        {PARAM_REQUESTED, utils.DATA_TYPE_DATETIME},

		HP_SLOGAN:        {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_WEBSITE:       {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_TYPE:          {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_INTRODUCTION:  {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_SCHEDULE:      {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_PRICE:         {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_SCORE:         {PARAM_OPTIONAL, utils.DATA_TYPE_FLOAT32},
		HP_SIGNUP_MALE:   {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_SIGNUP_FEMALE: {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
	}
}

func GetPartyEditHttpParametersMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_PARTY_ID: {PARAM_REQUESTED, utils.DATA_TYPE_STRING},

		HP_PARTY_NAME:      {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_LIMITATION:      {PARAM_OPTIONAL, utils.DATA_TYPE_UINT32},
		HP_COUNTRY:         {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_PROVINCE:        {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_CITY:            {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_ADDR:            {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_LIMITATION_TYPE: {PARAM_OPTIONAL, utils.DATA_TYPE_UINT8},
		HP_LOC_LAT:         {PARAM_OPTIONAL, utils.DATA_TYPE_FLOAT32},
		HP_LOC_LONG:        {PARAM_OPTIONAL, utils.DATA_TYPE_FLOAT32},
		HP_REG_START_TIME:  {PARAM_OPTIONAL, utils.DATA_TYPE_DATETIME},
		HP_REG_END_TIME:    {PARAM_OPTIONAL, utils.DATA_TYPE_DATETIME},
		HP_START_TIME:      {PARAM_OPTIONAL, utils.DATA_TYPE_DATETIME},
		HP_END_TIME:        {PARAM_OPTIONAL, utils.DATA_TYPE_DATETIME},

		HP_SLOGAN:        {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_WEBSITE:       {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_TYPE:          {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_INTRODUCTION:  {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_SCHEDULE:      {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_PRICE:         {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_SCORE:         {PARAM_OPTIONAL, utils.DATA_TYPE_FLOAT32},
		HP_SIGNUP_MALE:   {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_SIGNUP_FEMALE: {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
	}
}

func GetPartyListParamsMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{
		HP_PAGE_SIZE:     {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_PAGE_NO:       {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_MONTH:         {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_VALID_STATE:   {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_COUNTRY:       {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_PROVINCE:      {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_CITY:          {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_TAGS:          {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_ORDER_BY:      {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_INCLUDE_CLOSE: {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
		HP_USER_ID:       {PARAM_OPTIONAL, utils.DATA_TYPE_STRING},
		HP_TYPE:          {PARAM_OPTIONAL, utils.DATA_TYPE_INT},
	}

}

func GetPartyStateUpdateParametersMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_PARTY_ID:    {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_PARTY_STATE: {PARAM_REQUESTED, utils.DATA_TYPE_UINT8},
	}
}