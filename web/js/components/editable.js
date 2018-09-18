var component = {
	props: [
		'editing', //true,false
		'title',
		'value',
	],
	template: ` 
		<div class="row">
			<div class="col" v-if="title!=undefined">{{title}}</div>
			<div class="col" v-if="!editing">{{value}}</div>
			<div class="col" v-if="editing">
				<input class="form-control" v-bind:value="value" v-on:input="$emit('input', $event.target.value)"/>
			</div>
		</div>
	`,
};

export {component};