var getbloglist="http://127.0.0.1:8888/getbloglist";
var delblog="http://127.0.0.1:8888/delblog"
/**
 * 
 */
var bgsite="/bgsite"
/**
 * nav
 */
var info=new Vue({
    el:"#nav",
    
})
/**
 * 后台主页
 */
var info=new Vue({
    el:"#info",
    data:{
        username:"站长",
    
    }
})
// /**
//  * 后台列表
//  */
var list=new Vue({
    el: '#list',
    data:{
        bloglist:[]
    },
    methods: {
        delblog:function(item){
            $.ajax({
                type:"POST",
                url:delblog,
                data:{delblogid:item.id},
                success:function(res){
                    item.show=false;
                    alert(res.msg);
                },
                error : function(err){
                    alert(err.responseJSON.msg)
                    localStorage.setItem("token", null);
                    window.location.href="/static/view/background/login.html";
                }
            })
        }
    }
})
$.ajax({
    //请求方式
    type : "POST",
    //请求地址
    url : getbloglist,
    //向请求地址发送的数据
    data : null,
    //请求成功,result为返回的数据
    success : function(result) {
        for(var i in result.data){
            // console.log(result.data[i])
            list.$data.bloglist.push({
                title:result.data[i].Title,
                abstract:result.data[i].CreatedAt,
                id:result.data[i].ID,
                show:true
                })
        }
        
        
    },
    //请求失败，err为返回的错误信息
    error : function(err){
        alert(err.responseJSON.msg)
    }

});