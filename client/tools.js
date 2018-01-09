$(function(){
    var ws;
    if (window.WebSocket === undefined){
        alert("Sockets not supported")
        return;
    }
    ws = initWS();

    
    function initWS(){
        var socket = new WebSocket("ws://localhost:8000/dm_socket")
        socket.onopen = function(){
            alert("Socket Open");
        }
        socket.onmessage = function (e){
            alert(e)
        }
        socket.onclose = function(){
            alert("Socket closed");
        }
        return socket;

    }

    $("#sender").click(function (e){
        ws.send("Click");
    });
})