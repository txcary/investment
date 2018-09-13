var secureJson;
require.config({
	baseUrl: 'js',
	paths: {
		securejson: '3rd/securejson',
		elliptic: "3rd/elliptic/elliptic.min",
		sha3: "3rd/sha3/sha3.min",
		aes: "3rd/aes/index.min",
		base64: "3rd/base64/base64js.min"
	},
});
require(['securejson'], function(sj){
	secureJson = sj;
});
var app = new Vue({
	el: '#app',
	created: function() {
		this.userName = getCookie("portfolio","userName");
		this.userPasswd = getCookie("portfolio","userPasswd");
		if(this.userName!="" && this.userPassed!="") {
			this.isLogined=true;
		}else{
			this.logout();
		}
	},
	methods: {
		logout: function() {
			this.userName="";
			this.userPasswd="";
			this.isLogined=false;
			delCookie("portfolio", "userName");
			delCookie("portfolio", "userPasswd");
		},
		login: function() {
			userName = document.getElementById('login.user').value;
			if(userName===undefined || userName==="") {
				alert("User Name not valid!");
				return;
			}
			this.userName = userName;
			userPasswd = document.getElementById('login.passwd').value;
			if(userPasswd===undefined || userPasswd==="") {
				alert("User Password not valid!");
				return;
			}
			this.userPasswd = userPasswd;
			this.isLogined=true;
			setCookie("portfolio", "userName", this.userName);
			setCookie("portfolio", "userPasswd", this.userPasswd);
		},
		getPortfolioFromServer() {
			var str = secureJson.GenerateJson(this.userName, this.userPasswd, "");
			var xhr = new XMLHttpRequest();
			xhr.open("POST", "/portfolio/getjson", true);
			xhr.setRequestHeader('content-type', 'application/json');
			xhr.onreadystatechange = function() {
				if(xhr.readyState== XMLHttpRequest.DONE){
					alert(xhr.responseText);
				}
			}
			xhr.send(str);
		},
		putPortfolioToServer() {
			var str = secureJson.GenerateJson(this.userName, this.userPasswd, "MyData");
			var xhr = new XMLHttpRequest();
			xhr.open("POST", "/portfolio/putjson", true);
			xhr.setRequestHeader('content-type', 'application/json');
			xhr.onreadystatechange = function() {
				if(xhr.readyState== XMLHttpRequest.DONE){
					alert(xhr.responseText);
				}
			}
			xhr.send(str);
		},
		keyToName: function(key) {
			if(key===undefined){
				return "";
			}
			if(this.nameOfKeys[key]===undefined) {
				return key;
			}else{
				return this.nameOfKeys[key];
			}
		},
		setCommand: function(cmd, id) {
			this.command.pendingId = undefined;
			this.command.pendingAmount = undefined;
			this.command.pendingPrice = undefined;

			switch(cmd) {
			case "buy":
				this.command.pendingCommand = "Buy";
				this.command.cmdNeedId = true;
				this.command.cmdNeedAmount = true;
				this.command.cmdNeedPrice = true;
				this.command.cmdNeedSubmit = true;
				break;
			case "sell":
				this.command.pendingCommand = "Sell";
				this.command.cmdNeedId = true;
				this.command.cmdNeedAmount = true;
				this.command.cmdNeedPrice = true;
				this.command.cmdNeedSubmit = true;
				break;
			case "add":
				this.command.pendingCommand = "Add";
				this.command.cmdNeedId = false;
				this.command.cmdNeedAmount = true;
				this.command.cmdNeedPrice = false;
				this.command.cmdNeedSubmit = true;
				break;
			case "score":
				this.command.pendingCommand = "Score";
				this.command.cmdNeedId = true;
				this.command.cmdNeedAmount = true;
				this.command.cmdNeedPrice = false;
				this.command.cmdNeedSubmit = true;
				break;
			}
			if(id!=undefined) {
				this.command.pendingId = id;
			}
		},
	},
	data: {
		userName: "",
		userPasswd: "",
		isLogined: false,
		nameOfKeys: {
			refrate: "贴现率",
			hkExchangeRate: "港币汇率",
		},
		statistics: {
			refrate: 4.9,
			hkExchangeRate: 0.8432,
			hs300: 3225,
			stockCapitalization: 20000,
			hkCapitalization: 20000,
			cash: 20000,
			totalCapitalization: 60000,
			estimatedCash: 20000,
			totalProfitRate: 100,
			yearProfitRate: -10,
		},
		command: {
			pendingCommand: "",
			pendingId: "",
			pendingAmount: 0,
			pendingPrice: 0,
			cmdNeedId: false,
			cmdNeedAmount: false,
			cmdNeedPrice: false,
			cmdNeedSubmit: false,
		},
		portfolio: [
			{mark:"+", id: "600519", name: "贵州茅台", price: 646.0, share: 200, trade: 0, capitalization: 129200, level: 9.2, lowlevel: 11.1, uplevel:12.1, marketlevel: 14},
		],
	},
});


function setCookie(space,name,value,exdays){
    var d = new Date();
    var name = space+"_"+name;
    if(exdays===undefined) {
    	exdays = 1;
    }
    d.setTime(d.getTime()+(exdays*24*60*60*1000));
    var expires = "expires="+d.toGMTString();
    document.cookie = name+"="+value+"; "+expires;
}

function getCookie(space,name){
    var name = space+"_"+name + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i].trim();
        if (c.indexOf(name)==0) { return c.substring(name.length,c.length); }
    }
    return "";
}

function delCookie(space,name){
	setCookie(space,name,"",-1);
}
