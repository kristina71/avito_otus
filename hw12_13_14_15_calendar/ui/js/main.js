document.addEventListener('DOMContentLoaded', function () {
    document.getElementById("preloader").classList.remove("active");
    var xhr = new XMLHttpRequest();

    xhr.open('GET', '/calendar/get', false);
    xhr.send();
    if (xhr.status != 200) {
        alert(xhr.status + xhr.responseText);
    } else {
        var obj = JSON.parse(xhr.responseText);

        for (i = 0; i < obj.length; i++) {
            document.getElementById("app").innerHTML += "<div class=\"row\">" +
                "<p>"+obj[i].id+" "+obj[i].title+" "+obj[i].description+
                " "+obj[i].startAt+ " "+obj[i].endAt+ " "+obj[i].remindAt+"</p>";
        }
    }
});
document.getElementById("app").onerror = function () {
    alert("Something went wrong");
};