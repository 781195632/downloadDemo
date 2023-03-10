package downloadDemo

import (
	"fmt"
	"strconv"
)

type SelectItem struct {
	calculate string
	as        string
}

type ColumnDefine struct {
	label string
	value string
}

type ReportPlatformItem struct {
	Date                  string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`                                                                  // 日期
	Month                 string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`                                                                // 月
	ProjectId             string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`                                       // 应用id
	ProjectName           string `protobuf:"bytes,4,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`                                 // 应用
	PlatformId            string `protobuf:"bytes,5,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`                                    // 渠道id
	PlatformName          string `protobuf:"bytes,6,opt,name=platform_name,json=platformName,proto3" json:"platform_name,omitempty"`                              // 渠道
	AdvertiserId          string `protobuf:"bytes,7,opt,name=advertiser_id,json=advertiserId,proto3" json:"advertiser_id,omitempty"`                              // 账户id
	AdvertiserName        string `protobuf:"bytes,8,opt,name=advertiser_name,json=advertiserName,proto3" json:"advertiser_name,omitempty"`                        // 账户名称
	AdvertiserNameDisplay string `protobuf:"bytes,9,opt,name=advertiser_name_display,json=advertiserNameDisplay,proto3" json:"advertiser_name_display,omitempty"` // 账户名称显示
	AdgroupId             string `protobuf:"bytes,10,opt,name=adgroup_id,json=adgroupId,proto3" json:"adgroup_id,omitempty"`                                      // 广告组id
	AdgroupName           string `protobuf:"bytes,11,opt,name=adgroup_name,json=adgroupName,proto3" json:"adgroup_name,omitempty"`                                // 广告组名称
	AdgroupNameDisplay    string `protobuf:"bytes,12,opt,name=adgroup_name_display,json=adgroupNameDisplay,proto3" json:"adgroup_name_display,omitempty"`         // 广告组名称显示
	AdId                  string `protobuf:"bytes,13,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`                                                     // 广告id
	AdName                string `protobuf:"bytes,14,opt,name=ad_name,json=adName,proto3" json:"ad_name,omitempty"`                                               // 广告名称
	AdNameDisplay         string `protobuf:"bytes,15,opt,name=ad_name_display,json=adNameDisplay,proto3" json:"ad_name_display,omitempty"`                        // 广告名称显示
	StudioId              string `protobuf:"bytes,16,opt,name=studio_id,json=studioId,proto3" json:"studio_id,omitempty"`                                         // 营销工作室ID
	StudioName            string `protobuf:"bytes,17,opt,name=studio_name,json=studioName,proto3" json:"studio_name,omitempty"`                                   // 营销工作室
	OptimizerId           string `protobuf:"bytes,18,opt,name=optimizer_id,json=optimizerId,proto3" json:"optimizer_id,omitempty"`                                // 优化师ID
	OptimizerName         string `protobuf:"bytes,19,opt,name=optimizer_name,json=optimizerName,proto3" json:"optimizer_name,omitempty"`                          // 优化师
	DeviceOs              uint32 `protobuf:"varint,20,opt,name=device_os,json=deviceOs,proto3" json:"device_os,omitempty"`                                        // 设备系统id
	DeviceOsName          string `protobuf:"bytes,21,opt,name=device_os_name,json=deviceOsName,proto3" json:"device_os_name,omitempty"`                           // 设备系统
	Cost                  string `protobuf:"bytes,22,opt,name=cost,proto3" json:"cost,omitempty"`                                                                 // 消耗
	Show                  int64  `protobuf:"varint,23,opt,name=show,proto3" json:"show,omitempty"`                                                                // 展示数
	Click                 int64  `protobuf:"varint,24,opt,name=click,proto3" json:"click,omitempty"`                                                              // 点击数
	Convert               int64  `protobuf:"varint,25,opt,name=convert,proto3" json:"convert,omitempty"`                                                          // 转化数
	Active                int64  `protobuf:"varint,26,opt,name=active,proto3" json:"active,omitempty"`                                                            // 激活数
	ConvertCost           string `protobuf:"bytes,27,opt,name=convert_cost,json=convertCost,proto3" json:"convert_cost,omitempty"`                                // 转化成本
	ActiveCost            string `protobuf:"bytes,28,opt,name=active_cost,json=activeCost,proto3" json:"active_cost,omitempty"`                                   // 激活成本
	CostByThousandShow    string `protobuf:"bytes,29,opt,name=cost_by_thousand_show,json=costByThousandShow,proto3" json:"cost_by_thousand_show,omitempty"`       // 千次展示成本
	Ctr                   string `protobuf:"bytes,30,opt,name=ctr,proto3" json:"ctr,omitempty"`                                                                   // CTR
	Cvr                   string `protobuf:"bytes,31,opt,name=cvr,proto3" json:"cvr,omitempty"`                                                                   // CVR
	NewUser               int64  `protobuf:"varint,32,opt,name=new_user,json=newUser,proto3" json:"new_user,omitempty"`                                           // 新增用户
	Pay1                  int64  `protobuf:"varint,33,opt,name=pay1,proto3" json:"pay1,omitempty"`                                                                // 1日付费
	Pay2                  int64  `protobuf:"varint,34,opt,name=pay2,proto3" json:"pay2,omitempty"`                                                                // 2日付费
	Pay3                  int64  `protobuf:"varint,35,opt,name=pay3,proto3" json:"pay3,omitempty"`                                                                // 3日付费
	Pay4                  int64  `protobuf:"varint,36,opt,name=pay4,proto3" json:"pay4,omitempty"`                                                                // 4日付费
	Pay5                  int64  `protobuf:"varint,37,opt,name=pay5,proto3" json:"pay5,omitempty"`                                                                // 5日付费
	Pay6                  int64  `protobuf:"varint,38,opt,name=pay6,proto3" json:"pay6,omitempty"`                                                                // 6日付费
	Pay7                  int64  `protobuf:"varint,39,opt,name=pay7,proto3" json:"pay7,omitempty"`                                                                // 7日付费
	Pay8                  int64  `protobuf:"varint,40,opt,name=pay8,proto3" json:"pay8,omitempty"`                                                                // 8日付费
	Pay9                  int64  `protobuf:"varint,41,opt,name=pay9,proto3" json:"pay9,omitempty"`                                                                // 9日付费
	Pay10                 int64  `protobuf:"varint,42,opt,name=pay10,proto3" json:"pay10,omitempty"`                                                              // 10日付费
	Pay11                 int64  `protobuf:"varint,43,opt,name=pay11,proto3" json:"pay11,omitempty"`                                                              // 11日付费
	Pay12                 int64  `protobuf:"varint,44,opt,name=pay12,proto3" json:"pay12,omitempty"`                                                              // 12日付费
	Pay13                 int64  `protobuf:"varint,45,opt,name=pay13,proto3" json:"pay13,omitempty"`                                                              // 13日付费
	Pay14                 int64  `protobuf:"varint,46,opt,name=pay14,proto3" json:"pay14,omitempty"`                                                              // 14日付费
	Pay15                 int64  `protobuf:"varint,47,opt,name=pay15,proto3" json:"pay15,omitempty"`                                                              // 15日付费
	Pay30                 int64  `protobuf:"varint,48,opt,name=pay30,proto3" json:"pay30,omitempty"`                                                              // 30日付费
	Pay45                 int64  `protobuf:"varint,49,opt,name=pay45,proto3" json:"pay45,omitempty"`                                                              // 45日付费
	Pay60                 int64  `protobuf:"varint,50,opt,name=pay60,proto3" json:"pay60,omitempty"`                                                              // 60日付费
	Pay90                 int64  `protobuf:"varint,51,opt,name=pay90,proto3" json:"pay90,omitempty"`                                                              // 90日付费
	Pay120                int64  `protobuf:"varint,52,opt,name=pay120,proto3" json:"pay120,omitempty"`                                                            // 120日付费
	Pay150                int64  `protobuf:"varint,53,opt,name=pay150,proto3" json:"pay150,omitempty"`                                                            // 150日付费
	Pay180                int64  `protobuf:"varint,54,opt,name=pay180,proto3" json:"pay180,omitempty"`                                                            // 180日付费
	Pay210                int64  `protobuf:"varint,55,opt,name=pay210,proto3" json:"pay210,omitempty"`                                                            // 210日付费
	Pay240                int64  `protobuf:"varint,56,opt,name=pay240,proto3" json:"pay240,omitempty"`                                                            // 240日付费
	Pay270                int64  `protobuf:"varint,57,opt,name=pay270,proto3" json:"pay270,omitempty"`                                                            // 270日付费
	Pay300                int64  `protobuf:"varint,58,opt,name=pay300,proto3" json:"pay300,omitempty"`                                                            // 300日付费
	Pay360                int64  `protobuf:"varint,59,opt,name=pay360,proto3" json:"pay360,omitempty"`                                                            // 360日付费
	TotalPay              int64  `protobuf:"varint,60,opt,name=total_pay,json=totalPay,proto3" json:"total_pay,omitempty"`                                        // 累计付费
	Ltv1                  string `protobuf:"bytes,61,opt,name=ltv1,proto3" json:"ltv1,omitempty"`                                                                 // 首日LTV
	Ltv2                  string `protobuf:"bytes,62,opt,name=ltv2,proto3" json:"ltv2,omitempty"`                                                                 // 次日LTV
	Ltv3                  string `protobuf:"bytes,63,opt,name=ltv3,proto3" json:"ltv3,omitempty"`                                                                 // 3日LTV
	Ltv4                  string `protobuf:"bytes,64,opt,name=ltv4,proto3" json:"ltv4,omitempty"`                                                                 // 4日LTV
	Ltv5                  string `protobuf:"bytes,65,opt,name=ltv5,proto3" json:"ltv5,omitempty"`                                                                 // 5日LTV
	Ltv6                  string `protobuf:"bytes,66,opt,name=ltv6,proto3" json:"ltv6,omitempty"`                                                                 // 6日LTV
	Ltv7                  string `protobuf:"bytes,67,opt,name=ltv7,proto3" json:"ltv7,omitempty"`                                                                 // 7日LTV
	Ltv8                  string `protobuf:"bytes,68,opt,name=ltv8,proto3" json:"ltv8,omitempty"`                                                                 // 8日LTV
	Ltv9                  string `protobuf:"bytes,69,opt,name=ltv9,proto3" json:"ltv9,omitempty"`                                                                 // 9日LTV
	Ltv10                 string `protobuf:"bytes,70,opt,name=ltv10,proto3" json:"ltv10,omitempty"`                                                               // 10日LTV
	Ltv11                 string `protobuf:"bytes,71,opt,name=ltv11,proto3" json:"ltv11,omitempty"`                                                               // 11日LTV
	Ltv12                 string `protobuf:"bytes,72,opt,name=ltv12,proto3" json:"ltv12,omitempty"`                                                               // 12日LTV
	Ltv13                 string `protobuf:"bytes,73,opt,name=ltv13,proto3" json:"ltv13,omitempty"`                                                               // 13日LTV
	Ltv14                 string `protobuf:"bytes,74,opt,name=ltv14,proto3" json:"ltv14,omitempty"`                                                               // 14日LTV
	Ltv15                 string `protobuf:"bytes,75,opt,name=ltv15,proto3" json:"ltv15,omitempty"`                                                               // 15日LTV
	Ltv30                 string `protobuf:"bytes,76,opt,name=ltv30,proto3" json:"ltv30,omitempty"`                                                               // 30日LTV
	Ltv45                 string `protobuf:"bytes,77,opt,name=ltv45,proto3" json:"ltv45,omitempty"`                                                               // 45日LTV
	Ltv60                 string `protobuf:"bytes,78,opt,name=ltv60,proto3" json:"ltv60,omitempty"`                                                               // 60日LTV
	Ltv90                 string `protobuf:"bytes,79,opt,name=ltv90,proto3" json:"ltv90,omitempty"`                                                               // 90日LTV
	Ltv120                string `protobuf:"bytes,80,opt,name=ltv120,proto3" json:"ltv120,omitempty"`                                                             // 120日LTV
	Ltv150                string `protobuf:"bytes,81,opt,name=ltv150,proto3" json:"ltv150,omitempty"`                                                             // 150日LTV
	Ltv180                string `protobuf:"bytes,82,opt,name=ltv180,proto3" json:"ltv180,omitempty"`                                                             // 180日LTV
	Ltv210                string `protobuf:"bytes,83,opt,name=ltv210,proto3" json:"ltv210,omitempty"`                                                             // 210日LTV
	Ltv240                string `protobuf:"bytes,84,opt,name=ltv240,proto3" json:"ltv240,omitempty"`                                                             // 240日LTV
	Ltv270                string `protobuf:"bytes,85,opt,name=ltv270,proto3" json:"ltv270,omitempty"`                                                             // 270日LTV
	Ltv300                string `protobuf:"bytes,86,opt,name=ltv300,proto3" json:"ltv300,omitempty"`                                                             // 300日LTV
	Ltv360                string `protobuf:"bytes,87,opt,name=ltv360,proto3" json:"ltv360,omitempty"`                                                             // 360日LTV
	TotalLtv              string `protobuf:"bytes,88,opt,name=total_ltv,json=totalLtv,proto3" json:"total_ltv,omitempty"`                                         // 累计LTV
	Roi1                  string `protobuf:"bytes,89,opt,name=roi1,proto3" json:"roi1,omitempty"`                                                                 // 首日ROI
	Roi2                  string `protobuf:"bytes,90,opt,name=roi2,proto3" json:"roi2,omitempty"`                                                                 // 次日ROI
	Roi3                  string `protobuf:"bytes,91,opt,name=roi3,proto3" json:"roi3,omitempty"`                                                                 // 3日ROI
	Roi4                  string `protobuf:"bytes,92,opt,name=roi4,proto3" json:"roi4,omitempty"`                                                                 // 4日ROI
	Roi5                  string `protobuf:"bytes,93,opt,name=roi5,proto3" json:"roi5,omitempty"`                                                                 // 5日ROI
	Roi6                  string `protobuf:"bytes,94,opt,name=roi6,proto3" json:"roi6,omitempty"`                                                                 // 6日ROI
	Roi7                  string `protobuf:"bytes,95,opt,name=roi7,proto3" json:"roi7,omitempty"`                                                                 // 7日ROI
	Roi8                  string `protobuf:"bytes,96,opt,name=roi8,proto3" json:"roi8,omitempty"`                                                                 // 8日ROI
	Roi9                  string `protobuf:"bytes,97,opt,name=roi9,proto3" json:"roi9,omitempty"`                                                                 // 9日ROI
	Roi10                 string `protobuf:"bytes,98,opt,name=roi10,proto3" json:"roi10,omitempty"`                                                               // 10日ROI
	Roi11                 string `protobuf:"bytes,99,opt,name=roi11,proto3" json:"roi11,omitempty"`                                                               // 11日ROI
	Roi12                 string `protobuf:"bytes,100,opt,name=roi12,proto3" json:"roi12,omitempty"`                                                              // 12日ROI
	Roi13                 string `protobuf:"bytes,101,opt,name=roi13,proto3" json:"roi13,omitempty"`                                                              // 13日ROI
	Roi14                 string `protobuf:"bytes,102,opt,name=roi14,proto3" json:"roi14,omitempty"`                                                              // 14日ROI
	Roi15                 string `protobuf:"bytes,103,opt,name=roi15,proto3" json:"roi15,omitempty"`                                                              // 15日ROI
	Roi30                 string `protobuf:"bytes,104,opt,name=roi30,proto3" json:"roi30,omitempty"`                                                              // 30日ROI
	Roi45                 string `protobuf:"bytes,105,opt,name=roi45,proto3" json:"roi45,omitempty"`                                                              // 45日ROI
	Roi60                 string `protobuf:"bytes,106,opt,name=roi60,proto3" json:"roi60,omitempty"`                                                              // 60日ROI
	Roi90                 string `protobuf:"bytes,107,opt,name=roi90,proto3" json:"roi90,omitempty"`                                                              // 90日ROI
	Roi120                string `protobuf:"bytes,108,opt,name=roi120,proto3" json:"roi120,omitempty"`                                                            // 120日ROI
	Roi150                string `protobuf:"bytes,109,opt,name=roi150,proto3" json:"roi150,omitempty"`                                                            // 150日ROI
	Roi180                string `protobuf:"bytes,110,opt,name=roi180,proto3" json:"roi180,omitempty"`                                                            // 180日ROI
	Roi210                string `protobuf:"bytes,111,opt,name=roi210,proto3" json:"roi210,omitempty"`                                                            // 210日ROI
	Roi240                string `protobuf:"bytes,112,opt,name=roi240,proto3" json:"roi240,omitempty"`                                                            // 240日ROI
	Roi270                string `protobuf:"bytes,113,opt,name=roi270,proto3" json:"roi270,omitempty"`                                                            // 270日ROI
	Roi300                string `protobuf:"bytes,114,opt,name=roi300,proto3" json:"roi300,omitempty"`                                                            // 300日ROI
	Roi360                string `protobuf:"bytes,115,opt,name=roi360,proto3" json:"roi360,omitempty"`                                                            // 360日ROI
	TotalRoi              string `protobuf:"bytes,116,opt,name=total_roi,json=totalRoi,proto3" json:"total_roi,omitempty"`                                        // 累计ROI
	NewPaidUser           int64  `protobuf:"varint,117,opt,name=new_paid_user,json=newPaidUser,proto3" json:"new_paid_user,omitempty"`                            // 首日付费数
	NewPaidRate           string `protobuf:"bytes,118,opt,name=new_paid_rate,json=newPaidRate,proto3" json:"new_paid_rate,omitempty"`                             // 新增付费率
	Stay2                 string `protobuf:"bytes,119,opt,name=stay2,proto3" json:"stay2,omitempty"`                                                              // 次日留存
	Stay3                 string `protobuf:"bytes,120,opt,name=stay3,proto3" json:"stay3,omitempty"`                                                              // 3日留存
	Stay4                 string `protobuf:"bytes,121,opt,name=stay4,proto3" json:"stay4,omitempty"`                                                              // 4日留存
	Stay5                 string `protobuf:"bytes,122,opt,name=stay5,proto3" json:"stay5,omitempty"`                                                              // 5日留存
	Stay6                 string `protobuf:"bytes,123,opt,name=stay6,proto3" json:"stay6,omitempty"`                                                              // 6日留存
	Stay7                 string `protobuf:"bytes,124,opt,name=stay7,proto3" json:"stay7,omitempty"`                                                              // 7日留存
	Stay8                 string `protobuf:"bytes,125,opt,name=stay8,proto3" json:"stay8,omitempty"`                                                              // 8日留存
	Stay9                 string `protobuf:"bytes,126,opt,name=stay9,proto3" json:"stay9,omitempty"`                                                              // 9日留存
	Stay10                string `protobuf:"bytes,127,opt,name=stay10,proto3" json:"stay10,omitempty"`                                                            // 10日留存
	Stay11                string `protobuf:"bytes,128,opt,name=stay11,proto3" json:"stay11,omitempty"`                                                            // 11日留存
	Stay12                string `protobuf:"bytes,129,opt,name=stay12,proto3" json:"stay12,omitempty"`                                                            // 12日留存
	Stay13                string `protobuf:"bytes,130,opt,name=stay13,proto3" json:"stay13,omitempty"`                                                            // 13日留存
	Stay14                string `protobuf:"bytes,131,opt,name=stay14,proto3" json:"stay14,omitempty"`                                                            // 14日留存
	Stay15                string `protobuf:"bytes,132,opt,name=stay15,proto3" json:"stay15,omitempty"`                                                            // 15日留存
	Stay30                string `protobuf:"bytes,133,opt,name=stay30,proto3" json:"stay30,omitempty"`                                                            // 30日留存
	Stay45                string `protobuf:"bytes,134,opt,name=stay45,proto3" json:"stay45,omitempty"`                                                            // 45日留存
	Stay60                string `protobuf:"bytes,135,opt,name=stay60,proto3" json:"stay60,omitempty"`                                                            // 60日留存
	Stay90                string `protobuf:"bytes,136,opt,name=stay90,proto3" json:"stay90,omitempty"`                                                            // 90日留存
	Stay120               string `protobuf:"bytes,137,opt,name=stay120,proto3" json:"stay120,omitempty"`                                                          // 120日留存
	Stay150               string `protobuf:"bytes,138,opt,name=stay150,proto3" json:"stay150,omitempty"`                                                          // 150日留存
	Stay180               string `protobuf:"bytes,139,opt,name=stay180,proto3" json:"stay180,omitempty"`                                                          // 180日留存
	Times3                string `protobuf:"bytes,140,opt,name=times3,proto3" json:"times3,omitempty"`                                                            // 次三比
	Times4                string `protobuf:"bytes,141,opt,name=times4,proto3" json:"times4,omitempty"`                                                            // 次四比
	Times5                string `protobuf:"bytes,142,opt,name=times5,proto3" json:"times5,omitempty"`                                                            // 次五比
	Times6                string `protobuf:"bytes,143,opt,name=times6,proto3" json:"times6,omitempty"`                                                            // 次六比
	Times7                string `protobuf:"bytes,144,opt,name=times7,proto3" json:"times7,omitempty"`                                                            // 次七比
	Pay2StayRate          string `protobuf:"bytes,145,opt,name=pay2_stay_rate,json=pay2StayRate,proto3" json:"pay2_stay_rate,omitempty"`                          // 付费次留率
	Pay7StayRate          string `protobuf:"bytes,146,opt,name=pay7_stay_rate,json=pay7StayRate,proto3" json:"pay7_stay_rate,omitempty"`                          // 付费七留率
	ActiveArpu            string `protobuf:"bytes,147,opt,name=active_arpu,json=activeArpu,proto3" json:"active_arpu,omitempty"`                                  // 活跃arpu
	PaidArpu              string `protobuf:"bytes,148,opt,name=paid_arpu,json=paidArpu,proto3" json:"paid_arpu,omitempty"`                                        // 付费arpu
	NewUserCost           string `protobuf:"bytes,149,opt,name=new_user_cost,json=newUserCost,proto3" json:"new_user_cost,omitempty"`                             // 新增成本
	NewPaidUserCost       string `protobuf:"bytes,150,opt,name=new_paid_user_cost,json=newPaidUserCost,proto3" json:"new_paid_user_cost,omitempty"`               // 新增付费成本
}

