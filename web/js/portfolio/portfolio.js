"use strict"

import * as vtable from "../components/table.js";
import * as veditable from "../components/editable.js";
import * as vform from "../components/form.js";
import * as utils from "../utils.js";

import {Data} from "./portfolio_data.js"


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

var app = new Vue({
	el: '#app',
	components: {
		'vtable': vtable.component,
		'veditable': veditable.component,
		'vform': vform.component,
	},
	created: function() {
		this.userName = utils.GetCookie("portfolio","userName");
		this.userPasswd = utils.GetCookie("portfolio","userPasswd");
		if(this.userName!="" && this.userPassed!="") {
			this.isLogined=true;
			this.getPortfolioFromServer();
this.update();
		}else{
			this.logout();
		}
	},
	methods: {
		logout: function() {
			this.userName="";
			this.userPasswd="";
			this.isLogined=false;
			utils.DelCookie("portfolio", "userName");
			utils.DelCookie("portfolio", "userPasswd");
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
			utils.SetCookie("portfolio", "userName", this.userName);
			utils.SetCookie("portfolio", "userPasswd", this.userPasswd);
			this.getPortfolioFromServer();
		},
		getStock: function(id) {
			if(id===undefined) {
				return undefined;
			}
			if(this.stocks==undefined) {
				this.stocks = new Object();
			}
			if(this.stocks[id]==undefined) {
				var resp = utils.HttpGet("/stock/"+id);
				if(resp===undefined) {
					alert("Can not get data of id "+id);
					return undefined;
				}
				var stock = JSON.parse(resp);
				if(stock.code!=undefined) {
					if(stock.code>0) {
						alert(stock.msg);
						return undefined;
					}
				}
				if(stock.Name==undefined || stock.Id==undefined) {
					console.log("Error: stock data not valid.");
					console.log(stock);
					return undefined
				}
				this.stocks[id] = stock;
			}
			return this.stocks[id];
		},
		prepareServerDataJsonString() {
			var obj = Object();
			obj.statistics = this.componentData.statisticsData.data;
			obj.portfolio = this.componentData.tableData.datas;
			return JSON.stringify(obj);
		},
		getPortfolioFromServer() {
			var str = secureJson.GenerateJson(this.userName, this.userPasswd, "");
			var response = utils.HttpPostJson("/portfolio/getjson", str);
			if(response!=undefined) {
				var obj = JSON.parse(response);
				if(obj.code != undefined) {
					if(obj.code!=0) {
						alert(obj.msg);
						this.logout();
						return;
					}
				}
				if(obj.EncryptedData == undefined) {
					console.log("Error: No Ecrypted Data");
					return;
				}

				console.log("Encrypted-Data: "+obj.EncryptedData);
				var dataStr = secureJson.Decrypt(this.userName, this.userPasswd, obj.EncryptedData)
				console.log("Decrypted-Data: "+dataStr);
				var data = JSON.parse(dataStr);
				if(data.statistics != undefined) {
					this.componentData.statisticsData.data = data.statistics; 
				}
				if(data.portfolio!=undefined) {
					this.componentData.tableData.datas = data.portfolio;
				}
			} else {
				console.log("Error: can not connect to server!");
				this.logout();
			}
		},
		putPortfolioToServer() {
			var dataStr = this.prepareServerDataJsonString();
			var str = secureJson.GenerateJson(this.userName, this.userPasswd, dataStr);
			console.log(str);
			var callback = function(response) {
				console.log(response);
			}
			utils.HttpPostJsonAsync("/portfolio/putjson", str, callback);
		},
		getPortfolioIdx: function(id) {
			for(var i=0;i<this.componentData.tableData.datas.length;i++) {
				if(this.componentData.tableData.datas[i].id==id) {
					return i;
				}
			}
			return undefined;
		},
		update() {
			//TODO: compute
			var stat = this.componentData.statisticsData.data;
			var datas = this.componentData.tableData.datas;

			stat.totalCapitalization = 0;
			for(var i=0; i<datas.length; i++) {
				var stock = this.getStock(datas[i].id);
				datas[i].price = stock.Price;
				datas[i].capitalization = datas[i].price*datas[i].share;
				
				var factor1 = stock.SaftyFactor*stock.ExpectionFactor*stock.CashFactor*stock.FlowFactor;
				var factor2 = stock.SaftyFactor;
				datas[i].lowlevel = parseInt(100*factor1*(1-stat.refrate/(stock.ReturnStd*100))*0.2);
				datas[i].uplevel = parseInt(100*factor2*(1-stat.refrate/(stock.ReturnHigh*100))*0.2);

				stat.totalCapitalization += datas[i].capitalization;
			}

			for(var i=0; i<datas.length; i++) {
				datas[i].level = parseInt(100*datas[i].capitalization/stat.totalCapitalization);
			}
			this.putPortfolioToServer();
		},
		doTrade: function(id, amount, price) {
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
				var stock = this.getStock(id);

				var item = {
					id: id,
					name: stock.Name,
					share: 0,
					price: 0,
				};
				this.componentData.tableData.datas.push(item);
				idx = this.getPortfolioIdx(id);
			}
			var tradedShare = this.componentData.tableData.datas[idx].share + parseInt(amount);
			if(tradedShare <0 ) {
				alert("No enough share!");
				return;
			}
			if(tradedShare==0) {
				this.componentData.tableData.datas.splice(idx,1);
			} else {
				this.componentData.tableData.datas[idx].share = tradedShare;
				this.componentData.tableData.datas[idx].price = parseFloat(price);
			}
			this.update();
		},
		doScore: function(id, amount) {
			var idx = this.getPortfolioIdx(id);
			if(idx===undefined) {
				alert("ID not valid!");
				return;
			}
			this.componentData.tableData.datas[idx].score = amount;
			this.update();
		},
		doAdd: function(amount) {
			var amountFloat = parseFloat(amount);
			this.componentData.statisticsData.data.cash += amountFloat;
			this.componentData.statisticsData.data.virtualShare += amountFloat/this.componentData.statisticsData.data.virtualValue;
			this.update();
		},
		doStatistics() {
			this.componentData.statisticsData.editing=false;
			this.update();
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
	data: Data,
});



});