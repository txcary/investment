var component = {
	props: [
		'button_name',//"submit button's name"
		'datas',//[{Name: "title", Value: "initial value"},...]
	],
	template: ` 
		<div class="form-group">
			<div class="row" v-for="item in datas">
				<div class="col">{{item.Name}}</div>
				<div class="col">
					<input class="for-control" v-model="item.Value"/>
				</div>
			</div>
			<button class="btn btn-success" v-on:click="$emit('submit', datas)">{{button_name}}</button>
		</div>
	`,
};

export {component};