package testhtml

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	s := `<h1> This is a head </h1>
<p> This is a </br>paragraph </p>
<a href="http://www.baidu.com" style="text-decoration:none"> 前往百度 

</a>
<!--
<pre>	for i := 0; i < 5; i++ {
		i += 1
	}</pre>
-->
</br>
<adress>By <a href="2336260845@qq.com">马永真</adress>
<!注释是不会被显示的，哈哈哈哈>`
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, s)
	return
}

func Writer(ctx *gin.Context) {
	s := `<h>这是我的邮件</h>
<p>如果遇到了问题可以联系我</p>`

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, s)
	return
}

func Hello2(ctx *gin.Context) {
	s := `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>我的第一个html页面</title>
         
        <style type="text/css">
            .btn {
                height: 30px;
                width: 100px;
                background: #0c093e;
                color:#fff;
                font-family: "微软雅黑";
                border-radius: 5px;
                text-align: center;
                line-height: 30px;
                position: absolute;
                left: 50%;
                margin-left: -50px;
                top: 50%;
                margin-top: -15px;
                cursor: pointer;
                transition: background 1s ease;
            }
             
            .btn:hover{
                background: #251f8c;
            }
        </style>
    </head>
     
    <script type="text/javascript">
     
        /*var person = {
            name : "zhangsan"
        }
     
        person.name = "lisi";
         
        person.name = "wangwu"
         
        alert(person.name);*/
     
     
        window['onload'] = function(){
             
            //alert(1);
            //1. 获取按钮的dom元素
            var btn = document.getElementsByTagName('div')[0];
             
            btn.onclick = function(){
                alert('保存成功！');
            }
             
        }
         
        /*window.onload = function(){
            alert(2);
        }*/
         
        /*
         *  A 1
         *  B 2
         *  C 1,2
         *  D 报错
         * 
         * */
    </script>
     
    <body>
        <div class='btn'>
            保存
        </div>
    </body>
</html>`
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, s)
	return
}
