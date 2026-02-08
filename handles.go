package main

import (
	"net/http"
	"strings"
)

func router(mux *http.ServeMux) {
	mux.Handle("/api/config", handleConfig())
	mux.Handle("/api/finding/index", handleFinding())
	mux.Handle("/api/menu/control/list/{page}", handleMenu())
	mux.Handle("/api/we/run/index", handleWeRun())
	mux.Handle("/api/base/notice/page/{page}", handleNotice())
	mux.Handle("/api/portal/login/off_line", handleLoginOffline())
	mux.Handle("/api/sso/settings/{page}", handleSsoSetting())
	mux.Handle("/", handleDefault())
}

func handleConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"data":{"fileConfig":{"sushe":"https://upload-images.jianshu.io/upload_images/29040124-4842c3857b5d38e0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","pwdQuestion":"https://upload-images.jianshu.io/upload_images/29040124-0a13218e3233e7e5.png","mapiconNaviS":"https://upload-images.jianshu.io/upload_images/29040124-bdf177a17381ddcf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","xingzheng":"https://upload-images.jianshu.io/upload_images/29040124-914e260511260876.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","bannerImage":"https://upload-images.jianshu.io/upload_images/29040124-3d29c26b360da539.jpg?imageMogr2/auto-orient/strip|imageView2/2/w/1039/format/webp","defaultAvatar":"https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0","emptyCart":"https://upload-images.jianshu.io/upload_images/29040124-3ecea4f84d7baddb.png","schoolAerialView":"https://upload-images.jianshu.io/upload_images/29040124-c7c837ba17275c02.jpg?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp","mapiconNaviE":"https://upload-images.jianshu.io/upload_images/29040124-ff3f12e38d191edf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","others":"https://upload-images.jianshu.io/upload_images/29040124-fb4d2a5f3574108d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","parcel":"https://upload-images.jianshu.io/upload_images/29040124-aed6c371a0b004a6.png","myBg":"https://upload-images.jianshu.io/upload_images/4697920-bf033361b1498e88.gif","myBackgroundImage":"https://upload-images.jianshu.io/upload_images/4697920-48dab9eddafb6ce3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","achievementDoc":"https://fvti.oss-cn-hangzhou.aliyuncs.com/file/%E3%80%8A%E7%A6%8F%E5%B7%9E%E8%81%8C%E4%B8%9A%E6%8A%80%E6%9C%AF%E5%AD%A6%E9%99%A2%E5%AD%A6%E7%94%9F%E7%BB%BC%E5%90%88%E7%B4%A0%E8%B4%A8%E6%B5%8B%E8%AF%84%E5%8A%9E%E6%B3%95%E3%80%8B%28%E6%A6%95%E8%81%8C%E9%99%A2%E7%BB%BC%E3%80%942017%E3%80%9537%E5%8F%B7%29%282%29%20%282%29.doc","sports":"https://upload-images.jianshu.io/upload_images/29040124-75aa6adb05276f5b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","staticMap":"https://upload-images.jianshu.io/upload_images/29040124-83b69b66f3e83332.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","markerChecked":"https://upload-images.jianshu.io/upload_images/29040124-22fa5c5c5e3d8283.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","rank1":"https://upload-images.jianshu.io/upload_images/29040124-7433743505931949.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","rank2":"https://upload-images.jianshu.io/upload_images/29040124-40e0541e3a59ecd1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","jiaoxuelou":"https://upload-images.jianshu.io/upload_images/29040124-5563267f08b1c6e4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","moveSchool":"https://upload-images.jianshu.io/upload_images/29040124-a66cc850142e27bd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","rank3":"https://upload-images.jianshu.io/upload_images/29040124-8ceeb416b3ea6796.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","xiaomen":"https://upload-images.jianshu.io/upload_images/29040124-5d9f0a4e69d9f60e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","marker":"https://upload-images.jianshu.io/upload_images/29040124-b6be997b20cf0fb4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","loginCircleImage":"https://upload-images.jianshu.io/upload_images/4697920-3401f7949a9e8b5c.gif?imageMogr2/auto-orient/|imageView2","shitang":"https://upload-images.jianshu.io/upload_images/29040124-144b839a7b42afef.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","location":"https://upload-images.jianshu.io/upload_images/29040124-b3e22eb23325dd8e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","loadingGif":"https://upload-images.jianshu.io/upload_images/29040124-d83bca21b3f262c0.gif?imageMogr2/auto-orient/strip","rank4":"https://upload-images.jianshu.io/upload_images/29040124-5fd1531e618a568c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","schoolLogo":"https://upload-images.jianshu.io/upload_images/29040124-2e540a0b39ae228e.png?imageMogr2/auto-orient/strip|imageView2","rank5":"https://upload-images.jianshu.io/upload_images/29040124-89e35968aaab808e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","rank6":"https://upload-images.jianshu.io/upload_images/29040124-b3867ff95a5a0e10.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","courseTime":"https://upload-images.jianshu.io/upload_images/29040124-7540663d66aa3ce9.png"},"styleConfig":{"courseOfflineAdStyle":"fixed"}},"success":true,"server":"web1"}`))
	}
}

func handleFinding() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"data":[],"success":true,"server":"web1"}`))
	}
}

func handleMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := ExtractParams(r.Pattern, r.RequestURI)["page"]
		if page == "INDEX" {
			w.Write([]byte(`{"status":200,"data":[{"id":92,"menu":"新生找社团","menuName":"club","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":3,"bindtap":"toPageUnExamine","bindtapParams":"/pages/club/club","opentype":"","isDeleted":false,"createTime":"2023-09-03 00:17:58","updateTime":"2025-02-16 19:11:22"},{"id":1,"menu":"在线课表","menuName":"course","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":5,"bindtap":"toPage","bindtapParams":"/pages/course/course","opentype":"","isDeleted":false,"updateTime":"2025-02-16 19:11:22"},{"id":2,"menu":"成绩查询","menuName":"achievement","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":9,"bindtap":"toPage","bindtapParams":"/pages/achievement/achievement","opentype":"","isDeleted":false,"updateTime":"2025-02-16 19:11:22"},{"id":3,"menu":"考勤记录","menuName":"attendance","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":10,"bindtap":"toPage","bindtapParams":"/pages/attendance/attendance","opentype":"","isDeleted":false,"updateTime":"2025-02-16 19:11:22"},{"id":4,"menu":"去请假","menuName":"leave","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":11,"bindtap":"toPage","bindtapParams":"/pages/leave/leave","isDeleted":false,"updateTime":"2025-02-16 19:11:22"},{"id":82,"menu":"学业分析","menuName":"ana","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":13,"bindtap":"toPage","bindtapParams":"/pages/achievement/ana/ana","opentype":"","isDeleted":false,"updateTime":"2025-02-16 19:11:22"},{"id":61,"menu":"分享","menuName":"share","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":15,"bindtap":"","bindtapParams":"","opentype":"share","isDeleted":false,"updateTime":"2025-02-16 19:11:22"},{"id":16,"menu":"人工客服","menuName":"admin","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":18,"opentype":"contact","isDeleted":false,"updateTime":"2025-02-16 19:11:24"},{"id":84,"menu":"学业分排行","menuName":"achievement_score","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":18,"bindtap":"toPageUnExamine","bindtapParams":"/pages/achievement_score/achievement_score","opentype":"","isDeleted":false,"updateTime":"2025-02-16 19:11:24"},{"id":106,"menu":"确认课时","menuName":"confirm_class","menuColor":"#666","enable":true,"openAd":false,"page":"INDEX","orderNo":20,"bindtap":"toPage","bindtapParams":"/pages/confirm_class/confirm_class","opentype":"","isNew":false,"isDeleted":false,"createTime":"2024-02-27 23:32:11","updateTime":"2025-02-16 19:11:24"}],"success":true,"server":"web1"}`))
		} else if page == "APPLICATION" {
			w.Write([]byte(`{"status":200,"data":[{"id":96,"menu":"学业查询_激活新密码","menuName":"summary","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":1,"bindtap":"openArticle","bindtapParams":"https://mp.weixin.qq.com/s?__biz=MzkyMjIyODAzOA==&mid=2247483704&idx=1&sn=1d8155bcaf389801a98e331b3388d285&chksm=c1f6c511f6814c07ea43a71b097b8e4f79a0e7918c9574668253b3dfd5a3628eb50f9acab37d#rd","opentype":"","isDeleted":false,"createTime":"2023-09-05 21:38:57","updateTime":"2023-10-10 16:33:31"},{"id":67,"menu":"学业查询_学期课表","menuName":"class_course","menuColor":"#666","enable":true,"openAd":true,"page":"APPLICATION","orderNo":3,"bindtap":"toPage","bindtapParams":"/pages/course_class/course_class","opentype":"","isNew":false,"isDeleted":false,"updateTime":"2023-10-10 16:33:29"},{"id":75,"menu":"学业查询_学业地图","menuName":"study_map","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":5,"bindtap":"toPage","bindtapParams":"/pages/study_map/study_map","opentype":"","isNew":false,"isDeleted":false,"updateTime":"2025-02-09 02:07:40"},{"id":25,"menu":"学业查询_学业分排名（旧）","menuName":"achievement_term_rank","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":7,"bindtap":"toPage","bindtapParams":"/pages/achievement_term_rank/achievement_term_rank","opentype":"","isDeleted":false,"updateTime":"2025-02-09 02:07:40"},{"id":70,"menu":"学业查询_选修科目","menuName":"select_course","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":8,"bindtap":"toPage","bindtapParams":"/pages/select_course/select_course","opentype":"","isNew":false,"isDeleted":false,"updateTime":"2025-02-09 02:07:40"},{"id":68,"menu":"学业查询_毕业条件","menuName":"graduation","menuColor":"#666","enable":true,"openAd":true,"page":"APPLICATION","orderNo":10,"bindtap":"toPage","bindtapParams":"/pages/graduation/graduation","opentype":"","isNew":false,"isDeleted":false,"updateTime":"2025-02-09 02:09:11"},{"id":93,"menu":"校园服务_新生选课","menuName":"share","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":12,"bindtap":"openArticle","bindtapParams":"https://mp.weixin.qq.com/s?__biz=MzkyMjIyODAzOA==&mid=2247483686&idx=1&sn=7fc8e400d5d5926be49e8060b721f10a&chksm=c1f6c50ff6814c199e2ded2f0a76d22a80469e5ad48da7e081346443b8408b35761a43df15b8#rd","opentype":"","isDeleted":false,"createTime":"2023-09-05 18:27:33","updateTime":"2025-02-09 02:09:11"},{"id":76,"menu":"校园服务_校历","menuName":"calendar","menuColor":"#666","enable":true,"openAd":true,"page":"APPLICATION","orderNo":13,"bindtap":"previewImages","bindtapParams":"https://upload-images.jianshu.io/upload_images/29040124-d6afc7f12ec606cf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240;https://upload-images.jianshu.io/upload_images/29040124-a616c90cf0dac972.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240","opentype":"","isNew":false,"isDeleted":false,"updateTime":"2025-02-09 02:09:11"},{"id":77,"menu":"校园服务_校园导航","menuName":"navi","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":14,"bindtap":"toPageUnExamine","bindtapParams":"/pages/navi/navi","opentype":"","isNew":false,"isDeleted":false,"updateTime":"2025-02-09 02:09:11"},{"id":78,"menu":"校园服务_考试助手","menuName":"cert","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":15,"bindtap":"toPageUnExamine","bindtapParams":"/pages/cert/cert","opentype":"","isNew":false,"isDeleted":false,"updateTime":"2025-02-09 02:09:11"},{"id":71,"menu":"校园服务_宿舍报修","menuName":"repair","menuColor":"#666","enable":true,"openAd":true,"page":"APPLICATION","orderNo":16,"bindtap":"toOtherMP","bindtapParams":"wx680e5afb981350c6","opentype":"","isNew":false,"isDeleted":false,"updateTime":"2025-02-09 02:09:11"},{"id":32,"menu":"校园服务_四六级","menuName":"cet46","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":18,"bindtap":"toOtherMP","bindtapParams":"wxa56afc785454c86b","opentype":"","isDeleted":false,"updateTime":"2025-02-09 02:09:11"},{"id":88,"menu":"校园服务_运动排行","menuName":"sport","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":19,"bindtap":"toPageUnExamine","bindtapParams":"/pages/sport/sport","opentype":"","isDeleted":false,"createTime":"2023-08-29 18:05:58","updateTime":"2025-02-09 02:09:11"},{"id":91,"menu":"娱乐活动_商务合作","menuName":"qunliao","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":20,"bindtap":"openArticle","bindtapParams":"https://mp.weixin.qq.com/s/a6PmeGQ4eDGEDuY0aoWcIQ","opentype":"","isDeleted":false,"createTime":"2023-09-02 13:41:45","updateTime":"2025-02-09 02:09:11"},{"id":89,"menu":"校园服务_赞赏排行","menuName":"sponsor","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":21,"bindtap":"toPageNeedStudentId","bindtapParams":"/pages/sponsor/sponsor","opentype":"","isDeleted":false,"createTime":"2023-09-01 17:52:14","updateTime":"2025-02-09 02:09:11"},{"id":86,"menu":"校园服务_作息表","menuName":"rest","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":23,"bindtap":"openArticle","bindtapParams":"https://mp.weixin.qq.com/s?__biz=MzkyMjIyODAzOA==&mid=2247483661&idx=1&sn=5c22989bfde3bb0f51acc3c4aa56021d&chksm=c1f6c524f6814c32c332d19f89c10d7e71ed2ed56cbaea142323c5f7290693d9f0a27afd58de#rd","opentype":"","isDeleted":false,"updateTime":"2025-02-09 02:09:11"},{"id":87,"menu":"校园服务_找社团","menuName":"club","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":24,"bindtap":"toPageUnExamine","bindtapParams":"/pages/club/club","opentype":"","isDeleted":false,"updateTime":"2025-02-09 02:09:11"},{"id":94,"menu":"学业查询_绩点排行","menuName":"achievement_score_dot","menuColor":"#666","enable":true,"openAd":false,"page":"APPLICATION","orderNo":30,"bindtap":"toPage","bindtapParams":"/pages/achievement_score_dot/achievement_score_dot","opentype":"","isDeleted":false,"createTime":"2023-09-05 18:27:45","updateTime":"2025-02-09 16:19:05"}],"success":true,"server":"web1"}`))
		} else if page == "OTHER" {
			w.Write([]byte(`{"status":200,"data":[{"id":5,"menu":"个人信息","menuName":"profile","menuColor":"#666","enable":true,"openAd":true,"page":"OTHER","orderNo":1,"bindtap":"toPage","bindtapParams":"/pages/profile/profile","opentype":"","isDeleted":false,"updateTime":"2023-09-02 12:44:23"},{"id":66,"menu":"登陆密码","menuName":"profile","menuColor":"#666","enable":true,"openAd":true,"page":"OTHER","orderNo":3,"bindtap":"toPage","bindtapParams":"/pages/profile/profile","opentype":"","isDeleted":false,"updateTime":"2023-09-02 12:45:07"},{"id":85,"menu":"个人设置","menuName":"setting","menuColor":"#666","enable":true,"openAd":false,"page":"OTHER","orderNo":11,"bindtap":"toPageNeedStudentId","bindtapParams":"/pages/my/setting/setting","opentype":"","isDeleted":false,"updateTime":"2023-09-02 12:45:07"},{"id":27,"menu":"生活小助手","menuName":"attendance","menuColor":"#666","enable":true,"openAd":false,"page":"OTHER","orderNo":13,"bindtap":"toOtherMP","bindtapParams":"wx9f1a2a23c489b4cf","opentype":"","isDeleted":false,"updateTime":"2023-09-02 12:44:23"}],"success":true,"server":"web1"}`))
		} else {
			handleDefault()(w, r)
		}
	}
}

