package service

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

type resWbPy struct{
	Code string `json:"code"`
	Wb string `json:"wb"`
	Py string `json:"py"`
}

/*@Author:zengzeen
* @param:uipInterFunc struct
* @return:{code,msg,data}   json
 */
func GetWbPyService(data string) (resWbPy) {
	fmt.Println("this is FinTcDtl post!")
	var wbPy resWbPy
	wbPy.Py = LazyPinyin(data)
	wbPy.Wb = LazyCw(data)
	return wbPy
}

//拼音处理的相关函数
func LazyPinyin(s string) string {
	var lstr string
	pys := []string{}
	for _, v := range Pinyin(s) {
		pys = append(pys, v[0])
	}
	for _,row := range pys {
		lstr += row[:1]
	}
	return lstr
}

// Pinyin 汉字转拼音，支持多音字模式.
func Pinyin(s string) [][]string {
	pys := [][]string{}
	for _, r := range s {
		match, _ := regexp.MatchString("^[a-zA-Z0-9]", string(r))
		if match{
			ps :=[]string{string(r)}
			pys = append(pys, ps)
		}
		py := SinglePinyin(r)
		if len(py) > 0 {
			pys = append(pys, py)
		}
	}
	return pys
}

// SinglePinyin 把单个 `rune` 类型的汉字转换为拼音.
func SinglePinyin(r rune) []string {
	value, ok := PinyinDict[int(r)]
	pys := []string{}
	if ok {
		pys = strings.Split(value, ",")
	}
	for i,row := range pys  {
		pys[i] = row[:1]
	}
	return pys
}

//五笔编码相关的处理函数
func LazyCw(s string) string {
	//var pys []string
	var lstr string
	for _, v := range Wubi(s) {
		//pys = append(pys, v)
		//pys[i] = v[:1]
		lstr += v
	}
	return lstr
}

// Wubi 汉字转五笔，支持多音字模式.
func Wubi(s string) []string {
	var pys []string
	for _, r := range s {
		match, _ := regexp.MatchString("^[a-zA-Z0-9]", string(r))
		if match{
			pys = append(pys, string(r))
		}
		py := SingleWubi(string(r))
		if len(py) > 0 {
			pys = append(pys, py)
		}
	}
	for i,v := range pys {
		pys[i]=v[:1]
	}
	return pys
}

// SingleWubi 把单个 `rune` 类型的汉字转换为五笔编码.
func SingleWubi(r string) (s string) {
	uchar := strconv.QuoteToASCII(r)
	uchar = uchar[1 : len(uchar)-1]
	//fmt.Printf("the r here is %v, and the s unchar heer is %v:",r,uchar)
	value, ok := CW[string(uchar)]
	if ok {
		return value
	}
	return ""
}