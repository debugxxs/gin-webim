var uname = prompt('给自己起个响亮的名字吧');
const uuid = createUUID(10)
if(uname) {
    uname = uname.trim()
    var ws = new WebSocket("ws://localhost:8008/im");
    system("正在连接服务器...")

    ws.onopen = function() {
        ws.send(JSON.stringify({
            "type": "login",
            "uuid": uuid,
            "content": "Hello Go WebSocket",
            "username": uname
        }));
    };

    ws.onmessage = function(evt) {//绑定收到消息事件
        console.log( "Received Message: " + evt.data);
        const data = JSON.parse(evt.data)
        switch (data.type) {
            case "init":
                system("服务器连接成功", "success")
                break;
            case "login":
                userlistDom(data.user_list)
                system(`${data.username} 进入群聊`, "success")
                break;
            case "message":
                var { username, content } = data
                const message = `<span style="color: #40a9ff;">${username}</span>: ${content}`
                acceptMessage(message)
                break;
            case "private":
                var { username, content, touuid} = data
                if(touuid === uuid) {
                    const message = `<span style="color: red;">${username}</span> 对你说: ${content}`
                    acceptMessage(message)
                }
                break;
            case "logout":
                userlistDom(data.user_list)
                system(`${data.username} 已下线`, "error")
                break;
        }
    };

    ws.onclose = function() { //绑定关闭或断开连接事件
        system("与服务器连接断开", "error")
        ws.send(JSON.stringify({
            "type": "logout",
            "uuid": uuid,
            "content": "下线",
            "username": uname
        }))
    };
} else {
    system("服务器未连接，请给自己起个名字吧～，<a href=''>点我起名</a>")
}

document.onkeydown = function (event) {
    var e = event || window.event;
    if (e && e.keyCode === 13) {
        e.preventDefault()
        send()
    }
};

// 发送消息
function send() {
    const message = document.getElementById("content-value").value
    if(message.trim().toString().length <= 0) {
        alert("请输入发送的内容")
        return
    }
    const UUID = document.getElementById("UUID").value;
    const type = UUID ? "private" : "message"
    if(type === "private") {
        const info = `${uname}: ${message}`
        acceptMessage(info)
    }
    ws.send(JSON.stringify({
        "type": type,
        "content": message.trim(),
        "username": uname,
        "touuid": UUID
    }));
    document.getElementById("UUID").value = ""
    document.getElementById("content-value").value = ""
}

// 文本消息
function acceptMessage(message) {
    document.getElementById("content").innerHTML += `
                <div style="line-height: 30px;">${message}</div>
            `;
    setTimeout(() => {
        document.getElementById("content").scrollTo(0, document.getElementById("content").offsetHeight);
    }, 1000)
}

function system(message, type = "loading") {
    document.getElementById("content").innerHTML += `
                <div class="system ${type}">系统消息：${message}</div>
            `
}

function userlistDom(userList) {
    document.getElementById("user-list").innerHTML = ""
    userList.map(item => {
        console.log(uuid, item.uuid)
        if(uuid === item.uuid) {
            document.getElementById("user-list").innerHTML += `<li style="color: red;">${item.username}(我)</li>`
        } else {
            document.getElementById("user-list").innerHTML += `
                        <li onclick="privateMessage('${item.username}', '${item.uuid}')">${item.username}</li>
                    `
        }
    })
}

function privateMessage(user, uuid) {
    document.getElementById("UUID").value = uuid
    document.getElementById("content-value").value = `@${user} `
}

function createUUID(len, radix = null) {
    var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
    var uuid = [], i;
    radix = radix || chars.length;
    if (len) {
        for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random() * radix];
    } else {
        var r;
        uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
        uuid[14] = '4';
        for (i = 0; i < 36; i++) {
            if (!uuid[i]) {
                r = 0 | Math.random() * 16;
                uuid[i] = chars[(i === 19) ? (r & 0x3) | 0x8 : r];
            }
        }
    }
    return uuid.join('');
}