type ResponseReportPlatform struct {
	DataList []*ReportPlatformItem `protobuf:"bytes,1,rep,name=data_list,json=dataList,proto3" json:"data_list,omitempty"`
}

var excelColumnList []*ColumnDefine
var columnDefineMap map[string]*ColumnDefine
var result = &ResponseReportPlatform{}

var columnDefineList = []*ColumnDefine{
	{label: "千次展示成本", value: "cost_by_thousand_show"},
	{label: "新增付费成本", value: "new_paid_user_cost"},
	{label: "新增成本", value: "new_user_cost"},
	{label: "展示", value: "show"},
	{label: "点击", value: "click"},
	{label: "消耗", value: "cost"},
	{label: "转化", value: "convert"},
	{label: "激活", value: "active"},
	{label: "转化成本", value: "convert_cost"},
	{label: "激活成本", value: "active_cost"},
	{label: "CTR", value: "ctr"},
	{label: "CVR", value: "cvr"},
	{label: "日期", value: "date"},
	{label: "月", value: "month"},
	{label: "应用ID", value: "project_id"},
	{label: "应用", value: "project_name"},
	{label: "工作室ID", value: "studio_id"},
	{label: "工作室", value: "studio_name"},
	{label: "优化师ID", value: "optimizer_id"},
	{label: "优化师", value: "optimizer_name"},
	{label: "渠道", value: "platform_id"},
	{label: "设备系统", value: "device_os"},
	{label: "账户ID", value: "advertiser_id"},
	{label: "账户名称md5", value: "advertiser_name"},
	{label: "账户名称", value: "advertiser_name_display"},
	{label: "计划ID", value: "adgroup_id"},
	{label: "计划名称md5", value: "adgroup_name"},
	{label: "计划名称", value: "adgroup_name_display"},
	{label: "广告id", value: "ad_id"},
	{label: "广告名称md5", value: "ad_name"},
	{label: "广告名称", value: "ad_name_display"},
	{label: "新增设备数", value: "new_user"},
	{label: "1日付费", value: "pay1"},
	{label: "2日付费", value: "pay2"},
	{label: "3日付费", value: "pay3"},
	{label: "4日付费", value: "pay4"},
	{label: "5日付费", value: "pay5"},
	{label: "6日付费", value: "pay6"},
	{label: "7日付费", value: "pay7"},
	{label: "8日付费", value: "pay8"},
	{label: "9日付费", value: "pay9"},
	{label: "10日付费", value: "pay10"},
	{label: "11日付费", value: "pay11"},
	{label: "12日付费", value: "pay12"},
	{label: "13日付费", value: "pay13"},
	{label: "14日付费", value: "pay14"},
	{label: "15日付费", value: "pay15"},
	{label: "16日付费", value: "pay16"},
	{label: "17日付费", value: "pay17"},
	{label: "18日付费", value: "pay18"},
	{label: "19日付费", value: "pay19"},
	{label: "20日付费", value: "pay20"},
	{label: "21日付费", value: "pay21"},
	{label: "22日付费", value: "pay22"},
	{label: "23日付费", value: "pay23"},
	{label: "24日付费", value: "pay24"},
	{label: "25日付费", value: "pay25"},
	{label: "26日付费", value: "pay26"},
	{label: "27日付费", value: "pay27"},
	{label: "28日付费", value: "pay28"},
	{label: "29日付费", value: "pay29"},
	{label: "30日付费", value: "pay30"},
	{label: "31日付费", value: "pay31"},
	{label: "32日付费", value: "pay32"},
	{label: "33日付费", value: "pay33"},
	{label: "34日付费", value: "pay34"},
	{label: "35日付费", value: "pay35"},
	{label: "36日付费", value: "pay36"},
	{label: "37日付费", value: "pay37"},
	{label: "38日付费", value: "pay38"},
	{label: "39日付费", value: "pay39"},
	{label: "40日付费", value: "pay40"},
	{label: "41日付费", value: "pay41"},
	{label: "42日付费", value: "pay42"},
	{label: "43日付费", value: "pay43"},
	{label: "44日付费", value: "pay44"},
	{label: "45日付费", value: "pay45"},
	{label: "46日付费", value: "pay46"},
	{label: "47日付费", value: "pay47"},
	{label: "48日付费", value: "pay48"},
	{label: "49日付费", value: "pay49"},
	{label: "50日付费", value: "pay50"},
	{label: "51日付费", value: "pay51"},
	{label: "52日付费", value: "pay52"},
	{label: "53日付费", value: "pay53"},
	{label: "54日付费", value: "pay54"},
	{label: "55日付费", value: "pay55"},
	{label: "56日付费", value: "pay56"},
	{label: "57日付费", value: "pay57"},
	{label: "58日付费", value: "pay58"},
	{label: "59日付费", value: "pay59"},
	{label: "60日付费", value: "pay60"},
	{label: "61日付费", value: "pay61"},
	{label: "62日付费", value: "pay62"},
	{label: "63日付费", value: "pay63"},
	{label: "64日付费", value: "pay64"},
	{label: "65日付费", value: "pay65"},
	{label: "66日付费", value: "pay66"},
	{label: "67日付费", value: "pay67"},
	{label: "68日付费", value: "pay68"},
	{label: "69日付费", value: "pay69"},
	{label: "70日付费", value: "pay70"},
	{label: "71日付费", value: "pay71"},
	{label: "72日付费", value: "pay72"},
	{label: "73日付费", value: "pay73"},
	{label: "74日付费", value: "pay74"},
	{label: "75日付费", value: "pay75"},
	{label: "76日付费", value: "pay76"},
	{label: "77日付费", value: "pay77"},
	{label: "78日付费", value: "pay78"},
	{label: "79日付费", value: "pay79"},
	{label: "80日付费", value: "pay80"},
	{label: "81日付费", value: "pay81"},
	{label: "82日付费", value: "pay82"},
	{label: "83日付费", value: "pay83"},
	{label: "84日付费", value: "pay84"},
	{label: "85日付费", value: "pay85"},
	{label: "86日付费", value: "pay86"},
	{label: "87日付费", value: "pay87"},
	{label: "88日付费", value: "pay88"},
	{label: "89日付费", value: "pay89"},
	{label: "90日付费", value: "pay90"},
	{label: "120日付费", value: "pay120"},
	{label: "150日付费", value: "pay150"},
	{label: "180日付费", value: "pay180"},
	{label: "210日付费", value: "pay210"},
	{label: "240日付费", value: "pay240"},
	{label: "270日付费", value: "pay270"},
	{label: "300日付费", value: "pay300"},
	{label: "360日付费", value: "pay360"},
	{label: "累计付费", value: "total_pay"},
	//{label: "首日LTV", value: "ltv1"},
	//{label: "次日LTV", value: "ltv2"},
	//{label: "3日LTV", value: "ltv3"},
	//{label: "4日LTV", value: "ltv4"},
	//{label: "5日LTV", value: "ltv5"},
	//{label: "6日LTV", value: "ltv6"},
	//{label: "7日LTV", value: "ltv7"},
	//{label: "8日LTV", value: "ltv8"},
	//{label: "9日LTV", value: "ltv9"},
	//{label: "10日LTV", value: "ltv10"},
	//{label: "11日LTV", value: "ltv11"},
	//{label: "12日LTV", value: "ltv12"},
	//{label: "13日LTV", value: "ltv13"},
	//{label: "14日LTV", value: "ltv14"},
	//{label: "15日LTV", value: "ltv15"},
	//{label: "30日LTV", value: "ltv30"},
	//{label: "45日LTV", value: "ltv45"},
	//{label: "60日LTV", value: "ltv60"},
	//{label: "90日LTV", value: "ltv90"},
	//{label: "120日LTV", value: "ltv120"},
	//{label: "150日LTV", value: "ltv150"},
	//{label: "180日LTV", value: "ltv180"},
	//{label: "210日LTV", value: "ltv210"},
	//{label: "240日LTV", value: "ltv240"},
	//{label: "270日LTV", value: "ltv270"},
	//{label: "300日LTV", value: "ltv300"},
	//{label: "360日LTV", value: "ltv360"},
	//{label: "累计LTV", value: "total_ltv"},
	{label: "首日付费数", value: "new_paid_user"},
	{label: "新增付费率rpu", value: "new_paid_rate"},
	{label: "2日留存", value: "stay2"},
	{label: "3日留存", value: "stay3"},
	{label: "4日留存", value: "stay4"},
	{label: "5日留存", value: "stay5"},
	{label: "6日留存", value: "stay6"},
	{label: "7日留存", value: "stay7"},
	{label: "8日留存", value: "stay8"},
	{label: "9日留存", value: "stay9"},
	{label: "10日留存", value: "stay10"},
	{label: "11日留存", value: "stay11"},
	{label: "12日留存", value: "stay12"},
	{label: "13日留存", value: "stay13"},
	{label: "14日留存", value: "stay14"},
	{label: "15日留存", value: "stay15"},
	{label: "30日留存", value: "stay30"},
	{label: "45日留存", value: "stay45"},
	{label: "60日留存", value: "stay60"},
	{label: "90日留存", value: "stay90"},
	{label: "120日留存", value: "stay120"},
	{label: "150日留存", value: "stay150"},
	{label: "180日留存", value: "stay180"},
	{label: "DAU", value: "dau"},
	{label: "DAU流水", value: "inapp_purchase_flow"},
	{label: "付费DAU", value: "paid_dau"},
	{label: "活跃付费人数", value: "active_paid_user"},
	{label: "次三比", value: "times3"},
	{label: "次四比", value: "times4"},
	{label: "次五比", value: "times5"},
	{label: "次六比", value: "times6"},
	{label: "次七比", value: "times7"},
	{label: "付费次留率", value: "pay2_stay_rate"},
	{label: "付费七留率", value: "pay7_stay_rate"},
	{label: "活跃arpu", value: "active_arpu"},
	{label: "付费arpu", value: "paid_arpu"},
}
var diemColumnList = []*SelectItem{
	columnDate,
	columnDeviceOS,
	columnProjectName,
	columnStudioDisplay,
	columnOptimizerName,
	columnPlatform,
	columnAdvertiserID,
	columnAdvertiserNameDisplay,
	columnAdGroupID,
	columnAdGroupNameDisplay,
	columnAdID,
	columnAdNameDisplay,
}
var (
	columnMap                   = map[string]*SelectItem{}
	selectColumn                = ""
	columnDate                  = &SelectItem{calculate: "date", as: "date"}                                                     // 日期
	columnMonth                 = &SelectItem{calculate: "month", as: "month"}                                                   // 月
	columnDeviceOS              = &SelectItem{calculate: "device_os", as: "device_os"}                                           // 设备系统
	columnProject               = &SelectItem{calculate: "project_id", as: "project_id"}                                         // 应用
	columnProjectName           = &SelectItem{calculate: "projects.name", as: "project_name"}                                    // 应用
	columnStudio                = &SelectItem{calculate: "studio_id", as: "studio_id"}                                           // 工作室
	columnStudioDisplay         = &SelectItem{calculate: "studios.name", as: "studio_name"}                                      // 工作室名称
	columnStudioName            = &SelectItem{calculate: "studios.name", as: "studio_name"}                                      // 工作室
	columnOptimizer             = &SelectItem{calculate: "optimizer_id", as: "optimizer_id"}                                     // 优化师id
	columnOptimizerName         = &SelectItem{calculate: "users.name", as: "optimizer_name"}                                     // 优化师
	columnPlatform              = &SelectItem{calculate: "platform", as: "platform_id"}                                          // 渠道
	columnAdvertiserID          = &SelectItem{calculate: "advertiser_id", as: "advertiser_id"}                                   // 广告主ID
	columnAdvertiserName        = &SelectItem{calculate: "advertiser_name", as: "advertiser_name"}                               // 广告主名称
	columnAdvertiserNameDisplay = &SelectItem{calculate: "report_platform_advertiser_names.name", as: "advertiser_name_display"} // 广告主名称
	columnAdGroupID             = &SelectItem{calculate: "adgroup_id", as: "adgroup_id"}                                         // 广告组id
	columnAdGroupName           = &SelectItem{calculate: "adgroup_name", as: "adgroup_name"}                                     // 广告组名称
	columnAdGroupNameDisplay    = &SelectItem{calculate: "report_platform_ad_group_names.name", as: "adgroup_name_display"}      // 广告组名称
	columnAdID                  = &SelectItem{calculate: "ad_id", as: "ad_id"}                                                   // 广告id
	columnAdName                = &SelectItem{calculate: "ad_name", as: "ad_name"}                                               // 广告名称
	columnAdNameDisplay         = &SelectItem{calculate: "report_platform_ad_names.name", as: "ad_name_display"}                 // 广告名称显示
	SelectColumnSummary         = []*SelectItem{
		{"ROUND(SUM(`cost`), 2)", "cost"},
		{"SUM(`show`)", "show"},
		{"SUM(`click`)", "click"},
		{"SUM(`convert`)", "convert"},
		{"SUM(`active`)", "active"},
		{"ROUND(SUM(`cost`)/SUM(`convert`), 2)", "convert_cost"},
		{"ROUND(SUM(`cost`)/SUM(`active`), 2)", "active_cost"},
		{"ROUND(SUM(`cost`)*1000/SUM(`show`), 2)", "cost_by_thousand_show"},
		{"ROUND(SUM(`click`)/SUM(`show`), 4)", "ctr"},
		{"ROUND(SUM(`convert`)/SUM(`click`), 4)", "cvr"},
		{"SUM(`new_user`)", "new_user"},
		{"ROUND(SUM(`pay1`), 0)", "pay1"},
		{"ROUND(SUM(`pay2`), 0)", "pay2"},
		{"ROUND(SUM(`pay3`), 0)", "pay3"},
		{"ROUND(SUM(`pay4`), 0)", "pay4"},
		{"ROUND(SUM(`pay5`), 0)", "pay5"},
		{"ROUND(SUM(`pay6`), 0)", "pay6"},
		{"ROUND(SUM(`pay7`), 0)", "pay7"},
		{"ROUND(SUM(`pay8`), 0)", "pay8"},
		{"ROUND(SUM(`pay9`), 0)", "pay9"},
		{"ROUND(SUM(`pay10`), 0)", "pay10"},
		{"ROUND(SUM(`pay11`), 0)", "pay11"},
		{"ROUND(SUM(`pay12`), 0)", "pay12"},
		{"ROUND(SUM(`pay13`), 0)", "pay13"},
		{"ROUND(SUM(`pay14`), 0)", "pay14"},
		{"ROUND(SUM(`pay15`), 0)", "pay15"},
		{"ROUND(SUM(`pay30`), 0)", "pay30"},
		{"ROUND(SUM(`pay45`), 0)", "pay45"},
		{"ROUND(SUM(`pay60`), 0)", "pay60"},
		{"ROUND(SUM(`pay90`), 0)", "pay90"},
		{"ROUND(SUM(`pay120`), 0)", "pay120"},
		{"ROUND(SUM(`pay150`), 0)", "pay150"},
		{"ROUND(SUM(`pay180`), 0)", "pay180"},
		{"ROUND(SUM(`pay210`), 0)", "pay210"},
		{"ROUND(SUM(`pay240`), 0)", "pay240"},
		{"ROUND(SUM(`pay270`), 0)", "pay270"},
		{"ROUND(SUM(`pay300`), 0)", "pay300"},
		{"ROUND(SUM(`pay360`), 0)", "pay360"},
		{"ROUND(SUM(`total_pay`), 0)", "total_pay"},
		{"ROUND(SUM(`pay1`)/SUM(`new_user`), 2)", "ltv1"},
		{"ROUND(SUM(`pay2`)/SUM(`new_user`), 2)", "ltv2"},
		{"ROUND(SUM(`pay3`)/SUM(`new_user`), 2)", "ltv3"},
		{"ROUND(SUM(`pay4`)/SUM(`new_user`), 2)", "ltv4"},
		{"ROUND(SUM(`pay5`)/SUM(`new_user`), 2)", "ltv5"},
		{"ROUND(SUM(`pay6`)/SUM(`new_user`), 2)", "ltv6"},
		{"ROUND(SUM(`pay7`)/SUM(`new_user`), 2)", "ltv7"},
		{"ROUND(SUM(`pay8`)/SUM(`new_user`), 2)", "ltv8"},
		{"ROUND(SUM(`pay9`)/SUM(`new_user`), 2)", "ltv9"},
		{"ROUND(SUM(`pay10`)/SUM(`new_user`), 2)", "ltv10"},
		{"ROUND(SUM(`pay11`)/SUM(`new_user`), 2)", "ltv11"},
		{"ROUND(SUM(`pay12`)/SUM(`new_user`), 2)", "ltv12"},
		{"ROUND(SUM(`pay13`)/SUM(`new_user`), 2)", "ltv13"},
		{"ROUND(SUM(`pay14`)/SUM(`new_user`), 2)", "ltv14"},
		{"ROUND(SUM(`pay15`)/SUM(`new_user`), 2)", "ltv15"},
		{"ROUND(SUM(`pay30`)/SUM(`new_user`), 2)", "ltv30"},
		{"ROUND(SUM(`pay45`)/SUM(`new_user`), 2)", "ltv45"},
		{"ROUND(SUM(`pay60`)/SUM(`new_user`), 2)", "ltv60"},
		{"ROUND(SUM(`pay90`)/SUM(`new_user`), 2)", "ltv90"},
		{"ROUND(SUM(`pay120`)/SUM(`new_user`), 2)", "ltv120"},
		{"ROUND(SUM(`pay150`)/SUM(`new_user`), 2)", "ltv150"},
		{"ROUND(SUM(`pay180`)/SUM(`new_user`), 2)", "ltv180"},
		{"ROUND(SUM(`pay210`)/SUM(`new_user`), 2)", "ltv210"},
		{"ROUND(SUM(`pay240`)/SUM(`new_user`), 2)", "ltv240"},
		{"ROUND(SUM(`pay270`)/SUM(`new_user`), 2)", "ltv270"},
		{"ROUND(SUM(`pay300`)/SUM(`new_user`), 2)", "ltv300"},
		{"ROUND(SUM(`pay360`)/SUM(`new_user`), 2)", "ltv360"},
		{"ROUND(SUM(`total_pay`)/SUM(`new_user`), 2)", "total_ltv"},
		{"ROUND(SUM(`pay1`)/SUM(`cost`), 4)", "roi1"},
		{"ROUND(SUM(`pay2`)/SUM(`cost`), 4)", "roi2"},
		{"ROUND(SUM(`pay3`)/SUM(`cost`), 4)", "roi3"},
		{"ROUND(SUM(`pay4`)/SUM(`cost`), 4)", "roi4"},
		{"ROUND(SUM(`pay5`)/SUM(`cost`), 4)", "roi5"},
		{"ROUND(SUM(`pay6`)/SUM(`cost`), 4)", "roi6"},
		{"ROUND(SUM(`pay7`)/SUM(`cost`), 4)", "roi7"},
		{"ROUND(SUM(`pay8`)/SUM(`cost`), 4)", "roi8"},
		{"ROUND(SUM(`pay9`)/SUM(`cost`), 4)", "roi9"},
		{"ROUND(SUM(`pay10`)/SUM(`cost`), 4)", "roi10"},
		{"ROUND(SUM(`pay11`)/SUM(`cost`), 4)", "roi11"},
		{"ROUND(SUM(`pay12`)/SUM(`cost`), 4)", "roi12"},
		{"ROUND(SUM(`pay13`)/SUM(`cost`), 4)", "roi13"},
		{"ROUND(SUM(`pay14`)/SUM(`cost`), 4)", "roi14"},
		{"ROUND(SUM(`pay15`)/SUM(`cost`), 4)", "roi15"},
		{"ROUND(SUM(`pay30`)/SUM(`cost`), 4)", "roi30"},
		{"ROUND(SUM(`pay45`)/SUM(`cost`), 4)", "roi45"},
		{"ROUND(SUM(`pay60`)/SUM(`cost`), 4)", "roi60"},
		{"ROUND(SUM(`pay90`)/SUM(`cost`), 4)", "roi90"},
		{"ROUND(SUM(`pay120`)/SUM(`cost`), 4)", "roi120"},
		{"ROUND(SUM(`pay150`)/SUM(`cost`), 4)", "roi150"},
		{"ROUND(SUM(`pay180`)/SUM(`cost`), 4)", "roi180"},
		{"ROUND(SUM(`pay210`)/SUM(`cost`), 4)", "roi210"},
		{"ROUND(SUM(`pay240`)/SUM(`cost`), 4)", "roi240"},
		{"ROUND(SUM(`pay270`)/SUM(`cost`), 4)", "roi270"},
		{"ROUND(SUM(`pay300`)/SUM(`cost`), 4)", "roi300"},
		{"ROUND(SUM(`pay360`)/SUM(`cost`), 4)", "roi360"},
		{"ROUND(SUM(`total_pay`)/SUM(`cost`), 4)", "total_roi"},
		{"SUM(`new_paid_user`)", "new_paid_user"},
		{"ROUND(SUM(`new_paid_user`)/SUM(`new_user`), 4)", "new_paid_rate"},
		{"ROUND(SUM(`stay2`*`new_user`)/SUM(`new_user`), 4)", "stay2"},
		{"ROUND(SUM(`stay3`*`new_user`)/SUM(`new_user`), 4)", "stay3"},
		{"ROUND(SUM(`stay4`*`new_user`)/SUM(`new_user`), 4)", "stay4"},
		{"ROUND(SUM(`stay5`*`new_user`)/SUM(`new_user`), 5)", "stay5"},
		{"ROUND(SUM(`stay6`*`new_user`)/SUM(`new_user`), 6)", "stay6"},
		{"ROUND(SUM(`stay7`*`new_user`)/SUM(`new_user`), 7)", "stay7"},
		{"ROUND(SUM(`stay8`*`new_user`)/SUM(`new_user`), 8)", "stay8"},
		{"ROUND(SUM(`stay9`*`new_user`)/SUM(`new_user`), 9)", "stay9"},
		{"ROUND(SUM(`stay10`*`new_user`)/SUM(`new_user`), 4)", "stay10"},
		{"ROUND(SUM(`stay11`*`new_user`)/SUM(`new_user`), 4)", "stay11"},
		{"ROUND(SUM(`stay12`*`new_user`)/SUM(`new_user`), 4)", "stay12"},
		{"ROUND(SUM(`stay13`*`new_user`)/SUM(`new_user`), 4)", "stay13"},
		{"ROUND(SUM(`stay14`*`new_user`)/SUM(`new_user`), 4)", "stay14"},
		{"ROUND(SUM(`stay15`*`new_user`)/SUM(`new_user`), 4)", "stay15"},
		{"ROUND(SUM(`stay30`*`new_user`)/SUM(`new_user`), 4)", "stay30"},
		{"ROUND(SUM(`stay45`*`new_user`)/SUM(`new_user`), 4)", "stay45"},
		{"ROUND(SUM(`stay60`*`new_user`)/SUM(`new_user`), 4)", "stay60"},
		{"ROUND(SUM(`stay90`*`new_user`)/SUM(`new_user`), 4)", "stay90"},
		{"ROUND(SUM(`stay120`*`new_user`)/SUM(`new_user`), 4)", "stay120"},
		{"ROUND(SUM(`stay150`*`new_user`)/SUM(`new_user`), 4)", "stay150"},
		{"ROUND(SUM(`stay180`*`new_user`)/SUM(`new_user`), 4)", "stay180"},
		{"ROUND(SUM(`stay3`)/SUM(`stay2`), 4)", "times3"},
		{"ROUND(SUM(`stay4`)/SUM(`stay2`), 4)", "times4"},
		{"ROUND(SUM(`stay5`)/SUM(`stay2`), 4)", "times5"},
		{"ROUND(SUM(`stay6`)/SUM(`stay2`), 4)", "times6"},
		{"ROUND(SUM(`stay7`)/SUM(`stay2`), 4)", "times7"},
		{"ROUND(SUM(`pay2_stay_rate`*`new_paid_user`)/SUM(`new_paid_user`), 4)", "pay2_stay_rate"},
		{"ROUND(SUM(`pay7_stay_rate`*`new_paid_user`)/SUM(`new_paid_user`), 4)", "pay7_stay_rate"},
		{"ROUND(SUM(`inapp_purchase_flow`)/SUM(`paid_dau`), 2)", "paid_arpu"},
		{"ROUND(SUM(`inapp_purchase_flow`)/SUM(`dau`), 2)", "active_arpu"},
		{"ROUND(SUM(`inapp_purchase_flow`), 2)", "inapp_purchase_flow"},
		{"ROUND(SUM(`cost`)/SUM(`new_user`), 2)", "new_user_cost"},
		{"ROUND(SUM(`cost`)/SUM(`new_paid_user`), 2)", "new_paid_user_cost"},
	}
)

