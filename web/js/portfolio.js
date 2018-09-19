import * as vtable from "./components/table.js";
import * as veditable from "./components/editable.js";
import * as vform from "./components/form.js";
import * as utils from "./utils.js";

var app = new Vue({
	el: '#app',
	components: {
		'vtable': vtable.component,
		'veditable': veditable.component,
		'vform': vform.component,
	},
	created: function() {
		this.userName = utils.getCookie("portfolio","userName");
		this.userPasswd = utils.getCookie("portfolio","userPasswd");
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
			utils.delCookie("portfolio", "userName");
			utils.delCookie("portfolio", "userPasswd");
		},
		login: function() {
			var userName = document.getElementById('login.user').value;
			if(userName===undefined || userName==="") {
				alert("User Name not valid!");
				return;
			}
			this.userName = userName;
			var userPasswd = document.getElementById('login.passwd').value;
			if(userPasswd===undefined || userPasswd==="") {
				alert("User Password not valid!");
				return;
			}
			this.userPasswd = userPasswd;
			this.isLogined=true;
			utils.setCookie("portfolio", "userName", this.userName);
			utils.setCookie("portfolio", "userPasswd", this.userPasswd);
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
			for(var i=0;i<this.componentData.tableData.datas.length;i++) {
				if(this.componentData.tableData.datas[i].id==id) {
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
					self.componentData.tableData.datas.push(item);
					self.doTrade(id, amount, price);
				});
				return;
			}
			var tradedShare = this.componentData.tableData.datas[idx].share + parseInt(amount);
			if(tradedShare <0 ) {
				alert("No enough share!");
				return;
			}
			if(tradedShare==0) {
				this.componentData.tableData.datas.splice(idx,1);
				return;
			}
			this.componentData.tableData.datas[idx].share = tradedShare;
			this.componentData.tableData.datas[idx].price = parseFloat(price);
		},
		doScore: function(id, amount) {
			var idx = this.getPortfolioIdx(id);
			if(idx===undefined) {
				alert("ID not valid!");
				return;
			}
			this.componentData.tableData.datas[idx].score = amount;
		},
		doAdd: function(amount) {
			var amountFloat = parseFloat(amount);
			this.componentData.statisticsData.data.cash += amountFloat;
			this.componentData.statisticsData.data.virtualShare += amountFloat/this.componentData.statisticsData.data.virtualValue;
		},
		commandSubmit: function(cmd, event) {
			var formValue = Object();
			for(var idx in event) {
				formValue[event[idx].Key] = event[idx].Value;
			}
			switch(cmd) {
			case "sell":
				return this.doTrade(formValue['id'], -formValue['amount'], formValue['price']);
			case "buy":
				return this.doTrade(formValue['id'], formValue['amount'], formValue['price']);
			case "add":
				return this.doAdd(formValue['amount']);
			case "score":
				return this.doScore(formValue['id'], formValue['amount']);
			}
		},
		setCommand: function(cmd, id) {
			var tab = $('[data-toggle=tab]');
			for(var idx=0;idx<tab.length; idx++){
				if(tab[idx].href.match("#tab"+cmd)=="#tab"+cmd) {
					tab[idx].classList.add("active");
					tab[idx].classList.add("show");
				} else {
					tab[idx].classList.remove("active");
					tab[idx].classList.remove("show");
				}
			}
			var pane = $('[data-toggle=pane]');
			for(var idx=0; idx<pane.length; idx++){
				pane[idx].classList.remove("active");
				pane[idx].classList.remove("show");
			}
			document.getElementById("tab"+cmd).classList.add("active");
			document.getElementById("tab"+cmd).classList.add("show");
			for(var idx in this.componentData.commandData[cmd].datas) {
				if(this.componentData.commandData[cmd].datas[idx].Key=="id") {
					this.componentData.commandData[cmd].datas[idx].Value = id;
				}
			}
		},
	},
	data: {
		userName: "",
		userPasswd: "",
		isLogined: false,
		componentData: {
			commandData: {
				buy: {
					name: "买进",
					datas: [
						{Key: "id",	Name: "代码",	Value: ""},
						{Key: "amount",	Name: "买进数量",	Value: ""},
						{Key: "price",	Name: "买进价格",	Value: ""},
					],
				},
				sell: {
					name: "卖出",
					datas: [
						{Key: "id",	Name: "代码",	Value: ""},
						{Key: "amount",	Name: "卖出数量",	Value: ""},
						{Key: "price",	Name: "卖出价格",	Value: ""},
					],
				},
				add: {
					name: "入金",
					datas: [
						{Key: "amount",	Name: "金额",	Value: ""},
					],
				},
				score: {
					name: "评分",
					datas: [
						{Key: "id",	Name: "代码",	Value: ""},
						{Key: "amount",	Name: "分数",	Value: ""},
					],
				},
			},
			statisticsData: {
				editing: false,
				shownTagsAsEditable: [
					{key:'cash', displayAs:'现金'}, 
				],
				hidenTagsAsEditable: [
					{key:'virtualShare', displayAs:'组合份额'}, 
					{key:'virtualValue', displayAs:'组合净值'}, 
					{key:'virtualValueLYR', displayAs:'组合去年净值'},
				],
				shownTagsAsReadonly: [
					{key:'stockCapitalization', displayAs:'A股市值'}, 
					{key:'hkCapitalization', displayAs:'港股市值'}, 
					{key:'totalCapitalization', displayAs:'总市值'}, 
				],
				data: {
					//Configured
					cash: 20000,
					virtualShare: 100,
					virtualValue: 2.1410,
					virtualValueLYR: 5,
					// fetched
					refrate: 4.9,
					hkExchangeRate: 0.8432,
					hs300: 3225,
					hs300LYR: 4030,
					stockTrend: 100,
					hkTrend: 100,
					// Computed
					stockCapitalization: 20000,
					hkCapitalization: 20000,
					totalCapitalization: 60000,
					estimatedCash: 20000,
					totalProfitRate: 100,
					yearProfitRate: -10,
				},
			},
			tableData: {
				titles: [
						{Key:"mark",Name: "建议"},
						{Key:"id", Name: "代码"},
						{Key:"name", Name: "名称"},
						{Key:"score", Name: "评分"},
						{Key:"price", Name: "价格"},
						{Key:"share", Name: "数量"},
						{Key:"trade", Name: "交易建议"},
						{Key:"capitalization", Name: "市值"},
						{Key:"level", Name: "仓位"},
						{Key:"lowlevel", Name: "下限位"},
						{Key:"uplevel", Name: "上限位"},
						{Key:"marketlevel", Name: "市场仓位"},

				],//titles
				datas: [
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
				
				],//datas
				commands: [
					{Key:"buy", Name:"买进"},
					{Key:"sell", Name:"卖出"},
				],//commands
			},//tableData
		},//componentData
	},
});
