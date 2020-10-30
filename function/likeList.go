package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/axgle/mahonia"
)

//获取所有关注的吧
func GetLikeList() []string {
	client := &http.Client{}
	var req *http.Request

	req, err := http.NewRequest("GET", "http://tieba.baidu.com/f/like/mylike?v=1603974580023", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//req.Header.Add("Accept-Encoding", "gbk")
	//req.Header.Add("Accept-Charset", "utf-8;q=0.7,*;q=0.3")
	req.Header.Add("Host", "tieba.baidu.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Proxy-Connection", "keep-alive")
	req.Header.Add("Accept", "text/html, */*; q=0.01")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")
	req.Header.Add("Referer", "http://tieba.baidu.com/i/i/forum")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Cookie", "***")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(resp)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	bodystr := string(body)
	//fmt.Println(bodystr)

	result := ConvertToString(bodystr, "gbk", "utf-8")
	//fmt.Println(result)

	//匹配吧名
	re := regexp.MustCompile("<td><a href=\"[^\"]*\"")
	list := re.FindAllString(result, -1)
	fmt.Println(list)
	for i, j := range list {
		list[i] = j[12:]
	}
	//fmt.Println(list)

	return list
}

//golang默认支持UTF-8格式
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
