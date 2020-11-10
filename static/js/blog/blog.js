var getbloglist="http://127.0.0.1:8888/getbloglist";
var info = new Vue({
    el: '#info',
    data:{
      slogan1:"String1",
      slogan2:"String2",
      portrait:"/static/img/portrait.jpg"
    }
})
var list=new Vue({
    el: '#list',
    data:{
        bloglist:[]
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
                href:"/static/view/blog/blog_show.html?id="+result.data[i].ID
                })
        }
        
        
    },
    //请求失败，err为返回的错误信息
    error : function(err){
        alert(err.responseJSON.msg)
    }

});
// list.$data.bloglist=[
// {title:"1",abstract:"1" },
// {title:"2",abstract:"2" }
// ]