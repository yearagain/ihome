/**
 * 
 */

var bglogin="http://127.0.0.1:8888/bglogin";
var signin=new Vue({
    el:"#signin",
    data:{
        pwdregular:/[0-9A-Za-z]{8,16}$/, //密码正则表达式
        pwdrule:"口令应由8-16位字母数字组成！",    //密码规则叙述
        messtyle:"",    //提示框样式
        message:"", //提示框内容
        password:"",    //输入框内容
        pwdsalt:"yoursalt", //加密盐
        subData:{}  //提交数据
    },
    methods:{
        change:function(){
            if(this.pwdregular.exec(this.password)){
                this.messtyle="alert alert-primary";
                this.message="口令格式正确";
            }else{
                this.messtyle="alert alert-warning";
                this.message=this.pwdrule;
            }
            
        },
        submit:function(){
            if(!this.pwdregular.exec(this.password)){
                this.messtyle="alert alert-danger";
                this.message="口令格式错误！"+this.pwdrule;
                return
            }
            //加密提交
            this.subData.password=md5(this.password+this.pwdsalt);
            
            $.ajax({
                //请求方式
                type : "POST",
                //请求地址
                url : bglogin,
                //向请求地址发送的数据
                data : this.subData,
                //请求成功,result为返回的数据
                success : function(result) {
                    signin.$data.messtyle="alert alert-success";
                    signin.$data.message=result.msg;
                    localStorage.setItem("token", result.data);
                    //延迟跳转
                    setTimeout(function(){},"1000")
                    window.location.href="/static/view/background/bg.html";
                    
                },
                //请求失败，err为返回的错误信息
                error : function(err){
                    signin.$data.messtyle="alert alert-danger";
                    signin.$data.message=err.responseJSON.msg;
                }
            });
        }
    }
})
