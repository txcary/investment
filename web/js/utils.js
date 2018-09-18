function setCookie(space,name,value,exdays){
    var d = new Date();
    var name = space+"_"+name;
    if(exdays===undefined) {
    	exdays = 1;
    }
    d.setTime(d.getTime()+(exdays*24*60*60*1000));
    var expires = "expires="+d.toGMTString();
    document.cookie = name+"="+value+"; "+expires;
}

function getCookie(space,name){
    var name = space+"_"+name + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i].trim();
        if (c.indexOf(name)==0) { return c.substring(name.length,c.length); }
    }
    return "";
}

function delCookie(space,name){
	setCookie(space,name,"",-1);
}

export {setCookie, getCookie, delCookie};