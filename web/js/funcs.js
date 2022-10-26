
function Register() {

var name = document.getElementById("name").value;

var email = document.getElementById("email").value;

var password = document.getElementById("password").value;

    //Send a request to http://127.0.0.1:8080/v1/auth/register/ with user,email,password
    //Log response


    $.ajax({
        url: "http://127.0.0.1:8080/v1/auth/register/",
        contentType: "application/json",
        data: JSON.stringify({ 'name': name, 'user': email, 'password': password }),
        method: 'POST'
    }).done(function (_args) {
        console.log(_args);
    }
    ).fail(function (jqXHR, textStatus, errorThrown) {
        console.log(jqXHR);
        console.log(textStatus);
        console.log(errorThrown);
    }
    );

}