/*
alert("main.js");
import * as mod from "./module.js";

alert(mod.a);
alert(mod.GetB());
*/
import * as vtable from "../js/components/table.js";
import * as veditable from "../js/components/editable.js";
import * as vform from "../js/components/form.js";

var app = new Vue({
	el: '#app',
	data: {
		tableData: {
			titles: [
				{Key:"Id", Name:"ID"},
				{Key:"Name", Name:"名称"},
				{Key:"Value", Name:"值"},
			],
			datas: [
				{Id: "600519", Name:"贵州茅台", Value: "100"},
			],
			commands: [
				{Key:"cmd1", Name:"cmd2"},
			],
		},
		editableData: {
			editing: false,
			title: "my title",
			value: "my value",
		},
		formData: {
			datas: [
				{Name: "input1", Value: "value1"},
				{Name: "input2", Value: ""},
			],
		},
	},
	components: {
		'vtable': vtable.component,
		'veditable': veditable.component,
		'vform': vform.component,
	}
});