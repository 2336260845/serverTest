package testhtml

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	s := `
<h1>所有功能</h1>
<a href="file">智能切割图片 </a>
</adress>
</br>
<a href="email">延时发送邮件 </a>
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
	s := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>上传图片</title>
    <script type="text/javascript">
        var xhr

        localPath = "http://127.0.0.1:3333"

        function UpFile() {
            var fileop = document.getElementById("file").files[0];
            var url = localPath + "/fileop/push";

            var form = new FormData;

            form.append("file", fileop);

            xhr = new XMLHttpRequest();
            xhr.open("post", url, true);
            xhr.onload = UpComplete;
            xhr.onerror = UpFailed;

            xhr.upload.onprogress = progressFunction;//【上传进度调用方法实现】
            xhr.upload.onloadstart = function(){//上传开始执行方法
                ot = new Date().getTime();   //设置上传开始时间
                oloaded = 0;//设置上传开始时，以上传的文件大小为0
            };

            xhr.send(form);
        }

        function cropImageAndDownload() {
            var fileop = document.getElementById("file").files[0];
            var url = localPath + "/upload/images/cropfiles/" + fileop.name;

            var form = new FormData;
            form.append("file", fileop);
            var pullUrl = localPath + "/fileop/pull";

            xhr = new XMLHttpRequest();
            xhr.open("post", pullUrl, true);
            xhr.onload = downComplete;
            xhr.onerror = downFailed;

            xhr.send(form);

            document.getElementById("imgesShow").src = url;
        }

        function UpComplete(evt) {
            var data = JSON.parse(evt.target.responseText);
            if (data.ok == "success") {
                alert("上传成功");
            } else {
                alert("上传失败");
            }
        }

        function UpFailed(evt) {
            alert("上传失败");
        }

        function downComplete() {
            var data = JSON.parse(evt.target.responseText);
            if (data.ok == "success") {
                alert("下载成功");
            } else {
                alert("下载失败");
            }
        }

        function downFailed(evt) {
            alert("下载失败");
        }

        function CancleUpFile() {
            xhr.abort();
        }

        function progressFunction(evt) {
            var progressBar = document.getElementById("progressBar");
            var percentageDiv = document.getElementById("percentage");
            // event.total是需要传输的总字节，event.loaded是已经传输的字节。如果event.lengthComputable不为真，则event.total等于0
            if (evt.lengthComputable) {//
                progressBar.max = evt.total;
                progressBar.value = evt.loaded;
                percentageDiv.innerHTML = Math.round(evt.loaded / evt.total * 100) + "%";
            }
            var time = document.getElementById("time");
            var nt = new Date().getTime();//获取当前时间
            var pertime = (nt-ot)/1000; //计算出上次调用该方法时到现在的时间差，单位为s
            ot = new Date().getTime(); //重新赋值时间，用于下次计算
            var perload = evt.loaded - oloaded; //计算该分段上传的文件大小，单位b
            oloaded = evt.loaded;//重新赋值已上传文件大小，用以下次计算
            //上传速度计算
            var speed = perload/pertime;//单位b/s
            var bspeed = speed;
            var units = 'b/s';//单位名称
            if(speed/1024>1){
                speed = speed/1024;
                units = 'k/s';
            }
            if(speed/1024>1){
                speed = speed/1024;
                units = 'M/s';
            }
            speed = speed.toFixed(1);
            //剩余时间
            var resttime = ((evt.total-evt.loaded)/bspeed).toFixed(1);
            time.innerHTML = '，速度：'+speed+units+'，剩余时间：'+resttime+'s';
            if(bspeed==0) time.innerHTML = '上传已取消';
        }
    </script>
</head>
<body>
<progress id="progressBar" value="0" max="100" style="width: 300px;"></progress>
<span id="percentage"></span><span id="time"></span>
<br /><br />
<input type="file" id="file" name="myFile" />
<input type="button" onclick="UpFile()" value="上传文件" />
<input type="button" onclick="cropImageAndDownload()" value="智能切割图片并下载" />
<input type="button" onclick="CancleUpFile()" value="取消" />
</br>
<img src="" id="imgesShow" align=left/>
</body>
</html>

`

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, s)
	return
}
