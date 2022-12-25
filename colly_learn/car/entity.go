package main

type Car struct {
	ID int `json:"ID"`
	// Site          string      `json:"site"`
	BrandPro string `json:"brand_pro"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Style    string `json:"style"`
	Source   int    `json:"source"`
	// Level         string      `json:"level"`
	// ListedTime    int         `json:"listed_time"`
	// Engine        string      `json:"engine"`
	// CarLength     string      `json:"car_length"`
	// CarWidth      string      `json:"car_width"`
	// CarHeight     string      `json:"car_height"`
	// Wheelbase     string      `json:"wheelbase"`
	// FrontBraking  string      `json:"front_braking"`
	// RearBraking   string      `json:"rear_braking"`
	// TiresCount    int         `json:"tires_count"`
	// FrontTires    string      `json:"front_tires"`
	// RearTires     string      `json:"rear_tires"`
	// AddTime       string      `json:"add_time"`
	// UpdateTime    string      `json:"update_time"`
	// CarID         int         `json:"car_id"`
	// SerID         int         `json:"ser_id"`
	// BoxLength     interface{} `json:"Box_length"`
	// Drive         interface{} `json:"drive"`
	// TiresNorm     interface{} `json:"tires_norm"`
	// BulletinModel interface{} `json:"bulletin_model"`
	// Source        int         `json:"source"`
	// JiebaSortMd5  string      `json:"jieba_sort_md5"`
	// OldMd5        interface{} `json:"old_md5"`
	// LocalSource   int         `json:"local_source"`
}

type CarGroup struct {
	BrandPro      string `json:"brand_pro"`
	Year          int    `json:"year"`
	AID           int    `json:"a_id"`
	BID           int    `json:"b_id"`
	EqualStr      string `json:"equal_str"`
	EqualCount    int    `json:"equal_count"`
	NotEqualStr   string `json:"not_equal_str"`
	NotEqualCount int    `json:"not_equal_count"`
	Path          int    `json:"path"`
	LabelID       int    `json:"label_id"`
}

type CarGroup2 struct {
	BrandPro        string `json:"brand_pro"`
	Year            int    `json:"year"`
	AID             int    `json:"a_id"`
	BID             int    `json:"b_id"`
	BrandEqualStr   string `json:"brand_equal_str"`
	ModelEqualStr   string `json:"model_equal_str"`
	StyleEqualStr   string `json:"style_equal_str"`
	BrandEqualCount int    `json:"brand_equal_count"`
	ModelEqualCount int    `json:"model_equal_count"`
	StyleEqualCount int    `json:"style_equal_count"`
	AbBrandNes      string `json:"ab_brand_nes"`
	BaBrandNes      string `json:"ba_brand_nes"`
	AbModelNes      string `json:"ab_model_nes"`
	BaModelNes      string `json:"ba_model_nes"`
	AbStyleNes      string `json:"ab_style_nes"`
	BaStyleNes      string `json:"ba_style_nes"`
	AbBrandNec      int    `json:"ab_brand_nec"`
	BaBrandNec      int    `json:"ba_brand_nec"`
	AbModelNec      int    `json:"ab_model_nec"`
	BaModelNec      int    `json:"ba_model_nec"`
	AbStyleNec      int    `json:"ab_style_nec"`
	BaStyleNec      int    `json:"ba_style_nec"`
	AbBrandFlag     int    `json:"ab_brand_flag"`
	BaBrandFlag     int    `json:"ba_brand_flag"`
	AbModelFlag     int    `json:"ab_model_flag"`
	BaModelFlag     int    `json:"ba_model_flag"`
	AbStyleFlag     int    `json:"ab_style_flag"`
	BaStyleFlag     int    `json:"ba_style_flag"`
	Status          int    `json:"status"`
}

type ImgInfo struct {
	ImgURL  string `json:"img_url"`
	FdfsURL string `json:"fdfs_url"`
	Source  int    `json:"source"`
}