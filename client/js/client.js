var client_app = function(){
	var webSocket ;
	
	var initWcConn = function(){
	 	webSocket =  new WebSocket('ws://127.0.0.1:5498/ws');  
		webSocket.onerror = function(event) {  
		    onError(event)  
		};  

		webSocket.onopen = function(event) {  
		    onOpen(event)  
		};  

		webSocket.onmessage = function(event) {  
		    onMessage(event)  
		};  
	}
    
    function onMessage(event) {  

       var data = JSON.parse(event.data);
       var msg = "";
       if(data.Event == "self"){
          msg = ' <div class="mytalk"> \
                            <dl> \
                                <dt><img src="images/tx01.png" width="100%" /></dt>\
                                <dd><span>'+data.Body+'</span></dd> \
                            </dl> \
                        </div> \
                ';
       }else{
          msg = ' <div class="chatwith"> \
                            <dl> \
                                <dt><img src="images/tx02.png" width="100%" /></dt>\
                                <dd><em>客服欢欢</em></dd> \
                                <dd><span>'+data.Body+'</span></dd> \
                            </dl> \
                        </div> \
                ';

       }
      
        $('#jsb_show_message_box').append(msg);
   		$('#jsb_chatTop').scrollTop( $('#jsb_chatTop')[0].scrollHeight );
    }  

    function onOpen(event) {  
        // document.getElementById('messages').innerHTML = '连接成功';  
        console.info("连接成功");
    }  

    function onError(event) {  
        alert(event.data);  
        alert("error");  
    }  

    function send() {  
            var send_text = $('#jsb_message_box').val();
             $('#jsb_message_box').val("");
            if (send_text.trim().length <= 0){
                alert("请输入要发送的内容");
                return false ;
            }
            
            var datas = {};

            datas['Route'] = "ONE_TO_MANY";
            datas['Context'] = {};
            datas['Context']['Body'] = send_text.replace(/\n/g,"<BR>") ;
            datas['Context']['Event'] = "send"
            webSocket.send('ayou_heiyo\n'+ JSON.stringify(datas)+'\n');  
            return false;  
        }  
        
        document.onkeydown=function(event){
            var e = event || window.event || arguments.callee.caller.arguments[0];
               
             if(e && e.keyCode==13){ // enter 键
                return send()
            }

        }; 


        $('#jsb_send_btn').click(function(){
        	send();
        })
	return {
		init:function(){
			initWcConn()
		}
	}
}();