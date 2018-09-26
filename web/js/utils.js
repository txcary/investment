function SetCookie(space,name,value,exdays){
    var d = new Date();
    var name = space+"_"+name;
    if(exdays===undefined) {
    	exdays = 1;
    }
    d.setTime(d.getTime()+(exdays*24*60*60*1000));
    var expires = "expires="+d.toGMTString();
    document.cookie = name+"="+value+"; "+expires;
}

function GetCookie(space,name){
    var name = space+"_"+name + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i].trim();
        if (c.indexOf(name)==0) { return c.substring(name.length,c.length); }
    }
    return "";
}

function DelCookie(space,name){
	SetCookie(space,name,"",-1);
}

function HttpGetAsync(url, callback) {
	var xhr = new XMLHttpRequest();
	xhr.open("GET", url, true);
	xhr.onreadystatechange = function() {
		if(xhr.readyState== XMLHttpRequest.DONE) {
			callback(xhr.responseText);
		}
	}
	xhr.send();

}

function HttpGet(url) {
	var xhr = new XMLHttpRequest();
	xhr.open("GET", url, false);
	xhr.send();
	if(xhr.readyState== XMLHttpRequest.DONE) {
		return xhr.responseText;
	} else {
		console.log("Error: HttpGet Error!");
	}
	return undefined;
}

function HttpPostJsonAsync(url, jsonStr, callback) {
	var xhr = new XMLHttpRequest();
	xhr.open("POST", url, true);
	xhr.setRequestHeader('content-type', 'application/json');
	xhr.onreadystatechange = function() {
		if(xhr.readyState== XMLHttpRequest.DONE){
			callback(xhr.responseText);
		}
	}
	xhr.send(jsonStr);
}

function HttpPostJson(url, jsonStr) {
	var xhr = new XMLHttpRequest();
	xhr.open("POST", url, false);
	xhr.setRequestHeader('content-type', 'application/json');
	xhr.send(jsonStr);
	if(xhr.readyState== XMLHttpRequest.DONE){
		return xhr.responseText;
	} else {
		console.log("Error: HttpPostJson Error!");
		return undefined;
	}
}

export {SetCookie, GetCookie, DelCookie, HttpGet, HttpPostJson, HttpGetAsync, HttpPostJsonAsync};