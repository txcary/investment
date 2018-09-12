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
	methods: {
		logout: function() {
			this.userName="";
			this.userPasswd="";
			this.isLogined=false;
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
			var str = secureJson.GenerateJson(this.userName, this.userPasswd, "");
			alert(str);
			this.isLogined=true;
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