func (si *SelectItem) String() string {
	if si.calculate == "" {
		return fmt.Sprintf("`%s`", si.as)
	}
	return fmt.Sprintf("%s as `%s`", si.calculate, si.as)
}

func getLine(cd []*ColumnDefine, line *ReportPlatformItem) []string {
	ret := make([]string, 0, len(cd))
	for _, v := range cd {
		ret = append(ret, getCellValue(v, line))
	}
	return ret
}

func getCellValue(cd *ColumnDefine, line *ReportPlatformItem) string {
	switch cd.value {
	case "date":
		return getDateDisplay(line.Date)
	case "month":
		return getMonthDisplay(line.Month)
	case "project_id":
		return line.ProjectId
	case "project_name":
		return line.ProjectName
	case "studio_name":
		return line.StudioName
	case "optimizer_name":
		return line.OptimizerName
	case "optimizer_id":
		return line.OptimizerId
	case "platform_id":
		return "666"
		//return enum.GetPlatformDisplayByIDStr(line.PlatformId)
	case "device_os":
		return "ios"
		//return enum.DeviceOS(line.DeviceOs).String()
	case "advertiser_id":
		return line.AdvertiserId
	case "advertiser_name_display":
		return line.AdvertiserNameDisplay
	case "adgroup_id":
		return line.AdgroupId
	case "adgroup_name_display":
		return line.AdgroupNameDisplay
	case "ad_id":
		return line.AdId
	case "ad_name_display":
		return line.AdNameDisplay
	case "cost":
		return line.Cost
	case "show":
		return strconv.FormatInt(line.Show, 10)
	case "click":
		return strconv.FormatInt(line.Click, 10)
	case "convert":
		return strconv.FormatInt(line.Convert, 10)
	case "active":
		return strconv.FormatInt(line.Active, 10)
	case "convert_cost":
		return line.ConvertCost
	case "active_cost":
		return line.ActiveCost
	case "cost_by_thousand_show":
		return line.CostByThousandShow
	case "ctr":
		return line.Ctr
	case "cvr":
		return line.Cvr
	case "new_user":
		return strconv.FormatInt(line.NewUser, 10)
	case "pay1":
		return strconv.FormatInt(line.Pay1, 10)
	case "pay2":
		return strconv.FormatInt(line.Pay2, 10)
	case "pay3":
		return strconv.FormatInt(line.Pay3, 10)
	case "pay4":
		return strconv.FormatInt(line.Pay4, 10)
	case "pay5":
		return strconv.FormatInt(line.Pay5, 10)
	case "pay6":
		return strconv.FormatInt(line.Pay6, 10)
	case "pay7":
		return strconv.FormatInt(line.Pay7, 10)
	case "pay8":
		return strconv.FormatInt(line.Pay8, 10)
	case "pay9":
		return strconv.FormatInt(line.Pay9, 10)
	case "pay10":
		return strconv.FormatInt(line.Pay10, 10)
	case "pay11":
		return strconv.FormatInt(line.Pay11, 10)
	case "pay12":
		return strconv.FormatInt(line.Pay12, 10)
	case "pay13":
		return strconv.FormatInt(line.Pay13, 10)
	case "pay14":
		return strconv.FormatInt(line.Pay14, 10)
	case "pay15":
		return strconv.FormatInt(line.Pay15, 10)
	case "pay30":
		return strconv.FormatInt(line.Pay30, 10)
	case "pay45":
		return strconv.FormatInt(line.Pay45, 10)
	case "pay60":
		return strconv.FormatInt(line.Pay60, 10)
	case "pay90":
		return strconv.FormatInt(line.Pay90, 10)
	case "pay120":
		return strconv.FormatInt(line.Pay120, 10)
	case "pay150":
		return strconv.FormatInt(line.Pay150, 10)
	case "pay180":
		return strconv.FormatInt(line.Pay180, 10)
	case "pay210":
		return strconv.FormatInt(line.Pay210, 10)
	case "pay240":
		return strconv.FormatInt(line.Pay240, 10)
	case "pay270":
		return strconv.FormatInt(line.Pay270, 10)
	case "pay300":
		return strconv.FormatInt(line.Pay300, 10)
	case "pay360":
		return strconv.FormatInt(line.Pay360, 10)
	case "total_pay":
		return strconv.FormatInt(line.TotalPay, 10)
	case "ltv1":
		return line.Ltv1
	case "ltv2":
		return line.Ltv2
	case "ltv3":
		return line.Ltv3
	case "ltv4":
		return line.Ltv4
	case "ltv5":
		return line.Ltv5
	case "ltv6":
		return line.Ltv6
	case "ltv7":
		return line.Ltv7
	case "ltv8":
		return line.Ltv8
	case "ltv9":
		return line.Ltv9
	case "ltv10":
		return line.Ltv10
	case "ltv11":
		return line.Ltv11
	case "ltv12":
		return line.Ltv12
	case "ltv13":
		return line.Ltv13
	case "ltv14":
		return line.Ltv14
	case "ltv15":
		return line.Ltv15
	case "ltv30":
		return line.Ltv30
	case "ltv45":
		return line.Ltv45
	case "ltv60":
		return line.Ltv60
	case "ltv90":
		return line.Ltv90
	case "ltv120":
		return line.Ltv120
	case "ltv150":
		return line.Ltv150
	case "ltv180":
		return line.Ltv180
	case "ltv210":
		return line.Ltv210
	case "ltv240":
		return line.Ltv240
	case "ltv270":
		return line.Ltv270
	case "ltv300":
		return line.Ltv300
	case "ltv360":
		return line.Ltv360
	case "total_ltv":
		return line.TotalLtv
	case "roi1":
		return line.Roi1
	case "roi2":
		return line.Roi2
	case "roi3":
		return line.Roi3
	case "roi4":
		return line.Roi4
	case "roi5":
		return line.Roi5
	case "roi6":
		return line.Roi6
	case "roi7":
		return line.Roi7
	case "roi8":
		return line.Roi8
	case "roi9":
		return line.Roi9
	case "roi10":
		return line.Roi10
	case "roi11":
		return line.Roi11
	case "roi12":
		return line.Roi12
	case "roi13":
		return line.Roi13
	case "roi14":
		return line.Roi14
	case "roi15":
		return line.Roi15
	case "roi30":
		return line.Roi30
	case "roi45":
		return line.Roi45
	case "roi60":
		return line.Roi60
	case "roi90":
		return line.Roi90
	case "roi120":
		return line.Roi120
	case "roi150":
		return line.Roi150
	case "roi180":
		return line.Roi180
	case "roi210":
		return line.Roi210
	case "roi240":
		return line.Roi240
	case "roi270":
		return line.Roi270
	case "roi300":
		return line.Roi300
	case "roi360":
		return line.Roi360
	case "total_roi":
		return line.TotalRoi
	case "new_paid_user":
		return strconv.FormatInt(line.NewPaidUser, 10)
	case "new_paid_rate":
		return line.NewPaidRate
	case "stay2":
		return line.Stay2
	case "stay3":
		return line.Stay3
	case "stay4":
		return line.Stay4
	case "stay5":
		return line.Stay5
	case "stay6":
		return line.Stay6
	case "stay7":
		return line.Stay7
	case "stay8":
		return line.Stay8
	case "stay9":
		return line.Stay9
	case "stay10":
		return line.Stay10
	case "stay11":
		return line.Stay11
	case "stay12":
		return line.Stay12
	case "stay13":
		return line.Stay13
	case "stay14":
		return line.Stay14
	case "stay15":
		return line.Stay15
	case "stay30":
		return line.Stay30
	case "stay45":
		return line.Stay45
	case "stay60":
		return line.Stay60
	case "stay90":
		return line.Stay90
	case "stay120":
		return line.Stay120
	case "stay150":
		return line.Stay150
	case "stay180":
		return line.Stay180
	case "times3":
		return line.Times3
	case "times4":
		return line.Times4
	case "times5":
		return line.Times5
	case "times6":
		return line.Times6
	case "times7":
		return line.Times7
	case "pay2_stay_rate":
		return line.Pay2StayRate
	case "pay7_stay_rate":
		return line.Pay7StayRate
	case "active_arpu":
		return line.ActiveArpu
	case "paid_arpu":
		return line.PaidArpu
	case "new_user_cost":
		return line.NewUserCost
	case "new_paid_user_cost":
		return line.NewPaidUserCost
	}
	return ""
}

func getDateDisplay(date string) string {
	if len(date) == 8 {
		return fmt.Sprintf("%s-%s-%s", date[0:4], date[4:6], date[6:8])
	} else {
		return date
	}
}

func getMonthDisplay(month string) string {
	if len(month) == 6 {
		return fmt.Sprintf("%s-%s", month[0:4], month[4:6])
	} else {
		return month
	}
}
