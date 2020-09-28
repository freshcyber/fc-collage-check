package persist

import (
	public "github.com/freshcyber/fc-public"
)

// GetCollageOrderList 当前拼单列表
func (maria *Mariadb) GetCollageOrderList(time string) ([]public.NpCollageOrders, error) {
	var _obj []public.NpCollageOrders
	err := maria.db.Table("np_collage_orders").Where("status = 1 and end_time < ? and type =1",time).Order("end_time asc").Find(&_obj).Error
	if err != nil {
		return _obj, err
	}
	return _obj, nil
}


// GetCollageOrder 当前同一拼单
func (maria *Mariadb) GetCollageOrder(collage_no string) ([]public.NpCollageOrders, error) {
	var _obj []public.NpCollageOrders
	err := maria.db.Table("np_collage_orders").Where("status > 0 and collage_no = ?",collage_no).Find(&_obj).Error
	if err != nil {
		return _obj, err
	}
	return _obj, nil
}


// CollageInfo 拼单详情
func (maria *Mariadb) CollageInfo(collage_id int64) (public.NpCollageProducts, error) {
	var _obj public.NpCollageProducts
	err := maria.db.Table("np_collage_products").Where("id = ?", collage_id).Find(&_obj).Error
	if err != nil {
		return _obj, err
	}
	return _obj, nil
}

// GetCollageOrderCount 拼单单数
func (maria *Mariadb) GetCollageOrderCount(collage_no string) int64 {
	var count int64
	err := maria.db.Table("np_collage_orders").Where("collage_no = ? and status >0",collage_no).Count(&count).Error
	if err != nil {
		return count
	}
	return count
}

// UpdateCollage UpdateCollage
func (maria *Mariadb) UpdateCollage(collage_no string,status int64) error {
	var _map map[string]interface{}
	_map = make(map[string]interface{})
	_map["status"] = status
	_err := maria.db.Table("np_collage_orders").Where("collage_no = ? and status = 1", collage_no).Update(_map).Error
	if _err != nil {
		return _err
	}
	return nil
}

func (maria *Mariadb) GetUserInfo(uuid int64) (public.NpUserUserBase, error) {
	var _obj public.NpUserUserBase
	err := maria.db.Table("np_user_users").Where("uid = ?", uuid).Find(&_obj).Error
	if err != nil {
		return _obj, err
	}
	return _obj, nil
}