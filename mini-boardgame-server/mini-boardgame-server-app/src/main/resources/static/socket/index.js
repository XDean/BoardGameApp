new function () {
  var ws = null;
  var connected = false;

  var serverUrl;
  var connectionStatus;
  var topic;
  var addButton;
  var payload;
  var n = 0;

  var connectButton;
  var disconnectButton;
  var sendButton;

  var open = function () {
    var url = 'ws://'+window.location.host+'/socket/game/room/' + serverUrl.val();
    ws = new WebSocket(url);
    ws.onopen = onOpen;
    ws.onclose = onClose;
    ws.onmessage = onMessage;
    ws.onerror = onError;

    connectionStatus.text('OPENING ...');
    serverUrl.attr('disabled', 'disabled');
    connectButton.hide();
    disconnectButton.show();
  }

  var close = function () {
    if (ws) {
      console.log('CLOSING ...');
      ws.close();
    }
    connected = false;
    connectionStatus.text('CLOSED');

    serverUrl.removeAttr('disabled');
    connectButton.show();
    disconnectButton.hide();
    addButton.attr('disabled', 'disabled');
    payload.attr('disabled', 'disabled');
    topic.attr('disabled', 'disabled');
    sendButton.attr('disabled', 'disabled');
  }

  var clearLog = function () {
    $('#messages').html('');
  }

  var onOpen = function () {
    console.log('OPENED: ws://'+window.location.host+'/game/room/' + serverUrl.val());
    connected = true;
    connectionStatus.text('OPENED');
    payload.removeAttr('disabled');
    addButton.removeAttr('disabled');
    topic.removeAttr('disabled');
    sendButton.removeAttr('disabled');
  };

  var onClose = function () {
    console.log('CLOSED: ws://'+window.location.host+'/game/room/' + serverUrl.val());
    addMessage('CLOSED: ws://'+window.location.host+'/game/room/' + serverUrl.val());
    ws = null;
  };

  var onMessage = function (event) {
    var data = event.data;
    addMessage(data);
  };

  var onError = function (event) {
    alert(event.data);
  };

  var addMessage = function (data, type) {
    var msg = $('<pre>').text(data);
    if (type === 'SENT') {
      msg.addClass('sent');
    }
    var messages = $('#messages');
    messages.append(msg);

    var msgBox = messages.get(0);
    while (msgBox.childNodes.length > 1000) {
      msgBox.removeChild(msgBox.firstChild);
    }
    msgBox.scrollTop = msgBox.scrollHeight;
  }

  WebSocketClient = {
    init: function () {
      serverUrl = $('#serverUrl');
      connectionStatus = $('#connectionStatus');
      payload = $('#payload');
      topic = $('#topic');
      connectButton = $('#connectButton');
      disconnectButton = $('#disconnectButton');
      sendButton = $('#sendButton');
      addButton = $('#addButton');

      connectButton.click(function (e) {
        close();
        open();
      });

      disconnectButton.click(function (e) {
        close();
      });

      sendButton.click(function (e) {
        var attributesstr = '';
        for (var i=0;i<n;i++) {
          if ($('#attributekey'+i).length>0) {
            attributesstr+='"'+ $('#attributekey'+i).val() + '":"'+$('#attributevalue'+i).val() + '",';
          }
        }
        attributesstr=(attributesstr.slice(attributesstr.length-1)==',')?attributesstr.slice(0,-1):attributesstr;
        var msg = '{"topic":"' + $('#topic').val() + '","attributes":{' + attributesstr + '}, "payload":' + $('#payload').val() + '}';
        addMessage(msg, 'SENT');
        ws.send(msg);
      });

      addButton.click(function (e) {
        var attribute = document.createElement("div");
        attribute.id = "attribute" + n;
        attribute.innerHTML = "<input type='text' id='attributekey" + n + "' value=''/> = <input type='text' id='attributevalue" + n + "' value=''/><button class='removeButton' onclick='this.parentNode.remove()' id=" + n + ">remove</button>";

        document.getElementById("attributes").append(attribute);
        n++;

      });

      $('#clearMessage').click(function (e) {
        clearLog();
      });

      // var isCtrl;
      // sendMessage.keyup(function (e) {
      //   if (e.which == 17) isCtrl = false;
      // }).keydown(function (e) {
      //   if (e.which == 17) isCtrl = true;
      //   if (e.which == 13 && isCtrl == true) {
      //     sendButton.click();
      //     return false;
      //   }
      // });
    }
  };
}

$(function () {
  WebSocketClient.init();
});