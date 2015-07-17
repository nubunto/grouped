(function() {
    var App = {};
    var conn;
    var msg = $("#msg");
    var log = $("#log");
    var video = $("#video").get(0);

    function appendLog(msg) {
        var d = log[0]
        var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
        msg.appendTo(log)
        if (doScroll) {
            d.scrollTop = d.scrollHeight - d.clientHeight;
        }
    }
    App.appendLog = appendLog;

    $("#form").submit(function() {
        if (!conn) {
            return false;
        }
        if (!msg.val()) {
            return false;
        }
        conn.send(msg.val());
        msg.val("");
        return false
    });

    $('#pause').click(function() {
	    conn.send('@pause')
    })

    $('#unpause').click(function() {
	    conn.send('@unpause')
    })
    if (window["WebSocket"]) {
	App.start = function(host) {
			conn = new WebSocket('ws://' + host + '/ws');
	        conn.onclose = function(evt) {
	            appendLog($("<div><b>Connection closed.</b></div>"))
	        }
        	conn.onmessage = function(evt) {
		    console.log(evt)
		    if(evt.data === '@pause' || evt.data === '@unpause') {
			    if(evt.data === '@pause') {
				    video.pause()
			    } else if (evt.data === '@unpause') {
				    video.play()
			    }
		    } else {
	            	appendLog($("<div/>").text(evt.data))
		    }
	        }
	}
    } else {
	App.start = function() {
        	appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"))
	};
    }
    window.App = App;
})();
