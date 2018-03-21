package model

import "fmt"

type Ads struct {
	AdsId      int    `json:"ads_id" xorm:"not null pk autoincr INT(11)"`
	AdsTitle   string `json:"ads_title" xorm:"not null VARCHAR(255)"`
	Href       string `json:"href" xorm:"VARCHAR(255)"`
	ImgPath    string `json:"img_path" xorm:"VARCHAR(1024)"`
	ImgBaseUrl string `json:"img_base_url" xorm:"VARCHAR(1024)"`
	CreatedAt  int    `json:"created_at" xorm:"INT(11)"`
	UpdatedAt  int    `json:"updated_at" xorm:"INT(11)"`
}

func (t Ads) TableName() string {
	return "ty_ads"
}
func (t *Ads) AfterLoad() {
	t.Href = "http"

}

func (t Ads) BeforeUpdate() {

}

func GetAds() *Ads {

	ads := new(Ads)

	has, err := engine.Where("ads_id = ?", "2").Get(ads)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if has == false {
		return nil
	}
	return ads
}

func GetAdsForPage(page int, pageSize int) []*Ads {
	ads := make([]*Ads, 0, pageSize)

	err := engine.Limit(pageSize, (page-1)*pageSize).Find(&ads)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return ads
}
