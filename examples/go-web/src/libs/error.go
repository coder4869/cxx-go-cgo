package libs

const (
	RESULT_SUCCESS			= 0
	RESULT_PARAM_ERROR		= 1001
	RESULT_JSON_MA_ERROR	= 1002
)

type ErrInfo struct {
	ErrCode	int32  `json:"err_code"`
	ErrMsg	string `json:"err_msg"`
	Help	string `json:"help_info"`
}
type ErrSlice []ErrInfo

var ErrInfos = map[int32]ErrInfo {
	0: {
		ErrCode: 0,
		ErrMsg: "操作成功",
		Help:   "操作成功",
	},
	1001: {
		ErrCode: 1001,
		ErrMsg: "参数接收错误",
		Help:   "参数接收错误",
	},
	1002: {
		ErrCode: 1002,
		ErrMsg: "Json 解析失败",
		Help:   "Json 解析失败",
	},
} 

func (Info ErrSlice) Len() int {
	return len(Info)
}

func (Info ErrSlice) Swap(i, j int) {
	Info[i], Info[j] = Info[j], Info[i]
}

func (Info ErrSlice) Less(i, j int) bool {
	return Info[i].ErrCode < Info[j].ErrCode
}
