package testhtml

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	s := `<!DOCTYPE html>
<html>
<body>

<h1>我的第一张网页</h1>

<p>我的第一个段落</p>

<p id="demo"></p>

<script>
 document.getElementById("demo").innerHTML = 5 + 6;
</script>

</body>
</html>`
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
