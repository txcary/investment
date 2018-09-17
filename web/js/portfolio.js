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
		getStock: function(id,callback) {
			var xhr = new XMLHttpRequest();
			xhr.open("GET", "/stock/"+id, true);
			xhr.onreadystatechange = function() {
				if(xhr.readyState== XMLHttpRequest.DONE){
					callback(xhr.responseText);
				}
			}
			xhr.send();
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
		getPortfolioIdx: function(id) {
			for(var i=0;i<this.portfolio.length;i++) {
				if(this.portfolio[i].id==id) {
					return i;
				}
			}
			return undefined;
		},
		doTrade: function(id, amount, price) {
			var self = this;
			if(id===undefined) {
				alert("ID not valid!");
				return;
			}
			if(price<0) {
				alert("Price not valid!");
				return;
			}
			var idx = this.getPortfolioIdx(id);
			if(idx===undefined) {
				if(amount<0) {
					alert(id+" not exist in portfolio!");
					return;
				}
				this.getStock(id, function(resp){
					if(resp===undefined) {
						alert("Can not get data of id "+id);
						return;
					}
					var stock = JSON.parse(resp);
					var item = {
						id: id,
						name: stock.Name,
						share: 0,
						price: 0,
					};
					self.portfolio.push(item);
					self.doTrade(id, amount, price);
				});
				return;
			}
			var tradedShare = this.portfolio[idx].share + parseInt(amount);
			if(tradedShare <0 ) {
				alert("No enough share!");
				return;
			}
			if(tradedShare==0) {
				this.portfolio.splice(idx,1);
				return;
			}
			this.portfolio[idx].share = tradedShare;
			this.portfolio[idx].price = parseFloat(price);
		},
		doScore: function(id, amount) {
			var idx = this.getPortfolioIdx(id);
			if(idx===undefined) {
				alert("ID not valid!");
				return;
			}
			this.portfolio[idx].score = amount;
		},
		commandSubmit: function() {
			if(this.command.pendingCommand===undefined) {
				return;
			}
			var id = this.command.pendingId;
			var amount = this.command.pendingAmount;
			var price = this.command.pendingPrice;
			switch(this.command.pendingCommand) {
			case "Sell":
				return this.doTrade(id, -amount, price);
			case "Buy":
				return this.doTrade(id, amount, price);
			case "Add":
				//return this.doAdd(obj.pendingId, obj.pendingAmount);
			case "Score":
				return this.doScore(id, amount);
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
		statisticsView: {
			editing: false,
			shownTagsAsEditable: [
				{key:'cash', displayAs:'现金'}, 
			],
			hidenTagsAsEditable: [
				{key: 'virtualShare', displayAs:'组合份额'}, 
				{key: 'virtualValue', displayAs:'组合净值'}, 
				{key: 'virtualValueLYR', displayAs:'组合去年净值'},
			],
			shownTagsAsReadonly: [
				{key:'stockCapitalization', displayAs:'A股市值'}, 
				{key:'hkCapitalization', displayAs:'港股市值'}, 
				{key:'totalCapitalization', displayAs:'总市值'}, 
			],
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
/*
 * securejson = {
 * 	statistics: {cash,virtualShare,virtualValue,virtualValueLYR},
 * 	portfolio: [
 * 		{id,share,score,unit},
 * 		{...},
 * 	]
 * }
 */		
		statistics: {
			//Configured
			cash: 20000,
			virtualShare: 100,
			virtualValue: 2.1410,
			virtualValueLYR: 5,
			//fetched
			refrate: 4.9,
			hkExchangeRate: 0.8432,
			hs300: 3225,
			hs300LYR: 4030,
			stockTrend: 100,
			hkTrend: 100,
			//Computed
			stockCapitalization: 20000,
			hkCapitalization: 20000,
			totalCapitalization: 60000,
			estimatedCash: 20000,
			totalProfitRate: 100,
			yearProfitRate: -10,
		},
		portfolio: [
			{
				//configured
				id: "600519", 
				share: 200, 
				score: 100,
				unit: 100,
				discount: 1.00,
				//fetched
				name: "贵州茅台", 
				price: 646.0, 
				//computed
				mark:"+", 
				trade: 0, 
				capitalization: 129200, 
				level: 9.2, 
				lowlevel: 11.1, 
				uplevel:12.1, 
				marketlevel: 14,
			},
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
