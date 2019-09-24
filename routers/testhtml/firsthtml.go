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
<adress>
By <a href="2336260845@qq.com">马永真 </a>
</adress>
</br>
<button type="button" onclick="print()">上传文件</button>
<script>
function print() {
	alert("你点击了按钮")
}
</script>
</br>
<adress>
<a href="file">上传文件 </a>
</adress>

<!注释是不会被显示的，哈哈哈哈>`
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, s)
	return
}

func Writer(ctx *gin.Context) {
	s := `<h>这是我的邮件:2336260845@qq.com</h>
<p></p>`

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, s)
	return
}

func File(ctx *gin.Context) {
	s := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>上传图片</title>
    <script type="text/javascript">
        var xhr
        function UpFile() {
            var file = document.getElementById("file").files[0];
            var url = "http://127.0.0.1:3333" + "/fileop/push";

            var form = new FormData;
            form.append("file", file);

            xhr = new XMLHttpRequest();
            xhr.open("post", url, true);
            xhr.onload = UpComplete;
            xhr.onerror = UpFailed;

            //TODO 进度条

            xhr.send(form);
        }

        function UpComplete(evt) {
            var data = JSON.parse(evt.target.responseText);
            if (data.success) {
                alert("上传成功");
            } else {
                alert("上传失败");
            }
        }

        function UpFailed(evt) {
            alert("上传失败");
        }

        function CancleUpFile() {
            xhr.abort();
        }
    </script>
</head>
<body>
<progress id="progressBar" value="0" max="100" style="width: 300px;"></progress>
<span id="percentage"></span><span id="time"></span>
<br /><br />
<input type="file" id="file" name="myFile" />
<input type="button" onclick="UpFile()" value="上传文件" />
<input type="button" onclick="CancleUpFile()" value="取消" />
</body>
</html>
`

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, s)
	return
}
