# GoLang语言的应用 之 Demo_Xpath_Go
利用 Xpath 读取 html 内容, 通过 Github 来演示 Xpath 如何使用 <BR/>
<BR/> 
# 更新信息
开发者：沙振宇（沙师弟专栏） <BR/>
创建时间：2019-12-4<BR/>
最后一次更新时间：2019-12-5<BR/>
CSDN博客地址：https://shazhenyu.blog.csdn.net/article/details/103391918 <BR/>
<BR/>
# 先安装 Github 的其它包
https://github.com/pkg/errors <BR/>
https://github.com/lestrrat-go/libxml2 <BR/>
<BR/>
# 去除内容的空格和换行符
str = strings.Replace(str, " ", "", -1) // 去除空格 <BR/> 
str = strings.Replace(str, "\n", "", -1) // 去除换行符<BR/> 
<BR/>
# 运行效果
nodes type: types.NodeList,len: 3 <BR/>
<BR/>
nodes type: *dom.Text,text: WhyGitHub? <BR/>
nodes type: *dom.Text,text: Explore <BR/>
nodes type: *dom.Text,text: Pricing <BR/>

