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

function HttpGet(url, callback) {
	var xhr = new XMLHttpRequest();
	xhr.open("GET", url, true);
	xhr.onreadystatechange = function() {
		if(xhr.readyState== XMLHttpRequest.DONE) {
			callback(xhr.responseText);
		}
	}
	xhr.send();

}

function HttpPostJson(url, callback, jsonStr) {
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

export {SetCookie, GetCookie, DelCookie, HttpGet, HttpPostJson};