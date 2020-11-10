/**
 * 鉴权相关方法
 * @param {*} obj 
 */
$.ajax({
    type:"POST",
    url:"http://127.0.0.1:8888/test",
    data:queryParamsObj(),
    success : function(result) {
        
    },
    error : function(err){
        window.location.href="/static/view/background/login.html";
    }

  })
function queryParamsObj(obj) {
    // var username = localStorage.getItem("username");
    // var session= localStorage.getItem("session");
    // return "username=" + username + "&token=" + token;
    if(obj==null){
        var obj={}
        
        obj.token=localStorage.getItem("token");
        return obj;
    }
    obj.token=localStorage.getItem("token");
    return obj;
}