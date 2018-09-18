var component = {
	props: [
		'titles', //[{Key:"Id", Name:"ID"}, {Key:"Name", Name:"NAME",...]
		'datas', //[{Id: "600519", Name:"myname", Value: "100"},...]
		'commands', //[{Key:"cmd1", Name: "name1"},...]
	],
	template: ` 
		<table class="table table-sm">
			<thead>
				<tr>
					<th v-for="title in titles">{{title.Name}}</th>
					<th v-if="commands!=undefined && commands.length>0"></td>
				</tr>
			</thead>
			<tbody>
				<tr v-for="(item,itemIdx) in datas">
					<td v-for="title in titles">{{item[title.Key]}}</td>
					<td v-if="commands!=undefined && commands.length>0">
						<div class="btn-group" role="group">
							<button v-for="command in commands" v-on:click="$emit('command', {command:command.Key,index:itemIdx})" class="btn btn-link" type="button">{{command.Name}}</button>
						</div>
					</td>
				</tr>
			</tbody>
		</table>
	`,
};

export {component};