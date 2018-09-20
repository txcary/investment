var Data = {
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
};

export {Data};