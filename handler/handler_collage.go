package handler

import (
	"strings"
	"encoding/json"
	"github.com/freshcyber/fc-collage-check/utils"
	"github.com/freshcyber/fc-collage-check/config"
	"errors"
	"fmt"
	"time"
	"github.com/freshcyber/fc-collage-check/persist"
)

// CollageOrderInfo CollageOrderInfo
type CollageOrderInfo struct {
}

// Run Run
func (g CollageOrderInfo) Run() {
	fmt.Println("=== 进入定时 ===")
	err := GenCheckCollageOrder()
	if err != nil {
		fmt.Println("=== error : ===", err.Error())
	}
}

// GenCheckCollageOrder GenCheckCollageOrder
func GenCheckCollageOrder() error {
	t := time.Now()
	times := t.Format("2006-01-02 15:04:05")
	// 查询当前拼单
	_objs, err := persist.GMariadb.GetCollageOrderList(times)
	fmt.Println(err)
	if err != nil {
		return err
	}
	if len(_objs) == 0 {
		return errors.New("没有正在进行中的拼单")
	}
	for _,v := range _objs {
		// 查询拼单商品信息
		coll_pro,_ := persist.GMariadb.CollageInfo(v.CollageId)
		// 查询拼单单数
		coll_order_count := persist.GMariadb.GetCollageOrderCount(v.CollageNo)
		fmt.Println(coll_order_count)
		fmt.Println(coll_pro.CollageNumber)

		if coll_order_count == coll_pro.CollageNumber {
			// 拼单完成
			err := persist.GMariadb.UpdateCollage(v.CollageNo,2)
			if err != nil {
				return err
			}
		}else{
			// 拼单退款
			err := UpdateCollageOrder(v.CollageNo)
			if err != nil {
				fmt.Println(err)
				return err
			}
			// 更改拼单状态
			err = persist.GMariadb.UpdateCollage(v.CollageNo,3)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Return struct {
	Flag  		int64 	 `form:"flag" json:"flag"`    				// 状态
	Msg			string	 `form:"msg" json:"msg"`       			// 错误返回
}

func UpdateCollageOrder(collage_no string) error {
	// 查询拼团订单
	coll_order ,_ := persist.GMariadb.GetCollageOrder(collage_no)
	if len(coll_order) >0 {
		for _,v :=range coll_order {
			// 获取用户openid
			user,_ := persist.GMariadb.GetUserInfo(v.Uuid)
			// 退款
			urlStr := config.Server.Urlgo + "/weapp/v1/pay/order_refund"
			_props := make(map[string]interface{})
			_props["openid"] = user.WeChat
			_props["order_no"] = v.OrderNo
			_sign := utils.APICalcSign(_props, config.Server.APIAppendKey, config.Server.APIMd5Key)
			_props["api_token"] = strings.ToLower(_sign)
			_json, _ := json.Marshal(_props)
			_ret, err := PostJSONString(urlStr, string(_json))
			if err != nil {
				return err
			}
			str := []byte(_ret)
			stu := Return{}
			err = json.Unmarshal(str, &stu)
			if stu.Flag == 2 {
				return errors.New(stu.Msg)
			}
		}
	}
	return nil
}