func handleWeRun() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"data":{"weRunIndexRank":[]},"success":true,"server":"web1"}`))
	}
}

func handleNotice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"data":[],"success":true,"server":"web1"}`))
	}
}

func handleLoginOffline() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"message":"登录成功","data":{},"success":true,"server":"web1"}`))
	}
}

func handleSsoSetting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"data":{"jsLogin":true,"loginItems":[{"name":"统一身份","checked":"true","value":"SSO_LOGIN"}]},"success":true,"server":"web1"}`))
	}
}

func handleDefault() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":500,"message":"用不了了","success":true,"server":"web1"}`))
	}
}

func ExtractParams(pattern, path string) map[string]string {
	params := make(map[string]string)

	// 分割路径
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	// 检查部分数量是否一致
	if len(patternParts) != len(pathParts) {
		return nil
	}

	// 逐个部分匹配
	for i := 0; i < len(patternParts); i++ {
		patternPart := patternParts[i]
		pathPart := pathParts[i]

		// 检查是否是参数部分
		if strings.HasPrefix(patternPart, "{") && strings.HasSuffix(patternPart, "}") {
			// 提取参数名
			paramName := strings.TrimSuffix(strings.TrimPrefix(patternPart, "{"), "}")
			params[paramName] = pathPart
		} else if patternPart != pathPart {
			// 固定部分不匹配
			return nil
		}
	}

	return params
}
