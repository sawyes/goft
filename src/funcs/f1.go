package funcs

import "html/template"

//WebUtils.GenTplFunc 辅助生产funcmap文件


func Strong(txt string) template.HTML {
	return template.HTML("<strong style='color:red'>" + txt + "</strong>")
}

func Box(txt string) template.HTML {
	return template.HTML("<div class='.box'>" + txt + "</div>")
}

func Test() string {
	return "aaa"
}
