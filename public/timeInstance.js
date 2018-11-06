function checkTime(i){
    if (i<10){
        i="0" + i;
    }
    return i;
}

function startTime(){

    var today=new Date();
    var h=today.getHours();
    var m=today.getMinutes();
    var s=today.getSeconds();// 在小于10的数字前加一个‘0’
    m=checkTime(m);
    s=checkTime(s);
    document.getElementById("txt").innerHTML= h+":"+m+":"+s;
    t=setTimeout(function(){startTime()},500);
}

function ajax_Proc(){
    // $("#what").html("我是一个傻子吧");
    $.get("/admin/data",function(data,status){
        let obj = JSON.parse(data);
        var cols = ['Equipment'];

        for(var item in obj){
            // console.log('item',item);
            for (var col in obj[item]){
                // console.log('col',col);
                if (!(col in cols)){
                    // console.log('cols',cols);
                    cols.push(col)
                    // console.log('cols',cols);
                }
            }
        }

        var dt = [];
        for(var item in obj){
            var tmp = [item];
            var tmp_arr = cols.slice(1)
            for(var col in tmp_arr){

                // console.log('tmp',tmp,'val',obj[item][tmp_arr[col]]);
                tmp.push(obj[item][tmp_arr[col]]);
            }
            dt.push(tmp);
        }

        Table().init({
            id:'table',
            header:cols,
            data:dt
        });

        //获取table序号
        var x = $("#table");
        x.find("tr").each(function(){
            $(this).find("td:eq(1)").each(function(){
                // console.log($(this).text());
                var h_val = $(this).text();
                if (h_val>90 && h_val<=120){
                    $(this).css("background-color","#FF7575");
                }else if(h_val>80 && h_val<=90){
                    $(this).css("background-color","#F0FFF0");
                }else if (h_val<=80 && h_val>10){
                    $(this).css("background-color","#FFFFCE");
                }else {
                    $(this).css("background-color","#8E8E8E");
                }
            });
        });
    }
);
}




