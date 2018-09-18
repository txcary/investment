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
